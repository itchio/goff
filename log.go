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
// void goff_log_callback_trampoline(void *ptr, int level, const char *fmt, va_list vl);
import "C"

var log_callback LogCallback
var log_max_level LogLevel

type LogCallback func(level LogLevel, line string)
type CLogCallback = *[0]byte

func LogSetCallback(maxLevel LogLevel, lc LogCallback) {
	log_callback = lc
	log_max_level = maxLevel

	if log_callback == nil {
		C.av_log_set_callback(CLogCallback(C.av_log_default_callback))
	} else {
		C.av_log_set_callback(CLogCallback(C.goff_log_callback_trampoline))
	}
}

type LogLevel C.int

var (
	// Print no output
	LogLevel_QUIET LogLevel = C.AV_LOG_QUIET
	// Something went really wrong and we will crash now
	LogLevel_PANIC LogLevel = C.AV_LOG_PANIC
	// Something went wrong and recovery is not possible
	LogLevel_FATAL LogLevel = C.AV_LOG_FATAL
	// Something went wrong and cannot losslessly be recovered
	LogLevel_ERROR LogLevel = C.AV_LOG_ERROR
	// Something somehow does not look correct
	LogLevel_WARNING LogLevel = C.AV_LOG_WARNING
	// Standard information
	LogLevel_INFO LogLevel = C.AV_LOG_INFO
	// Detailed information
	LogLevel_VERBOSE LogLevel = C.AV_LOG_VERBOSE
	// Stuff which is only useful for libav* developrs
	LogLevel_DEBUG LogLevel = C.AV_LOG_DEBUG
	// Extremely verbose debugging, useful for libav* development
	LogLevel_TRACE LogLevel = C.AV_LOG_TRACE
)

func (ll LogLevel) String() string {
	switch ll {
	case LogLevel_PANIC:
		return "panic"
	case LogLevel_FATAL:
		return "fatal"
	case LogLevel_ERROR:
		return "error"
	case LogLevel_WARNING:
		return "warning"
	case LogLevel_INFO:
		return "info"
	case LogLevel_VERBOSE:
		return "verbose"
	case LogLevel_DEBUG:
		return "debug"
	case LogLevel_TRACE:
		return "trace"
	default:
		return ""
	}
}

//export goff_send_log_to_go
func goff_send_log_to_go(level C.int, line *C.char) {
	if log_callback != nil {
		log_callback(LogLevel(level), C.GoString(line))
	}
}

//export goff_should_send_log
func goff_should_send_log(level C.int) C.int {
	if log_callback == nil {
		return 0
	}
	if level > C.int(log_max_level) {
		return 0
	}
	return 1
}
