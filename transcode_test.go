package goff_test

import (
	"testing"
	"unsafe"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestTranscode(t *testing.T) {
	assert := assert.New(t)

	loggedObjects := make(map[uintptr]string)

	goff.LogSetCallback(goff.LogLevel_DEBUG, func(ptr uintptr, level goff.LogLevel, line string, printPrefix bool) {
		if lo, ok := loggedObjects[ptr]; ok {
			t.Logf("(%s) %s", lo, line)
		} else {
			t.Logf("(unknown %x) %s", ptr, line)
		}
	})

	must := func(err error) {
		if err != nil {
			assert.NoError(err)
			t.FailNow()
		}
	}

	t.Logf("Opening input...")

	inputPath := "testdata/sample.mp4"

	inputFormatContext := goff.FormatAllocContext()
	loggedObjects[uintptr(unsafe.Pointer(inputFormatContext))] = "demuxer"

	err := inputFormatContext.OpenInput(inputPath, nil, nil)
	must(err)

	assert.EqualValues(2, inputFormatContext.NbStreams())

	inputVideoStream := inputFormatContext.Streams()[0]
	assert.EqualValues("video", inputVideoStream.CodecParameters().CodecType().String())

	inputFormatContext.DumpFormat(0, inputPath, false)

	numFrames := 0

	inputVideoDecoder := inputVideoStream.CodecParameters().CodecID().FindDecoder()
	assert.NotNil(inputVideoDecoder)

	inputVideoDecoderContext := inputVideoDecoder.AllocContext3()
	assert.NotNil(inputVideoDecoderContext)
	loggedObjects[uintptr(unsafe.Pointer(inputVideoDecoderContext))] = "decoder"

	inputVideoStream.CodecParameters().ToContext(inputVideoDecoderContext)

	err = inputVideoDecoderContext.Open2(inputVideoDecoder, nil)
	must(err)

	t.Logf("Opening output...")

	outPath := "out.mp4"

	outputIOContext, err := goff.IOOpen(outPath, goff.IO_FLAG_WRITE)
	must(err)
	defer outputIOContext.Close()
	loggedObjects[uintptr(unsafe.Pointer(outputIOContext))] = "output-io"

	outputFormat := goff.GuessFormat("mp4", "", "")
	if outputFormat == nil {
		assert.NotNil(outputFormat)
		t.FailNow()
	}
	t.Logf("Guessed format: %s", outputFormat.LongName())

	outputFormatContext, err := goff.FormatAllocOutputContext2(outputFormat, "", outPath)
	must(err)
	defer outputFormatContext.Free()
	loggedObjects[uintptr(unsafe.Pointer(outputFormatContext))] = "muxer"

	outputFormatContext.SetPB(outputIOContext)
	outputFormatContext.SetOutputFormat(outputFormat)

	outputVideoStream := outputFormatContext.NewStream(inputVideoDecoder)
	outputVideoStream.SetID(0)

	outputCodecID := goff.CodecID_H264
	outputVideoEncoder := outputCodecID.FindEncoder()
	assert.NotNil(outputVideoEncoder)

	outputVideoEncoderContext := outputVideoEncoder.AllocContext3()
	assert.NotNil(outputVideoEncoderContext)
	loggedObjects[uintptr(unsafe.Pointer(outputVideoEncoderContext))] = "encoder"

	outputVideoEncoderContext.SetCodecType(goff.MediaType_Video)
	outputVideoEncoderContext.SetCodecID(goff.CodecID_H264)
	outputVideoEncoderContext.SetPixelFormat(goff.PixelFormat_YUV420P)

	outputVideoStream.SetTimeBase(goff.TIME_BASE_Q)
	outputVideoEncoderContext.SetTimeBase(goff.NewRational(1, 120))
	outputVideoEncoderContext.SetWidth(inputVideoDecoderContext.Width())
	outputVideoEncoderContext.SetHeight(inputVideoDecoderContext.Height())

	goff.OptSet(outputVideoEncoderContext.PrivData(), "preset", "ultrafast", goff.SearchFlags_CHILDREN)

	err = outputVideoEncoderContext.Open2(outputVideoEncoder, nil)
	must(err)

	err = outputVideoStream.CodecParameters().FromContext(outputVideoEncoderContext)
	must(err)

	// outputFormatContext.DumpFormat(0, outPath, true)

	must(outputFormatContext.WriteHeader(nil))

	inputFrame := goff.FrameAlloc()
	assert.NotNil(inputFrame)
	defer inputFrame.Free()

	var packet goff.Packet

	writeEncodedPackets := func(last bool) {
		for {
			var outPacket goff.Packet
			outPacket.Init()

			err := outputVideoEncoderContext.ReceivePacket(&outPacket)
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

			outPacket.SetStreamIndex(outputVideoStream.Index())
			outPacket.RescaleTs(outputVideoEncoderContext.TimeBase(), outputVideoStream.TimeBase())

			err = outputFormatContext.InterleavedWriteFrame(&outPacket)
			must(err)
		}
	}

	receiveDecodedFrames := func() {
		for {
			err = inputVideoDecoderContext.ReceiveFrame(inputFrame)
			if err != nil {
				if goff.IsEAGAIN(err) || goff.IsEOF(err) {
					return
				}
				must(err)
			}
			numFrames++

			scaledPts := inputFrame.Pts().Rescale(inputVideoDecoderContext.TimeBase(), outputVideoEncoderContext.TimeBase())
			if scaledPts.IsNop() {
				scaledPts = 0
			}
			inputFrame.SetPts(scaledPts)

			err = outputVideoEncoderContext.SendFrame(inputFrame)
			must(err)

			writeEncodedPackets(false)
		}
	}

	for {
		packet.Init()

		err = inputFormatContext.ReadFrame(&packet)
		if err != nil {
			if goff.IsEOF(err) {
				break
			}
			must(err)
		}

		if packet.StreamIndex() != inputVideoStream.Index() {
			// ignore audio packets
			continue
		}

		packet.RescaleTs(inputVideoStream.TimeBase(), inputVideoDecoderContext.TimeBase())

		err = inputVideoDecoderContext.SendPacket(&packet)
		if err != nil {
			must(err)
		}

		receiveDecodedFrames()
	}
	receiveDecodedFrames()

	flushEncoder := func() {
		err := outputVideoEncoderContext.SendFrame(nil)
		must(err)

		writeEncodedPackets(true)
	}

	flushEncoder()

	t.Logf("Processed %d inputFrames in total", numFrames)
	assert.EqualValues(23, numFrames)

	err = outputFormatContext.WriteTrailer()
	must(err)

	defer inputFormatContext.Free()
}
