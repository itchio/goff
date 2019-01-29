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
import (
	"fmt"
	"time"
)

type Rational = C.struct_AVRational
type Timing = C.int64_t

var (
	// Internal time base represented as fractional value.
	TIME_BASE_Q Rational = C.AV_TIME_BASE_Q

	NOPTS_VALUE Timing = C.AV_NOPTS_VALUE
)

// Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
//
// Timestamps with unknown values (AV_NOPTS_VALUE) will be ignored.
func (t Timing) Rescale(bq Rational, cq Rational) Timing {
	if t.IsNop() {
		return t
	}
	return C.av_rescale_q(t, bq, cq)
}

func (t Timing) AsDuration(timebase Rational) time.Duration {
	return time.Duration(float64(time.Second) * float64(t) * timebase.Float())
}

func (t Timing) IsNop() bool {
	return t == NOPTS_VALUE
}

func (r Rational) Float() float64 {
	return float64(r.num) / float64(r.den)
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.num, r.den)
}

func NewRational(num int, den int) Rational {
	var r Rational
	r.num = C.int(num)
	r.den = C.int(den)
	return r
}

func (r Rational) Num() int {
	return int(r.num)
}

func (r Rational) Den() int {
	return int(r.den)
}
