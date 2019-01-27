package goff

//#cgo pkg-config: libavutil libavcodec
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
import "C"

type (
	Packet = C.struct_AVPacket
)

// Initialize optional fields of a packet with default values.
//
// Note, this does not touch the data and size members, which have to be
// initialized separately.
func (pkt *Packet) Init() {
	C.av_init_packet(pkt)
}

func (pkt *Packet) StreamIndex() int {
	return int(pkt.stream_index)
}

func (pkt *Packet) SetStreamIndex(i int) {
	pkt.stream_index = C.int(i)
}

// Presentation timestamp in AVStream->time_base units; the time at which the
// decompressed packet will be presented to the user
func (pkt *Packet) Pts() Timing {
	return pkt.pts
}

func (pkt *Packet) SetPts(pts Timing) {
	pkt.pts = pts
}

// Decompression timestamp in AVStream->time_base units; the time at which the
// packet is decompressed.
func (pkt *Packet) Dts() Timing {
	return pkt.dts
}

func (pkt *Packet) SetDts(dts Timing) {
	pkt.dts = dts
}

// Duration of this packet in AVStream->time_base units, 0 if unknown.
func (pkt *Packet) Duration() Timing {
	return pkt.duration
}

func (pkt *Packet) Size() int {
	return int(pkt.size)
}

// byte position in stream, -1 if unknown
func (pkt *Packet) Pos() int64 {
	return int64(pkt.pos)
}

// Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
//
// Timestamps with unknown values (AV_NOPTS_VALUE) will be ignored.
//
// Parameters
//     pkt	packet on which the conversion will be performed
//     tb_src	source timebase, in which the timing fields in pkt are expressed
//     tb_dst	destination timebase, to which the timing fields will be converted
func (pkt *Packet) RescaleTs(tbSrc Rational, tbDst Rational) {
	C.av_packet_rescale_ts(pkt, tbSrc, tbDst)
}
