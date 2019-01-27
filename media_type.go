package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
import "C"

type MediaType C.enum_AVMediaType

var (
	// Usually treated as AVMEDIA_TYPE_DATA.
	MediaType_Unknown MediaType = C.AVMEDIA_TYPE_UNKNOWN
	MediaType_Video   MediaType = C.AVMEDIA_TYPE_VIDEO
	MediaType_Audio   MediaType = C.AVMEDIA_TYPE_AUDIO
	// Opaque data information usually continuous.
	MediaType_Data     MediaType = C.AVMEDIA_TYPE_DATA
	MediaType_Subtitle MediaType = C.AVMEDIA_TYPE_SUBTITLE
	// Opaque data information usually sparse.
	MediaType_Attachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT
	MediaType_Nb         MediaType = C.AVMEDIA_TYPE_NB
)

func (mt MediaType) String() string {
	return C.GoString(C.av_get_media_type_string(mt.C()))
}

func (mt MediaType) C() C.enum_AVMediaType {
	return C.enum_AVMediaType(mt)
}
