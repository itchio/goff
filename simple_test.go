package goff_test

import (
	"log"
	"testing"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
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
	for {
		packet.Init()

		err = ctx.ReadFrame(&packet)
		if err != nil {
			break
		}
		log.Printf("ReadFrame, got packet for stream %d", packet.StreamIndex())
		// assert.NoError(err)

		if packet.StreamIndex() != vst.Index() {
			// ignore audio packets
			continue
		}

		packet.RescaleTs(vst.TimeBase(), vdec.TimeBase())

		err = vdec.SendPacket(&packet)
		assert.NoError(err)

		for {
			err = vdec.ReceiveFrame(frame)
			if err != nil {
				log.Printf("in receive_frame: %v", err)
				break
			}
			// assert.NoError(err)

			numFrames++
		}
	}
	assert.EqualValues(23, numFrames)

	defer ctx.Free()
}
