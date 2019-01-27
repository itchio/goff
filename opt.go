package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
import "C"
import "unsafe"

type SearchFlags = C.int

var (
	SearchFlags_CHILDREN SearchFlags = C.AV_OPT_SEARCH_CHILDREN
)

func OptSet(obj unsafe.Pointer, name string, val string, searchFlags SearchFlags) error {
	name_ := CString(name)
	defer FreeString(name_)

	val_ := CString(val)
	defer FreeString(val_)

	return CheckErr(C.av_opt_set(obj, name_, val_, C.int(searchFlags)))
}
