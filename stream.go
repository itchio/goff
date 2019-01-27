package goff

//#cgo pkg-config: libavutil libavformat
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavformat/avformat.h>
import "C"

type Stream = C.struct_AVStream

func (s *Stream) Index() int {
	return int(s.index)
}

func (s *Stream) ID() int {
	return int(s.id)
}

func (s *Stream) SetID(id int) {
	s.id = C.int(id)
}

func (s *Stream) TimeBase() Rational {
	return s.time_base
}

func (s *Stream) SetTimeBase(tb Rational) {
	s.time_base = tb
}

func (s *Stream) StartTime() Timing {
	return s.start_time
}

func (s *Stream) SetStartTime(st Timing) {
	s.start_time = st
}

func (s *Stream) Duration() Timing {
	return s.duration
}

func (s *Stream) NbFrames() int64 {
	return int64(s.nb_frames)
}

func (s *Stream) CodecParameters() *CodecParameters {
	return s.codecpar
}

func (s *Stream) SetCodecParameters(cp *CodecParameters) {
	s.codecpar = cp
}
