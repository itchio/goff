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
import (
	"reflect"
	"unsafe"
)

type FormatContext = C.struct_AVFormatContext
type InputFormat = C.struct_AVInputFormat
type OutputFormat = C.struct_AVOutputFormat

// Open an input stream and read the header.
//
// The codecs are not opened. The stream must be closed with avformat_close_input().
func FormatOpenInput(url string, format *InputFormat, options **Dictionary) (*FormatContext, error) {
	var ctx *FormatContext

	url_ := CString(url)
	defer FreeString(url_)

	ret := C.avformat_open_input(&ctx, url_, format, options)
	return ctx, CheckErr(ret)
}

func FormatAllocOutputContext2(oformat *OutputFormat, formatName string, filename string) (*FormatContext, error) {
	var ctx *FormatContext

	formatName_ := CString(formatName)
	defer FreeString(formatName_)
	filename_ := CString(filename)
	defer FreeString(filename_)

	ret := C.avformat_alloc_output_context2(&ctx, oformat, formatName_, filename_)
	return ctx, CheckErr(ret)
}

func (ctx *FormatContext) PB() *IOContext {
	return ctx.pb
}

func (ctx *FormatContext) SetPB(pb *IOContext) {
	ctx.pb = pb
}

func (ctx *FormatContext) Free() {
	C.avformat_free_context(ctx)
}

// Read packets of a media file to get stream information.
//
// This is useful for file formats with no headers such as MPEG. This function
// also computes the real framerate in case of MPEG-2 repeat frame mode. The
// logical file position is not changed by this function; examined packets may
// be buffered for later processing.
func (ctx *FormatContext) FindStreamInfo(options **Dictionary) error {
	return CheckErr(C.avformat_find_stream_info(ctx, options))
}

// Number of elements in AVFormatContext.streams.
func (ctx *FormatContext) NbStreams() int {
	return int(ctx.nb_streams)
}

// Position of the first frame of the component, in AV_TIME_BASE fractional seconds.
func (ctx *FormatContext) StartTime() Timing {
	return ctx.start_time
}

// A list of all streams in the file.
func (ctx *FormatContext) Streams() []*Stream {
	var streams []*Stream
	{
		sh := (*reflect.SliceHeader)((unsafe.Pointer(&streams)))
		sh.Cap = int(ctx.nb_streams)
		sh.Len = int(ctx.nb_streams)
		sh.Data = uintptr(unsafe.Pointer(ctx.streams))
	}
	return streams
}

func (ctx *FormatContext) DumpFormat(index int, url string, isOutput bool) {
	url_ := CString(url)
	defer FreeString(url_)

	C.av_dump_format(ctx, C.int(index), url_, BoolToInt(isOutput))
}

// Return the next frame of a stream.
//
// This function returns what is stored in the file, and does not validate that
// what is there are valid frames for the decoder. It will split what is stored
// in the file into frames and return one for each call. It will not omit
// invalid data between valid frames so as to give the decoder the maximum
// information possible for decoding.
//
// If pkt->buf is NULL, then the packet is valid until the next av_read_frame()
// or until avformat_close_input(). Otherwise the packet is valid indefinitely.
// In both cases the packet must be freed with av_packet_unref when it is no
// longer needed. For video, the packet contains exactly one frame. For audio,
// it contains an integer number of frames if each frame has a known fixed size
// (e.g. PCM or ADPCM data). If the audio frames have a variable size (e.g. MPEG
// audio), then it contains one frame.
//
// pkt->pts, pkt->dts and pkt->duration are always set to correct values in
// AVStream.time_base units (and guessed if the format cannot provide them).
// pkt->pts can be AV_NOPTS_VALUE if the video format has B-frames, so it is
// better to rely on pkt->dts if you do not decompress the payload.
//
// Returns
//     0 if OK, < 0 on error or end of file
func (ctx *FormatContext) ReadFrame(pkt *Packet) error {
	return CheckErr(C.av_read_frame(ctx, pkt))
}

// Write a packet to an output media file ensuring correct interleaving.
//
// This function will buffer the packets internally as needed to make sure the
// packets in the output file are properly interleaved in the order of
// increasing dts. Callers doing their own interleaving should call
// av_write_frame() instead of this function.
//
// Using this function instead of av_write_frame() can give muxers advance
// knowledge of future packets, improving e.g. the behaviour of the mp4 muxer
// for VFR content in fragmenting mode.
//
// Parameters
//     s media file handle
//     pkt The packet containing the data to be written.
//     If the packet is reference-counted, this function will take ownership of this
//     reference and unreference it later when it sees fit. The caller must not
//     access the data through this reference after this function returns. If the
//     packet is not reference-counted, libavformat will make a copy. This parameter
//     can be NULL (at any time, not just at the end), to flush the interleaving
//     queues. Packet's stream_index field must be set to the index of the
//     corresponding stream in s->streams. The timestamps (pts, dts) must be set to
//     correct values in the stream's timebase (unless the output format is flagged
//     with the AVFMT_NOTIMESTAMPS flag, then they can be set to AV_NOPTS_VALUE).
//     The dts for subsequent packets in one stream must be strictly increasing
//     (unless the output format is flagged with the AVFMT_TS_NONSTRICT, then they
//     merely have to be nondecreasing). duration) should also be set if known.
//
// Returns
//     0 on success, a negative AVERROR on error. Libavformat will always take care of freeing the packet, even if this function fails.
func (ctx *FormatContext) InterleavedWriteFrame(pkt *Packet) error {
	return CheckErr(C.av_interleaved_write_frame(ctx, pkt))
}

// Add a new stream to a media file.
//
// When demuxing, it is called by the demuxer in read_header(). If the flag
// AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also be called in
// read_packet().
//
// When muxing, should be called by the user before avformat_write_header().
//
// User is required to call avcodec_close() and avformat_free_context() to clean
// up the allocation by avformat_new_stream().
func (ctx *FormatContext) NewStream(c *Codec) *Stream {
	return C.avformat_new_stream(ctx, c)
}

func (ctx *FormatContext) WriteTrailer() error {
	return CheckErr(C.av_write_trailer(ctx))
}

// Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.

// Parameters
//   short_name if non-NULL checks if short_name matches with the names of the
//   registered formats
//   filename if non-NULL checks if filename terminates with the extensions of the
//   registered formats
//   mime_type if non-NULL checks if mime_type matches with the MIME type of the
//   registered formats
func GuessFormat(shortName string, filename string, mimeType string) *OutputFormat {
	shortName_ := CString(shortName)
	defer FreeString(shortName_)

	filename_ := CString(filename)
	defer FreeString(filename_)

	mimeType_ := CString(mimeType)
	defer FreeString(mimeType_)

	return C.av_guess_format(shortName_, filename_, mimeType_)
}

func (of *OutputFormat) Name() string {
	return C.GoString(of.name)
}

func (of *OutputFormat) LongName() string {
	return C.GoString(of.long_name)
}

func (of *OutputFormat) MimeType() string {
	return C.GoString(of.mime_type)
}

func (of *OutputFormat) Extensions() string {
	return C.GoString(of.extensions)
}

func (of *OutputFormat) AudioCodec() CodecID {
	return CodecID(of.audio_codec)
}

func (of *OutputFormat) VideoCodec() CodecID {
	return CodecID(of.video_codec)
}
