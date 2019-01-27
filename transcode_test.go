package goff_test

import (
	"testing"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestTranscode(t *testing.T) {
	assert := assert.New(t)

	must := func(err error) {
		if err != nil {
			assert.NoError(err)
			t.FailNow()
		}
	}

	t.Logf("Opening input...")

	ctx, err := goff.FormatOpenInput("testdata/sample.mp4", nil, nil)
	must(err)

	err = ctx.FindStreamInfo(nil)
	must(err)

	assert.EqualValues(2, ctx.NbStreams())

	vst := ctx.Streams()[0]
	assert.EqualValues("video", vst.CodecParameters().CodecType().String())

	numFrames := 0

	vcod := vst.CodecParameters().CodecID().FindDecoder()
	assert.NotNil(vcod)

	vdec := vcod.AllocContext3()
	assert.NotNil(vdec)

	vst.CodecParameters().ToContext(vdec)

	err = vdec.Open2(vcod, nil)
	must(err)

	t.Logf("Opening output...")

	outPath := "out.mp4"

	var pb *goff.IOContext
	err = goff.IOOpen(&pb, outPath, goff.IO_FLAG_WRITE)
	must(err)
	defer pb.Close()

	ofmt := goff.GuessFormat("mp4", "", "")
	if ofmt == nil {
		assert.NotNil(ofmt)
		t.FailNow()
	}
	t.Logf("Guessed format: %s", ofmt.LongName())

	octx, err := goff.FormatAllocOutputContext2(ofmt, "", outPath)
	must(err)
	defer octx.Free()

	octx.SetPB(pb)

	ovst := octx.NewStream(vcod)
	ovst.SetID(0)

	venc := vcod.AllocContext3()
	assert.NotNil(venc)

	venc.SetCodecType(goff.MediaType_Video)
	venc.SetCodecID(goff.CodecID_H264)
	venc.SetPixelFormat(goff.PixelFormat_YUV420P)

	ovst.SetTimeBase(goff.TIME_BASE_Q)
	venc.SetTimeBase(ovst.TimeBase())
	venc.SetGOPSize(120)
	venc.SetMaxBFrames(16)

	crf := 20
	venc.SetQMin(crf)
	venc.SetQMax(crf)
	venc.OrFlag(goff.CodecFlags_GLOBAL_HEADER)
	venc.SetProfile(goff.Profile_H264_BASELINE)
	goff.OptSet(venc.PrivData(), "preset", "ultrafast", goff.SearchFlags_CHILDREN)

	err = venc.Open2(vcod, nil)
	must(err)

	err = ovst.CodecParameters().FromContext(venc)

	frame := goff.FrameAlloc()
	assert.NotNil(frame)
	defer frame.Free()

	var packet goff.Packet

	readFrames := func() {
		for {
			err = vdec.ReceiveFrame(frame)
			if err != nil {
				if goff.IsEAGAIN(err) || goff.IsEOF(err) {
					return
				}
				must(err)
			}
			numFrames++

			err = venc.SendFrame(frame)
			must(err)

			for {
				var opkt goff.Packet
				opkt.Init()

				err = venc.ReceivePacket(&opkt)
				if err != nil {
					if goff.IsEAGAIN(err) {
						return
					}
					must(err)
				}

				opkt.SetStreamIndex(ovst.Index())
				err = octx.InterleavedWriteFrame(&packet)
				must(err)
			}
		}
	}

	for {
		packet.Init()

		err = ctx.ReadFrame(&packet)
		if err != nil {
			if goff.IsEOF(err) {
				break
			}
			must(err)
		}

		if packet.StreamIndex() != vst.Index() {
			// ignore audio packets
			continue
		}

		packet.RescaleTs(vst.TimeBase(), vdec.TimeBase())

		err = vdec.SendPacket(&packet)
		if err != nil {
			must(err)
		}

		readFrames()
	}
	readFrames()

	flushEncoder := func() {
		err := venc.SendFrame(nil)
		must(err)

		for {
			var opkt goff.Packet
			opkt.Init()

			err := venc.ReceivePacket(&opkt)
			if err != nil {
				if goff.IsEOF(err) {
					// all flushed!
					return
				}
				must(err)
			}

			opkt.SetStreamIndex(ovst.Index())
			err = octx.InterleavedWriteFrame(&packet)
			must(err)
		}
	}

	flushEncoder()

	t.Logf("Processed %d frames in total", numFrames)
	assert.EqualValues(23, numFrames)

	defer ctx.Free()
}
