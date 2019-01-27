package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"unsafe"

	"github.com/pkg/errors"
)

type ErrNum = C.int

type Error struct {
	errnum ErrNum
}

func (e *Error) Error() string {
	buf := make([]byte, 1024)
	cstr := (*C.char)(unsafe.Pointer(&buf[0]))
	C.av_strerror(e.errnum, cstr, (C.size_t)(len(buf)))
	return C.GoString(cstr)
}

func IsErrNum(e error, errnum ErrNum) bool {
	if ee, ok := errors.Cause(e).(*Error); ok {
		return ee.errnum == errnum
	}
	return false
}

func IsEOF(e error) bool {
	return IsErrNum(e, ERROR_EOF)
}

func IsEAGAIN(e error) bool {
	return IsErrNum(e, ERROR_EAGAIN)
}

func CheckErr(errnum ErrNum) error {
	if errnum < 0 {
		return &Error{errnum: errnum}
	}
	return nil
}

var (
	ERROR_EAGAIN ErrNum = -C.EAGAIN
	// Bitstream filter not found.
	ERROR_BSF_NOT_FOUND ErrNum = C.AVERROR_BSF_NOT_FOUND
	// Internal bug, also see AVERROR_BUG2
	ERROR_BUG ErrNum = C.AVERROR_BUG
	// Buffer too small.
	ERROR_BUFFER_TOO_SMALL ErrNum = C.AVERROR_BUFFER_TOO_SMALL
	// Decoder not found
	ERROR_DECODER_NOT_FOUND ErrNum = C.AVERROR_DECODER_NOT_FOUND
	// Demuxer not found
	ERROR_DEMUXER_NOT_FOUND ErrNum = C.AVERROR_DEMUXER_NOT_FOUND
	// Encoder not found
	ERROR_ENCODER_NOT_FOUND ErrNum = C.AVERROR_ENCODER_NOT_FOUND
	// End of file
	ERROR_EOF ErrNum = C.AVERROR_EOF
	// Immediate exit was requested; the called function should not be restarted
	ERROR_EXIT ErrNum = C.AVERROR_EXIT
	// Generic error in an external library
	ERROR_EXTERNAL ErrNum = C.AVERROR_EXTERNAL
	// Filter not found
	ERROR_FILTER_NOT_FOUND ErrNum = C.AVERROR_FILTER_NOT_FOUND
	// Invalid data found when processing input
	ERROR_INVALIDDATA ErrNum = C.AVERROR_INVALIDDATA
	// Muxer not found
	ERROR_MUXER_NOT_FOUND ErrNum = C.AVERROR_MUXER_NOT_FOUND
	// Option not found
	ERROR_OPTION_NOT_FOUND ErrNum = C.AVERROR_OPTION_NOT_FOUND
	// Not yet implemented in FFmpeg, patches welcome.
	ERROR_PATCHWELCOME ErrNum = C.AVERROR_PATCHWELCOME
	// Protocol not found
	ERROR_PROTOCOL_NOT_FOUND ErrNum = C.AVERROR_PROTOCOL_NOT_FOUND
	// Stream not found
	ERROR_STREAM_NOT_FOUND ErrNum = C.AVERROR_STREAM_NOT_FOUND
	// This is semantically identical to AVERROR_BUG it has been introduced in Libav after our AVERROR_BUG and with a modified value.
	ERROR_BUG2 ErrNum = C.AVERROR_BUG2
	// Unknown error, typically from an external library.
	ERROR_UNKNOWN ErrNum = C.AVERROR_UNKNOWN
	// Requested feature is flagged experimental. Set strict_std_compliance if you really want to use it.
	ERROR_EXPERIMENTAL ErrNum = C.AVERROR_EXPERIMENTAL
	// Input changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_OUTPUT_CHANGED)
	ERROR_INPUT_CHANGED ErrNum = C.AVERROR_INPUT_CHANGED
	// Output changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_INPUT_CHANGED)
	ERROR_OUTPUT_CHANGED    ErrNum = C.AVERROR_OUTPUT_CHANGED
	ERROR_HTTP_BAD_REQUEST  ErrNum = C.AVERROR_HTTP_BAD_REQUEST
	ERROR_HTTP_UNAUTHORIZED ErrNum = C.AVERROR_HTTP_UNAUTHORIZED
	ERROR_HTTP_FORBIDDEN    ErrNum = C.AVERROR_HTTP_FORBIDDEN
	ERROR_HTTP_NOT_FOUND    ErrNum = C.AVERROR_HTTP_NOT_FOUND
	ERROR_HTTP_OTHER_4XX    ErrNum = C.AVERROR_HTTP_OTHER_4XX
	ERROR_HTTP_SERVER_ERROR ErrNum = C.AVERROR_HTTP_SERVER_ERROR
)
