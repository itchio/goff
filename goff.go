package goff

//#cgo pkg-config: libavformat libavcodec libavutil libswscale x264
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/pixdesc.h>
//#include <libavutil/rational.h>
//#include <libswscale/swscale.h>
import "C"

import (
	"reflect"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

type (
	FormatContext   = C.struct_AVFormatContext
	InputFormat     = C.struct_AVInputFormat
	OutputFormat    = C.struct_AVOutputFormat
	Dictionary      = C.struct_AVDictionary
	Stream          = C.struct_AVStream
	Rational        = C.struct_AVRational
	CodecParameters = C.struct_AVCodecParameters
	Codec           = C.struct_AVCodec
	CodecContext    = C.struct_AVCodecContext
	Packet          = C.struct_AVPacket
	Frame           = C.struct_AVFrame

	ErrNum = C.int
	Timing = C.int64_t
)

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

var (
	TIME_BASE_Q Rational = C.AV_TIME_BASE_Q
)

type MediaType C.enum_AVMediaType

var (
	// Usually treated as AVMEDIA_TYPE_DATA.
	MediaType_Unknown MediaType = C.AVMEDIA_TYPE_UNKNOWN
	MediaType_Video   MediaType = C.AVMEDIA_TYPE_VIDEO
	MediaType_Audio   MediaType = C.AVMEDIA_TYPE_AUDIO
	// Opaque data information usually continuous.
	MediaType_Data     MediaType = C.AVMEDIA_TYPE_DATA
	MediaType_Subtitle MediaType = C.AVMEDIA_TYPE_SUBTITLE
	// Opaque data information usually sparse.
	MediaType_Attachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT
	MediaType_Nb         MediaType = C.AVMEDIA_TYPE_NB
)

func (mt MediaType) String() string {
	return C.GoString(C.av_get_media_type_string(mt.C()))
}

func (mt MediaType) C() C.enum_AVMediaType {
	return C.enum_AVMediaType(mt)
}

type PixelFormat C.enum_AVPixelFormat

// TODO: add missing pixel formats
var (
	PixelFormat_None PixelFormat = C.AV_PIX_FMT_NONE
	// planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	PixelFormat_YUV420P PixelFormat = C.AV_PIX_FMT_YUV420P
	// packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	PixelFormat_RGBA PixelFormat = C.AV_PIX_FMT_RGBA
)

type PixFmtDescriptor = C.struct_AVPixFmtDescriptor

// Returns a pixel format descriptor for provided pixel format or NULL if this pixel format is unknown.
func (pf PixelFormat) Desc() *PixFmtDescriptor {
	return C.av_pix_fmt_desc_get(pf.C())
}

// Return the short name for a pixel format, NULL in case pix_fmt is unknown.
func (pf PixelFormat) Name() string {
	return C.GoString(C.av_get_pix_fmt_name(pf.C()))
}

func (pf PixelFormat) String() string {
	return pf.Name()
}

func (pf PixelFormat) C() C.enum_AVPixelFormat {
	return C.enum_AVPixelFormat(pf)
}

type SampleFormat C.enum_AVSampleFormat

func (sf SampleFormat) C() C.enum_AVSampleFormat {
	return C.enum_AVSampleFormat(sf)
}

type CodecID C.enum_AVCodecID

func (cid CodecID) C() C.enum_AVCodecID {
	return C.enum_AVCodecID(cid)
}

// Find a registered decoder with a matching codec ID.
//
// Parameters
//   id AVCodecID of the requested decoder
//
// Returns
//   A decoder if one was found, NULL otherwise.
func (cid CodecID) FindDecoder() *Codec {
	return C.avcodec_find_decoder(cid.C())
}

// Find a registered encoder with a matching codec ID.
//
// Parameters
//   id	AVCodecID of the requested encoder
//
// Returns
//   An encoder if one was found, NULL otherwise.
func (cid CodecID) FindEncoder() *Codec {
	return C.avcodec_find_encoder(cid.C())
}

// Allocate an AVCodecContext and set its fields to default values.
//
// The resulting struct should be freed with avcodec_free_context().
//
// Parameters
//     codec if non-NULL, allocate private data and initialize defaults for the
//     given codec. It is illegal to then call avcodec_open2() with a different
//     codec. If NULL, then the codec-specific defaults won't be initialized, which
//     may result in suboptimal default settings (this is important mainly for
//     encoders, e.g. libx264).
//
// Returns
//     An AVCodecContext filled with default values or NULL on failure.
func (codec *Codec) AllocContext3() *CodecContext {
	return C.avcodec_alloc_context3(codec)
}

func (codec *Codec) String() string {
	return codec.Name()
}

func (codec *Codec) Name() string {
	return C.GoString(codec.name)
}

func (codec *Codec) LongName() string {
	return C.GoString(codec.long_name)
}

// Initialize the AVCodecContext to use the given AVCodec.
//
// Prior to using this function the context has to be allocated with
// avcodec_alloc_context3().
//
// The functions avcodec_find_decoder_by_name(), avcodec_find_encoder_by_name(),
// avcodec_find_decoder() and avcodec_find_encoder() provide an easy way for
// retrieving a codec.
//
// Warning: This function is not thread safe!
//
// Note: Always call this function before using decoding routines (such as
// avcodec_receive_frame()).
func (cctx *CodecContext) Open2(codec *Codec, options **Dictionary) error {
	return CheckErr(C.avcodec_open2(cctx, codec, options))
}

func (cctx *CodecContext) TimeBase() Rational {
	return cctx.time_base
}

// Supply raw packet data as input to a decoder.
//
// Internally, this call will copy relevant AVCodecContext fields, which can
// influence decoding per-packet, and apply them when the packet is actually
// decoded. (For example AVCodecContext.skip_frame, which might direct the
// decoder to drop the frame contained by the packet sent with this function.)
//
// Warning
//     The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE larger
//     than the actual read bytes because some optimized bitstream readers read 32
//     or 64 bits at once and could read over the end. Do not mix this API with the
//     legacy API (like avcodec_decode_video2()) on the same AVCodecContext. It will
//     return unexpected results now or in future libavcodec versions.
//
// Note
//     The AVCodecContext MUST have been opened with avcodec_open2() before packets
//     may be fed to the decoder.
func (cctx *CodecContext) SendPacket(pkt *Packet) error {
	return CheckErr(C.avcodec_send_packet(cctx, pkt))
}

// Return decoded output data from a decoder.
// Parameters
// avctx	codec context
// frame	This will be set to a reference-counted video or audio frame (depending
// on the decoder type) allocated by the decoder. Note that the function will
// always call av_frame_unref(frame) before doing anything else.
//
// Returns 0: success, a frame was returned AVERROR(EAGAIN): output is not
// available in this state - user must try to send new input AVERROR_EOF: the
// decoder has been fully flushed, and there will be no more output frames
// AVERROR(EINVAL): codec not opened, or it is an encoder other negative values:
// legitimate decoding errors
func (cctx *CodecContext) ReceiveFrame(frame *Frame) error {
	return CheckErr(C.avcodec_receive_frame(cctx, frame))
}

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

func (s *Stream) Index() int {
	return int(s.index)
}

func (s *Stream) ID() int {
	return int(s.id)
}

func (s *Stream) TimeBase() Rational {
	return s.time_base
}

func (s *Stream) StartTime() Timing {
	return s.start_time
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

func (cp *CodecParameters) CodecType() MediaType {
	return MediaType(cp.codec_type)
}

func (cp *CodecParameters) CodecID() CodecID {
	return CodecID(cp.codec_id)
}

func (cp *CodecParameters) PixelFormat() PixelFormat {
	return PixelFormat(cp.format)
}

func (cp *CodecParameters) SampleFormat() SampleFormat {
	return SampleFormat(cp.format)
}

// Fill the codec context based on the values from the supplied codec parameters.
//
// Any allocated fields in codec that have a corresponding field in par are
// freed and replaced with duplicates of the corresponding field in par. Fields
// in codec that do not have a counterpart in par are not touched.
//
// Returns
//     >= 0 on success, a negative AVERROR code on failure.
func (cp *CodecParameters) ToContext(cctx *CodecContext) error {
	return CheckErr(C.avcodec_parameters_to_context(cctx, cp))
}

func (t Timing) Rescale(bq Rational, cq Rational) Timing {
	return C.av_rescale_q(t, bq, cq)
}

func (t Timing) AsDuration(timebase Rational) time.Duration {
	return time.Duration(float64(time.Second) * float64(t) * timebase.Float())
}

func (r Rational) Float() float64 {
	return float64(r.num) / float64(r.den)
}

// Allocate an AVFrame and set its fields to default values.
//
// The resulting struct must be freed using av_frame_free().
//
// Returns
//     An AVFrame filled with default values or NULL on failure.
//
// Note
//     this only allocates the AVFrame itself, not the data buffers. Those must be
//     allocated through other means, e.g. with av_frame_get_buffer() or manually.
func FrameAlloc() *Frame {
	return C.av_frame_alloc()
}

// Free the frame and any dynamically allocated objects in it, e.g.
//
// extended_data. If the frame is reference counted, it will be unreferenced first.
func (frame *Frame) Free() {
	C.av_frame_free(&frame)
}

// Allocate new buffer(s) for audio or video data.
//
// The following fields must be set on frame before calling this function:
//
// format (pixel format for video, sample format for audio)
// width and height for video
// nb_samples and channel_layout for audio
//
// This function will fill AVFrame.data and AVFrame.buf arrays and, if
// necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
// For planar formats, one buffer will be allocated for each plane.
//
// Warning
// if frame already has been allocated, calling this function will leak
// memory. In addition, undefined behavior can occur in certain cases.
//
// Parameters
// frame	frame in which to store the new buffers.
// align	Required buffer size alignment. If equal to 0, alignment will be chosen
// automatically for the current CPU. It is highly recommended to pass 0 here
// unless you know what you are doing.
//
// Returns
// 0 on success, a negative AVERROR on error.
func (frame *Frame) GetBuffer(align int) error {
	return CheckErr(C.av_frame_get_buffer(frame, C.int(align)))
}

// Ensure that the frame data is writable, avoiding data copy if possible.
//
// Do nothing if the frame is writable, allocate new buffers and copy the data
// if it is not.
//
// Returns
//     0 on success, a negative AVERROR on error.
func (frame *Frame) MakeWritable() error {
	return CheckErr(C.av_frame_make_writable(frame))
}

// Presentation timestamp in time_base units (time when frame should be shown to user).
func (frame *Frame) Pts() Timing {
	return frame.pts
}

func (frame *Frame) Width() int {
	return int(frame.width)
}

func (frame *Frame) SetWidth(width int) {
	frame.width = C.int(width)
}

func (frame *Frame) Height() int {
	return int(frame.height)
}

func (frame *Frame) SetHeight(height int) {
	frame.height = C.int(height)
}

func (frame *Frame) Format() PixelFormat {
	return PixelFormat(frame.format)
}

func (frame *Frame) SetFormat(pf PixelFormat) {
	frame.format = C.int(pf)
}

type Plane = *C.uchar
type Planes = [8]Plane

func (frame *Frame) Data() Planes {
	return frame.data
}

func (frame *Frame) PlaneData(index int) []uint8 {
	linesize := frame.Linesize()[index]
	data := frame.Data()[index]

	var d []uint8
	dsize := int(linesize) * frame.Height()
	{
		sh := (*reflect.SliceHeader)((unsafe.Pointer(&d)))
		sh.Cap = dsize
		sh.Len = dsize
		sh.Data = uintptr(unsafe.Pointer(data))
	}
	return d
}

type Linesize = C.int
type Linesizes = [8]Linesize

func (frame *Frame) Linesize() Linesizes {
	return frame.linesize
}

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

// Presentation timestamp in AVStream->time_base units; the time at which the
// decompressed packet will be presented to the user
func (pkt *Packet) Pts() Timing {
	return pkt.pts
}

// Decompression timestamp in AVStream->time_base units; the time at which the
// packet is decompressed.
func (pkt *Packet) Dts() Timing {
	return pkt.dts
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

type SwsContext = C.struct_SwsContext
type SwsFilter = C.struct_SwsFilter

type SwsFlags = int

var (
	SwsFlags_Bicubic SwsFlags = C.SWS_BICUBIC
)

// Check if context can be reused, otherwise reallocate a new one.
// If context is NULL, just calls sws_getContext() to get a new context.
// Otherwise, checks if the parameters are the ones already saved in context. If
// that is the case, returns the current context. Otherwise, frees context and
// gets a new context with the new parameters.
//
// Be warned that srcFilter and dstFilter are not checked, they are assumed to
// remain the same.
func SwsGetCachedContext(
	context *SwsContext,
	srcW int,
	srcH int,
	srcFormat PixelFormat,
	dstW int,
	dstH int,
	dstFormat PixelFormat,
	flags SwsFlags,
	srcFilter *SwsFilter,
	dstFilter *SwsFilter,
	param *float64,
) *SwsContext {
	return C.sws_getCachedContext(
		context,
		C.int(srcW), C.int(srcH), srcFormat.C(),
		C.int(dstW), C.int(dstH), dstFormat.C(),
		C.int(flags), srcFilter, dstFilter,
		(*C.double)(param),
	)
}

func (swctx *SwsContext) Scale(
	srcSlice Planes, srcStride Linesizes,
	srcSliceY int, srcSliceH int,
	dst Planes, dstStride Linesizes,
) error {
	return CheckErr(C.sws_scale(
		swctx,
		(**C.uchar)(&srcSlice[0]),
		(*C.int)(&srcStride[0]),
		C.int(srcSliceY),
		C.int(srcSliceH),
		(**C.uchar)(&dst[0]),
		(*C.int)(&dstStride[0]),
	))
}

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

func CString(s string) *C.char {
	if s == "" {
		return (*C.char)(nil)
	}
	return C.CString(s)
}

func FreeString(c *C.char) {
	if c != nil {
		C.free(unsafe.Pointer(c))
	}
}

func BoolToInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
