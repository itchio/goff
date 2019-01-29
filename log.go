package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/log.h>
//
// void goff_log_callback(void *ptr, int level, const char *fmt, va_list vl);
import "C"

type LogCallback func(line string)
type CLogCallback = *[0]byte

var logCallback LogCallback

func LogSetCallback(lc LogCallback) {
	logCallback = lc

	if logCallback == nil {
		C.av_log_set_callback(CLogCallback(C.av_log_default_callback))
	} else {
		C.av_log_set_callback(CLogCallback(C.goff_log_callback))
	}
}

//export goff_send_log_to_go
func goff_send_log_to_go(line *C.char) {
	if logCallback != nil {
		logCallback(C.GoString(line))
	}
}
