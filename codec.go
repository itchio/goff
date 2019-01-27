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

type Codec = C.struct_AVCodec

// Allocate an AVCodecContext and set its fields to default values.
//
// The resulting struct should be freed with avcodec_free_context().
//
// Parameters
//     codec if non-NULL, allocate private data and initialize defaults for the
//     given codec. It is illegal to then call avcodec_open2() with a different
//     codec. If NULL, then the codec-specific defaults won't be initialized, which
//     may result in suboptimal default settings (this is important mainly for
//     encoders, e.g. libx264).
//
// Returns
//     An AVCodecContext filled with default values or NULL on failure.
func (codec *Codec) AllocContext3() *CodecContext {
	return C.avcodec_alloc_context3(codec)
}

func (codec *Codec) String() string {
	return codec.Name()
}

func (codec *Codec) Name() string {
	return C.GoString(codec.name)
}

func (codec *Codec) LongName() string {
	return C.GoString(codec.long_name)
}

func (codec *Codec) ID() CodecID {
	return CodecID(codec.id)
}

func (codec *Codec) Type() MediaType {
	return MediaType(codec._type)
}
