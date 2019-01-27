package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/frame.h>
import "C"
import (
	"reflect"
	"unsafe"
)

//------------------------------------
// @Frame
//------------------------------------

type (
	Frame = C.struct_AVFrame
)

// Allocate an AVFrame and set its fields to default values.
//
// The resulting struct must be freed using av_frame_free().
//
// Returns
//     An AVFrame filled with default values or NULL on failure.
//
// Note
//     this only allocates the AVFrame itself, not the data buffers. Those must be
//     allocated through other means, e.g. with av_frame_get_buffer() or manually.
func FrameAlloc() *Frame {
	return C.av_frame_alloc()
}

// Free the frame and any dynamically allocated objects in it, e.g.
//
// extended_data. If the frame is reference counted, it will be unreferenced first.
func (frame *Frame) Free() {
	C.av_frame_free(&frame)
}

// Allocate new buffer(s) for audio or video data.
//
// The following fields must be set on frame before calling this function:
//
// format (pixel format for video, sample format for audio)
// width and height for video
// nb_samples and channel_layout for audio
//
// This function will fill AVFrame.data and AVFrame.buf arrays and, if
// necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
// For planar formats, one buffer will be allocated for each plane.
//
// Warning
// if frame already has been allocated, calling this function will leak
// memory. In addition, undefined behavior can occur in certain cases.
//
// Parameters
// frame	frame in which to store the new buffers.
// align	Required buffer size alignment. If equal to 0, alignment will be chosen
// automatically for the current CPU. It is highly recommended to pass 0 here
// unless you know what you are doing.
//
// Returns
// 0 on success, a negative AVERROR on error.
func (frame *Frame) GetBuffer(align int) error {
	return CheckErr(C.av_frame_get_buffer(frame, C.int(align)))
}

// Ensure that the frame data is writable, avoiding data copy if possible.
//
// Do nothing if the frame is writable, allocate new buffers and copy the data
// if it is not.
//
// Returns
//     0 on success, a negative AVERROR on error.
func (frame *Frame) MakeWritable() error {
	return CheckErr(C.av_frame_make_writable(frame))
}

// Presentation timestamp in time_base units (time when frame should be shown to user).
func (frame *Frame) Pts() Timing {
	return frame.pts
}

func (frame *Frame) Width() int {
	return int(frame.width)
}

func (frame *Frame) SetWidth(width int) {
	frame.width = C.int(width)
}

func (frame *Frame) Height() int {
	return int(frame.height)
}

func (frame *Frame) SetHeight(height int) {
	frame.height = C.int(height)
}

func (frame *Frame) Format() PixelFormat {
	return PixelFormat(frame.format)
}

func (frame *Frame) SetFormat(pf PixelFormat) {
	frame.format = C.int(pf)
}

type Plane = *C.uchar
type Planes = [8]Plane

func (frame *Frame) Data() Planes {
	return frame.data
}

func (frame *Frame) PlaneData(index int) []uint8 {
	linesize := frame.Linesize()[index]
	data := frame.Data()[index]

	var d []uint8
	dsize := int(linesize) * frame.Height()
	{
		sh := (*reflect.SliceHeader)((unsafe.Pointer(&d)))
		sh.Cap = dsize
		sh.Len = dsize
		sh.Data = uintptr(unsafe.Pointer(data))
	}
	return d
}

type Linesize = C.int
type Linesizes = [8]Linesize

func (frame *Frame) Linesize() Linesizes {
	return frame.linesize
}
