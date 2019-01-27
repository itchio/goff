package goff_test

import (
	"testing"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	assert := assert.New(t)

	ctx, err := goff.FormatOpenInput("testdata/sample.mp4", nil, nil)
	assert.NoError(err)

	err = ctx.FindStreamInfo(nil)
	assert.NoError(err)

	assert.EqualValues(2, ctx.NbStreams())

	vst := ctx.Streams()[0]
	assert.EqualValues("video", vst.CodecParameters().CodecType().String())

	ast := ctx.Streams()[1]
	assert.EqualValues("audio", ast.CodecParameters().CodecType().String())

	numFrames := 0
	vcod := vst.CodecParameters().CodecID().FindDecoder()
	assert.NotNil(vcod)
	assert.EqualValues("h264", vcod.Name())

	vdec := vcod.AllocContext3()
	assert.NotNil(vdec)

	vst.CodecParameters().ToContext(vdec)

	err = vdec.Open2(vcod, nil)
	assert.NoError(err)

	frame := goff.FrameAlloc()
	assert.NotNil(frame)
	defer frame.Free()

	var packet goff.Packet

	readFrames := func() {
		for {
			err = vdec.ReceiveFrame(frame)
			if err != nil {
				if goff.IsEAGAIN(err) {
					return
				}
				if goff.IsEOF(err) {
					return
				}
				assert.NoError(err)
				t.FailNow()
			}

			numFrames++
		}
	}

	for {
		packet.Init()

		err = ctx.ReadFrame(&packet)
		if err != nil {
			if goff.IsEOF(err) {
				break
			}
			assert.NoError(err)
			t.FailNow()
		}

		if packet.StreamIndex() != vst.Index() {
			// ignore audio packets
			continue
		}

		packet.RescaleTs(vst.TimeBase(), vdec.TimeBase())

		err = vdec.SendPacket(&packet)
		if err != nil {
			assert.NoError(err)
			t.FailNow()
		}

		readFrames()
	}
	readFrames()

	t.Logf("Processed %d frames in total", numFrames)
	assert.EqualValues(23, numFrames)

	defer ctx.Free()
}
