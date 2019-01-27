package goff

//#cgo pkg-config: libavutil libavcodec
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
import "C"

type CodecParameters = C.struct_AVCodecParameters

func (cp *CodecParameters) CodecType() MediaType {
	return MediaType(cp.codec_type)
}

func (cp *CodecParameters) CodecID() CodecID {
	return CodecID(cp.codec_id)
}

func (cp *CodecParameters) PixelFormat() PixelFormat {
	return PixelFormat(cp.format)
}

func (cp *CodecParameters) SampleFormat() SampleFormat {
	return SampleFormat(cp.format)
}

// Fill the codec context based on the values from the supplied codec parameters.
//
// Any allocated fields in codec that have a corresponding field in par are
// freed and replaced with duplicates of the corresponding field in par. Fields
// in codec that do not have a counterpart in par are not touched.
//
// Returns
//     >= 0 on success, a negative AVERROR code on failure.
func (cp *CodecParameters) ToContext(cctx *CodecContext) error {
	return CheckErr(C.avcodec_parameters_to_context(cctx, cp))
}

// Fill the parameters struct based on the values from the supplied codec
// context.
//
// Any allocated fields in par are freed and replaced with duplicates of the
// corresponding fields in codec.
//
// Returns
//     >= 0 on success, a negative AVERROR code on failure
func (cp *CodecParameters) FromContext(cctx *CodecContext) error {
	return CheckErr(C.avcodec_parameters_from_context(cp, cctx))
}

func (cp *CodecParameters) Copy(cpDest *CodecParameters) error {
	return CheckErr(C.avcodec_parameters_copy(cp, cpDest))
}
