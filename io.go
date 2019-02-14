package goff

//#cgo pkg-config: libavutil libavformat
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavformat/avio.h>
import "C"

type IOFlags = int

var (
	IO_FLAG_READ       IOFlags = C.AVIO_FLAG_READ
	IO_FLAG_WRITE      IOFlags = C.AVIO_FLAG_WRITE
	IO_FLAG_READ_WRITE IOFlags = C.AVIO_FLAG_READ_WRITE
)

type IOContext = C.AVIOContext

// Create and initialize a AVIOContext for accessing the resource indicated by url.
//
// Note
//   When the resource indicated by url has been opened in read+write mode, the AVIOContext can be used only for writing.
//
// Parameters
//   s	Used to return the pointer to the created AVIOContext. In case of failure the pointed to value is set to NULL.
//   url	resource to access
//   flags	flags which control how the resource indicated by url is to be opened
//
// Returns
//   >= 0 in case of success, a negative value corresponding to an AVERROR code in case of failure
func IOOpen(url string, flags IOFlags) (*IOContext, error) {
	url_ := CString(url)
	defer FreeString(url_)

	var ctx *IOContext
	err := CheckErr(C.avio_open(&ctx, url_, C.int(flags)))
	return ctx, err
}

func (ctx *IOContext) Close() error {
	return CheckErr(C.avio_close(ctx))
}

func (ctx *IOContext) PutStr(s string) error {
	s_ := CString(s)
	defer FreeString(s_)

	return CheckErr(C.avio_put_str(ctx, s_))
}
