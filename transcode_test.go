package goff_test

import (
	"testing"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestTranscode(t *testing.T) {
	assert := assert.New(t)

	goff.LogSetCallback(func(line string) {
		t.Logf(line)
	})

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

	ctx.DumpFormat(0, "", false)

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
	octx.SetOutputFormat(ofmt)

	ovst := octx.NewStream(vcod)
	ovst.SetID(0)

	ocid := goff.CodecID_H264
	ovcod := ocid.FindEncoder()
	assert.NotNil(ovcod)

	venc := ovcod.AllocContext3()
	assert.NotNil(venc)

	venc.SetCodecType(goff.MediaType_Video)
	venc.SetCodecID(goff.CodecID_H264)
	venc.SetPixelFormat(goff.PixelFormat_YUV420P)

	ovst.SetTimeBase(goff.TIME_BASE_Q)
	venc.SetTimeBase(ovst.TimeBase())
	venc.SetGOPSize(120)
	venc.SetMaxBFrames(16)
	venc.SetWidth(vdec.Width())
	venc.SetHeight(vdec.Height())

	crf := 20
	venc.SetQMin(crf)
	venc.SetQMax(crf)
	venc.OrFlag(goff.CodecFlags_GLOBAL_HEADER)
	venc.SetProfile(goff.Profile_H264_BASELINE)
	goff.OptSet(venc.PrivData(), "preset", "ultrafast", goff.SearchFlags_CHILDREN)

	err = venc.Open2(ovcod, nil)
	must(err)

	err = ovst.CodecParameters().FromContext(venc)
	must(err)

	octx.DumpFormat(0, "", true)

	frame := goff.FrameAlloc()
	assert.NotNil(frame)
	defer frame.Free()

	var oPts goff.Timing = 0

	var packet goff.Packet

	writeFrames := func(last bool) {
		for {
			var opkt goff.Packet
			opkt.Init()

			err := venc.ReceivePacket(&opkt)
			if err != nil {
				if goff.IsEOF(err) {
					// all flushed!
					return
				}
				if !last && goff.IsEAGAIN(err) {
					// will get more packets later
					return
				}
				must(err)
			}

			opkt.SetPts(oPts)
			oPts += 1

			opkt.SetStreamIndex(ovst.Index())
			err = octx.InterleavedWriteFrame(&opkt)
			must(err)
		}
	}

	readFrames := func() {
		for {
			err = vdec.ReceiveFrame(frame)
			if err != nil {
				if goff.IsEAGAIN(err) || goff.IsEOF(err) {
					return
				}
				must(err)
			}
			t.Logf("Received frame! PTS %v", frame.Pts().AsDuration(vdec.TimeBase()))
			t.Logf("Frame format: %s", frame.Format().Name())
			t.Logf("Frame res: %dx%d", frame.Width(), frame.Height())
			numFrames++

			t.Logf("Sending frame..")
			err = venc.SendFrame(frame)
			must(err)

			writeFrames(false)
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

		writeFrames(true)
	}

	flushEncoder()

	t.Logf("Processed %d frames in total", numFrames)
	assert.EqualValues(23, numFrames)

	err = octx.WriteTrailer()
	must(err)

	defer ctx.Free()
}
