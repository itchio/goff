package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
import "C"

type PictureType C.enum_AVPictureType

var (
	// Undefined
	PictureType_None PictureType = C.AV_PICTURE_TYPE_NONE
	// Intra
	PictureType_I PictureType = C.AV_PICTURE_TYPE_I
	// Predicted
	PictureType_P PictureType = C.AV_PICTURE_TYPE_P
	// Bi-dir predicted
	PictureType_B PictureType = C.AV_PICTURE_TYPE_B
	// S(GMC)-VOP MPEG-4
	PictureType_S PictureType = C.AV_PICTURE_TYPE_S
	// Switching Intra.
	PictureType_SI PictureType = C.AV_PICTURE_TYPE_SI
	// Switching Predicted
	PictureType_SP PictureType = C.AV_PICTURE_TYPE_SP
	// BI TYPE
	PictureType_BI PictureType = C.AV_PICTURE_TYPE_BI
)

func (pt PictureType) String() string {
	c := byte(C.av_get_picture_type_char(pt.C()))
	return string([]byte{c})
}

func (pt PictureType) C() C.enum_AVPictureType {
	return C.enum_AVPictureType(pt)
}
