package goff

//#cgo pkg-config: libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libswscale/swscale.h>
import "C"

type SwsContext = C.struct_SwsContext
type SwsFilter = C.struct_SwsFilter

type SwsFlags = int

var (
	SwsFlags_Bicubic SwsFlags = C.SWS_BICUBIC
)

// Check if context can be reused, otherwise reallocate a new one.
// If context is NULL, just calls sws_getContext() to get a new context.
// Otherwise, checks if the parameters are the ones already saved in context. If
// that is the case, returns the current context. Otherwise, frees context and
// gets a new context with the new parameters.
//
// Be warned that srcFilter and dstFilter are not checked, they are assumed to
// remain the same.
func SwsGetCachedContext(
	context *SwsContext,
	srcW int,
	srcH int,
	srcFormat PixelFormat,
	dstW int,
	dstH int,
	dstFormat PixelFormat,
	flags SwsFlags,
	srcFilter *SwsFilter,
	dstFilter *SwsFilter,
	param *float64,
) *SwsContext {
	return C.sws_getCachedContext(
		context,
		C.int(srcW), C.int(srcH), srcFormat.C(),
		C.int(dstW), C.int(dstH), dstFormat.C(),
		C.int(flags), srcFilter, dstFilter,
		(*C.double)(param),
	)
}

func (swctx *SwsContext) Scale(
	srcSlice Planes, srcStride Linesizes,
	srcSliceY int, srcSliceH int,
	dst Planes, dstStride Linesizes,
) error {
	return CheckErr(C.sws_scale(
		swctx,
		(**C.uchar)(&srcSlice[0]),
		(*C.int)(&srcStride[0]),
		C.int(srcSliceY),
		C.int(srcSliceH),
		(**C.uchar)(&dst[0]),
		(*C.int)(&dstStride[0]),
	))
}

func (swctx *SwsContext) Free() {
	C.sws_freeContext(swctx)
}
