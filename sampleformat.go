package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/samplefmt.h>
import "C"

type (
	SampleFormat C.enum_AVSampleFormat
)

func (sf SampleFormat) C() C.enum_AVSampleFormat {
	return C.enum_AVSampleFormat(sf)
}

var (
	SampleFormat_NONE SampleFormat = C.AV_SAMPLE_FMT_NONE
	// unsigned 8 bits
	SampleFormat_U8 SampleFormat = C.AV_SAMPLE_FMT_U8
	// signed 16 bits
	SampleFormat_S16 SampleFormat = C.AV_SAMPLE_FMT_S16
	// signed 32 bits
	SampleFormat_S32 SampleFormat = C.AV_SAMPLE_FMT_S32
	// float
	SampleFormat_FLT SampleFormat = C.AV_SAMPLE_FMT_FLT
	// double
	SampleFormat_DBL SampleFormat = C.AV_SAMPLE_FMT_DBL
	// unsigned 8 bits, planar
	SampleFormat_U8P SampleFormat = C.AV_SAMPLE_FMT_U8P
	// signed 16 bits, planar
	SampleFormat_S16P SampleFormat = C.AV_SAMPLE_FMT_S16P
	// signed 32 bits, planar
	SampleFormat_S32P SampleFormat = C.AV_SAMPLE_FMT_S32P
	// float, planar
	SampleFormat_FLTP SampleFormat = C.AV_SAMPLE_FMT_FLTP
	// double, planar
	SampleFormat_DBLP SampleFormat = C.AV_SAMPLE_FMT_DBLP
	// signed 64 bits
	SampleFormat_S64 SampleFormat = C.AV_SAMPLE_FMT_S64
	// signed 64 bits, planar
	SampleFormat_S64P SampleFormat = C.AV_SAMPLE_FMT_S64P
)
