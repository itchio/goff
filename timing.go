package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/rational.h>
import "C"
import "time"

//------------------------------------
// @Timing @Rational
//------------------------------------

type (
	Rational = C.struct_AVRational
	Timing   = C.int64_t
)

var (
	// Internal time base represented as fractional value.
	TIME_BASE_Q Rational = C.AV_TIME_BASE_Q
)

func (t Timing) Rescale(bq Rational, cq Rational) Timing {
	return C.av_rescale_q(t, bq, cq)
}

func (t Timing) AsDuration(timebase Rational) time.Duration {
	return time.Duration(float64(time.Second) * float64(t) * timebase.Float())
}

func (r Rational) Float() float64 {
	return float64(r.num) / float64(r.den)
}
