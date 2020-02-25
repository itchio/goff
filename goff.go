package goff

//#cgo pkg-config: libavutil libavformat libavcodec libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//
//#include <libavformat/avformat.h>
//#include <libavformat/avio.h>
//
//#include <libavutil/avutil.h>
//#include <libavutil/rational.h>
//#include <libavutil/opt.h>
//#include <libavutil/frame.h>
//#include <libavutil/pixdesc.h>
//#include <libavutil/samplefmt.h>
//#include <libavutil/log.h>
//
//#include <libavcodec/avcodec.h>
//
//#include <libswscale/swscale.h>
//
// void goff_log_callback_trampoline(void *ptr, int level, const char *fmt, va_list vl);
// int goff_reader_read_packet_trampoline(void *opaque, uint8_t *buf, int buf_size);
// int64_t goff_reader_seek_trampoline(void *opaque, int64_t offset, int whence);
// int goff_writer_write_packet_trampoline(void *opaque, uint8_t *buf, int buf_size);
// int64_t goff_writer_seek_trampoline(void *opaque, int64_t offset, int whence);
//
import "C"
import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

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

type CodecContext = C.struct_AVCodecContext

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

// Free the codec context and everything associated with it and write NULL to
// the provided pointer.
func (cctx *CodecContext) Free() {
	C.avcodec_free_context(&cctx)
}

func (cctx *CodecContext) CodecID() CodecID {
	return CodecID(cctx.codec_id)
}

func (cctx *CodecContext) SetCodecID(cid CodecID) {
	cctx.codec_id = uint32(cid)
}

func (cctx *CodecContext) CodecType() MediaType {
	return MediaType(cctx.codec_type)
}

func (cctx *CodecContext) SetCodecType(mt MediaType) {
	cctx.codec_type = int32(mt)
}

func (cctx *CodecContext) ThreadCount() int {
	return int(cctx.thread_count)
}

func (cctx *CodecContext) SetThreadCount(tc int) {
	cctx.thread_count = C.int(tc)
}

type ThreadType = C.int

var (
	ThreadType_UNSET ThreadType = 0
	ThreadType_FRAME ThreadType = C.FF_THREAD_FRAME
	ThreadType_SLICE ThreadType = C.FF_THREAD_SLICE
)

func (cctx *CodecContext) ThreadType() ThreadType {
	return ThreadType(cctx.thread_type)
}

func (cctx *CodecContext) SetThreadType(tt ThreadType) {
	cctx.thread_type = C.int(tt)
}

func (cctx *CodecContext) PixelFormat() PixelFormat {
	return PixelFormat(cctx.pix_fmt)
}

func (cctx *CodecContext) SetPixelFormat(pf PixelFormat) {
	cctx.pix_fmt = int32(pf)
}

func (cctx *CodecContext) TimeBase() Rational {
	return cctx.time_base
}

func (cctx *CodecContext) SetTimeBase(tb Rational) {
	cctx.time_base = tb
}

func (cctx *CodecContext) Width() int {
	return int(cctx.width)
}

func (cctx *CodecContext) SetWidth(x int) {
	cctx.width = C.int(x)
}

func (cctx *CodecContext) Height() int {
	return int(cctx.height)
}

func (cctx *CodecContext) SetHeight(x int) {
	cctx.height = C.int(x)
}

func (cctx *CodecContext) GOPSize() int {
	return int(cctx.gop_size)
}

func (cctx *CodecContext) SetGOPSize(x int) {
	cctx.gop_size = C.int(x)
}

func (cctx *CodecContext) QMin() int {
	return int(cctx.qmin)
}

func (cctx *CodecContext) SetQMin(x int) {
	cctx.qmin = C.int(x)
}

func (cctx *CodecContext) QMax() int {
	return int(cctx.qmax)
}

func (cctx *CodecContext) SetQMax(x int) {
	cctx.qmax = C.int(x)
}

func (cctx *CodecContext) MaxBFrames() int {
	return int(cctx.max_b_frames)
}

func (cctx *CodecContext) SetMaxBFrames(x int) {
	cctx.max_b_frames = C.int(x)
}

type CodecFlags = C.int

var (
	CodecFlags_UNALIGNED      CodecFlags = C.AV_CODEC_FLAG_UNALIGNED
	CodecFlags_QSCALE         CodecFlags = C.AV_CODEC_FLAG_QSCALE
	CodecFlags_4MV            CodecFlags = C.AV_CODEC_FLAG_4MV
	CodecFlags_OUTPUT_CORRUPT CodecFlags = C.AV_CODEC_FLAG_OUTPUT_CORRUPT
	CodecFlags_QPEL           CodecFlags = C.AV_CODEC_FLAG_QPEL
	CodecFlags_PASS1          CodecFlags = C.AV_CODEC_FLAG_PASS1
	CodecFlags_PASS2          CodecFlags = C.AV_CODEC_FLAG_PASS2
	CodecFlags_LOOP_FILTER    CodecFlags = C.AV_CODEC_FLAG_LOOP_FILTER
	CodecFlags_GRAY           CodecFlags = C.AV_CODEC_FLAG_GRAY
	CodecFlags_PSNR           CodecFlags = C.AV_CODEC_FLAG_PSNR
	CodecFlags_TRUNCATED      CodecFlags = C.AV_CODEC_FLAG_TRUNCATED
	CodecFlags_INTERLACED_DCT CodecFlags = C.AV_CODEC_FLAG_INTERLACED_DCT
	CodecFlags_LOW_DELAY      CodecFlags = C.AV_CODEC_FLAG_LOW_DELAY
	CodecFlags_GLOBAL_HEADER  CodecFlags = C.AV_CODEC_FLAG_GLOBAL_HEADER
	CodecFlags_BITEXACT       CodecFlags = C.AV_CODEC_FLAG_BITEXACT
	CodecFlags_AC_PRED        CodecFlags = C.AV_CODEC_FLAG_AC_PRED
	CodecFlags_INTERLACED_ME  CodecFlags = C.AV_CODEC_FLAG_INTERLACED_ME
	// FIXME: this overflows int?
	// CodecFlags_CLOSED_GOP     CodecFlags = C.AV_CODEC_FLAG_CLOSED_GOP
)

func (cctx *CodecContext) Flags() CodecFlags {
	return CodecFlags(cctx.flags)
}

func (cctx *CodecContext) SetFlags(x CodecFlags) {
	cctx.flags = C.int(x)
}

func (cctx *CodecContext) OrFlag(x CodecFlags) {
	cctx.flags |= C.int(x)
}

func (cctx *CodecContext) Profile() Profile {
	return Profile(cctx.profile)
}

func (cctx *CodecContext) SetProfile(p Profile) {
	cctx.profile = C.int(p)
}

func (cctx *CodecContext) PrivData() unsafe.Pointer {
	return cctx.priv_data
}

// Read encoded data from the encoder.
//
// Parameters
//     avctx	codec context
//     avpkt This will be set to a reference-counted packet allocated by the
//     encoder. Note that the function will always call av_frame_unref(frame) before
//     doing anything else.
//
// Returns
//     0 on success, otherwise negative error code: AVERROR(EAGAIN): output is not
//     available in the current state - user must try to send input AVERROR_EOF: the
//     encoder has been fully flushed, and there will be no more output packets
//     AVERROR(EINVAL): codec not opened, or it is an encoder other errors:
//     legitimate decoding errors
func (cctx *CodecContext) ReceivePacket(pkt *Packet) error {
	return CheckErr(C.avcodec_receive_packet(cctx, pkt))
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

// Supply a raw video or audio frame to the encoder.
//
// Use avcodec_receive_packet() to retrieve buffered output packets.
//
// Parameters
//     	avctx	codec context
//     [in] frame AVFrame containing the raw audio or video frame to be encoded.
//     Ownership of the frame remains with the caller, and the encoder will not
//     write to the frame. The encoder may create a reference to the frame data (or
//     copy it if the frame is not reference-counted). It can be NULL, in which case
//     it is considered a flush packet. This signals the end of the stream. If the
//     encoder still has packets buffered, it will return them after this call. Once
//     flushing mode has been entered, additional flush packets are ignored, and
//     sending frames will return AVERROR_EOF.
//
// For audio: If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame can
// have any number of samples. If it is not set, frame->nb_samples must be equal
// to avctx->frame_size for all frames except the last. The final frame may be
// smaller than avctx->frame_size.
//
// Returns
//     0 on success, otherwise negative error code: AVERROR(EAGAIN): input is not accepted in the current state - user must read output with avcodec_receive_packet() (once all output is read, the packet should be resent, and the call will not fail with EAGAIN). AVERROR_EOF: the encoder has been flushed, and no new frames can be sent to it AVERROR(EINVAL): codec not opened, refcounted_frames not set, it is a decoder, or requires flush AVERROR(ENOMEM): failed to add packet to internal queue, or similar other errors: legitimate decoding errors
func (cctx *CodecContext) SendFrame(frame *Frame) error {
	return CheckErr(C.avcodec_send_frame(cctx, frame))
}

type Profile = C.int

var (
	Profile_AAC_MAIN      Profile = C.FF_PROFILE_AAC_MAIN
	Profile_AAC_LOW       Profile = C.FF_PROFILE_AAC_LOW
	Profile_AAC_SSR       Profile = C.FF_PROFILE_AAC_SSR
	Profile_AAC_LTP       Profile = C.FF_PROFILE_AAC_LTP
	Profile_AAC_HE        Profile = C.FF_PROFILE_AAC_HE
	Profile_AAC_HE_V2     Profile = C.FF_PROFILE_AAC_HE_V2
	Profile_AAC_ELD       Profile = C.FF_PROFILE_AAC_ELD
	Profile_MPEG2_AAC_LOW Profile = C.FF_PROFILE_MPEG2_AAC_LOW
	Profile_MPEG2_AAC_HE  Profile = C.FF_PROFILE_MPEG2_AAC_HE

	Profile_DNXHD     Profile = C.FF_PROFILE_DNXHD
	Profile_DNXHR_LB  Profile = C.FF_PROFILE_DNXHR_LB
	Profile_DNXHR_SQ  Profile = C.FF_PROFILE_DNXHR_SQ
	Profile_DNXHR_HQ  Profile = C.FF_PROFILE_DNXHR_HQ
	Profile_DNXHR_HQX Profile = C.FF_PROFILE_DNXHR_HQX
	Profile_DNXHR_444 Profile = C.FF_PROFILE_DNXHR_444

	Profile_DTS         Profile = C.FF_PROFILE_DTS
	Profile_DTS_ES      Profile = C.FF_PROFILE_DTS_ES
	Profile_DTS_96_24   Profile = C.FF_PROFILE_DTS_96_24
	Profile_DTS_HD_HRA  Profile = C.FF_PROFILE_DTS_HD_HRA
	Profile_DTS_HD_MA   Profile = C.FF_PROFILE_DTS_HD_MA
	Profile_DTS_EXPRESS Profile = C.FF_PROFILE_DTS_EXPRESS

	Profile_MPEG2_422          Profile = C.FF_PROFILE_MPEG2_422
	Profile_MPEG2_HIGH         Profile = C.FF_PROFILE_MPEG2_HIGH
	Profile_MPEG2_SS           Profile = C.FF_PROFILE_MPEG2_SS
	Profile_MPEG2_SNR_SCALABLE Profile = C.FF_PROFILE_MPEG2_SNR_SCALABLE
	Profile_MPEG2_MAIN         Profile = C.FF_PROFILE_MPEG2_MAIN
	Profile_MPEG2_SIMPLE       Profile = C.FF_PROFILE_MPEG2_SIMPLE

	Profile_H264_CONSTRAINED Profile = C.FF_PROFILE_H264_CONSTRAINED
	Profile_H264_INTRA       Profile = C.FF_PROFILE_H264_INTRA

	Profile_H264_BASELINE             Profile = C.FF_PROFILE_H264_BASELINE
	Profile_H264_CONSTRAINED_BASELINE Profile = C.FF_PROFILE_H264_CONSTRAINED_BASELINE
	Profile_H264_MAIN                 Profile = C.FF_PROFILE_H264_MAIN
	Profile_H264_EXTENDED             Profile = C.FF_PROFILE_H264_EXTENDED
	Profile_H264_HIGH                 Profile = C.FF_PROFILE_H264_HIGH
	Profile_H264_HIGH_10              Profile = C.FF_PROFILE_H264_HIGH_10
	Profile_H264_HIGH_10_INTRA        Profile = C.FF_PROFILE_H264_HIGH_10_INTRA
	Profile_H264_MULTIVIEW_HIGH       Profile = C.FF_PROFILE_H264_MULTIVIEW_HIGH
	Profile_H264_HIGH_422             Profile = C.FF_PROFILE_H264_HIGH_422
	Profile_H264_HIGH_422_INTRA       Profile = C.FF_PROFILE_H264_HIGH_422_INTRA
	Profile_H264_STEREO_HIGH          Profile = C.FF_PROFILE_H264_STEREO_HIGH
	Profile_H264_HIGH_444             Profile = C.FF_PROFILE_H264_HIGH_444
	Profile_H264_HIGH_444_PREDICTIVE  Profile = C.FF_PROFILE_H264_HIGH_444_PREDICTIVE
	Profile_H264_HIGH_444_INTRA       Profile = C.FF_PROFILE_H264_HIGH_444_INTRA
	Profile_H264_CAVLC_444            Profile = C.FF_PROFILE_H264_CAVLC_444

	Profile_VC1_SIMPLE   Profile = C.FF_PROFILE_VC1_SIMPLE
	Profile_VC1_MAIN     Profile = C.FF_PROFILE_VC1_MAIN
	Profile_VC1_COMPLEX  Profile = C.FF_PROFILE_VC1_COMPLEX
	Profile_VC1_ADVANCED Profile = C.FF_PROFILE_VC1_ADVANCED

	Profile_MPEG4_SIMPLE                    Profile = C.FF_PROFILE_MPEG4_SIMPLE
	Profile_MPEG4_SIMPLE_SCALABLE           Profile = C.FF_PROFILE_MPEG4_SIMPLE_SCALABLE
	Profile_MPEG4_CORE                      Profile = C.FF_PROFILE_MPEG4_CORE
	Profile_MPEG4_MAIN                      Profile = C.FF_PROFILE_MPEG4_MAIN
	Profile_MPEG4_N_BIT                     Profile = C.FF_PROFILE_MPEG4_N_BIT
	Profile_MPEG4_SCALABLE_TEXTURE          Profile = C.FF_PROFILE_MPEG4_SCALABLE_TEXTURE
	Profile_MPEG4_SIMPLE_FACE_ANIMATION     Profile = C.FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION
	Profile_MPEG4_BASIC_ANIMATED_TEXTURE    Profile = C.FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE
	Profile_MPEG4_HYBRID                    Profile = C.FF_PROFILE_MPEG4_HYBRID
	Profile_MPEG4_ADVANCED_REAL_TIME        Profile = C.FF_PROFILE_MPEG4_ADVANCED_REAL_TIME
	Profile_MPEG4_CORE_SCALABLE             Profile = C.FF_PROFILE_MPEG4_CORE_SCALABLE
	Profile_MPEG4_ADVANCED_CODING           Profile = C.FF_PROFILE_MPEG4_ADVANCED_CODING
	Profile_MPEG4_ADVANCED_CORE             Profile = C.FF_PROFILE_MPEG4_ADVANCED_CORE
	Profile_MPEG4_ADVANCED_SCALABLE_TEXTURE Profile = C.FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE
	Profile_MPEG4_SIMPLE_STUDIO             Profile = C.FF_PROFILE_MPEG4_SIMPLE_STUDIO
	Profile_MPEG4_ADVANCED_SIMPLE           Profile = C.FF_PROFILE_MPEG4_ADVANCED_SIMPLE

	Profile_JPEG2000_CSTREAM_RESTRICTION_0  Profile = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0
	Profile_JPEG2000_CSTREAM_RESTRICTION_1  Profile = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1
	Profile_JPEG2000_CSTREAM_NO_RESTRICTION Profile = C.FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION
	Profile_JPEG2000_DCINEMA_2K             Profile = C.FF_PROFILE_JPEG2000_DCINEMA_2K
	Profile_JPEG2000_DCINEMA_4K             Profile = C.FF_PROFILE_JPEG2000_DCINEMA_4K

	Profile_VP9_0 Profile = C.FF_PROFILE_VP9_0
	Profile_VP9_1 Profile = C.FF_PROFILE_VP9_1
	Profile_VP9_2 Profile = C.FF_PROFILE_VP9_2
	Profile_VP9_3 Profile = C.FF_PROFILE_VP9_3

	Profile_HEVC_MAIN               Profile = C.FF_PROFILE_HEVC_MAIN
	Profile_HEVC_MAIN_10            Profile = C.FF_PROFILE_HEVC_MAIN_10
	Profile_HEVC_MAIN_STILL_PICTURE Profile = C.FF_PROFILE_HEVC_MAIN_STILL_PICTURE
	Profile_HEVC_REXT               Profile = C.FF_PROFILE_HEVC_REXT

	Profile_MJPEG_HUFFMAN_BASELINE_DCT            Profile = C.FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT
	Profile_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT Profile = C.FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT
	Profile_MJPEG_HUFFMAN_PROGRESSIVE_DCT         Profile = C.FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT
	Profile_MJPEG_HUFFMAN_LOSSLESS                Profile = C.FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS
	Profile_MJPEG_JPEG_LS                         Profile = C.FF_PROFILE_MJPEG_JPEG_LS

	Profile_SBC_MSBC Profile = C.FF_PROFILE_SBC_MSBC
)

type Codec = C.struct_AVCodec

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

func (codec *Codec) ID() CodecID {
	return CodecID(codec.id)
}

func (codec *Codec) Type() MediaType {
	return MediaType(codec._type)
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

var (
	//--------------------------------------
	// Video codecs
	//--------------------------------------

	CodecID_NONE            CodecID = C.AV_CODEC_ID_NONE
	CodecID_MPEG1VIDEO      CodecID = C.AV_CODEC_ID_MPEG1VIDEO
	CodecID_MPEG2VIDEO      CodecID = C.AV_CODEC_ID_MPEG2VIDEO
	CodecID_H261            CodecID = C.AV_CODEC_ID_H261
	CodecID_H263            CodecID = C.AV_CODEC_ID_H263
	CodecID_RV10            CodecID = C.AV_CODEC_ID_RV10
	CodecID_RV20            CodecID = C.AV_CODEC_ID_RV20
	CodecID_MJPEG           CodecID = C.AV_CODEC_ID_MJPEG
	CodecID_MJPEGB          CodecID = C.AV_CODEC_ID_MJPEGB
	CodecID_LJPEG           CodecID = C.AV_CODEC_ID_LJPEG
	CodecID_SP5X            CodecID = C.AV_CODEC_ID_SP5X
	CodecID_JPEGLS          CodecID = C.AV_CODEC_ID_JPEGLS
	CodecID_MPEG4           CodecID = C.AV_CODEC_ID_MPEG4
	CodecID_RAWVIDEO        CodecID = C.AV_CODEC_ID_RAWVIDEO
	CodecID_MSMPEG4V1       CodecID = C.AV_CODEC_ID_MSMPEG4V1
	CodecID_MSMPEG4V2       CodecID = C.AV_CODEC_ID_MSMPEG4V2
	CodecID_MSMPEG4V3       CodecID = C.AV_CODEC_ID_MSMPEG4V3
	CodecID_WMV1            CodecID = C.AV_CODEC_ID_WMV1
	CodecID_WMV2            CodecID = C.AV_CODEC_ID_WMV2
	CodecID_H263P           CodecID = C.AV_CODEC_ID_H263P
	CodecID_H263I           CodecID = C.AV_CODEC_ID_H263I
	CodecID_FLV1            CodecID = C.AV_CODEC_ID_FLV1
	CodecID_SVQ1            CodecID = C.AV_CODEC_ID_SVQ1
	CodecID_SVQ3            CodecID = C.AV_CODEC_ID_SVQ3
	CodecID_DVVIDEO         CodecID = C.AV_CODEC_ID_DVVIDEO
	CodecID_HUFFYUV         CodecID = C.AV_CODEC_ID_HUFFYUV
	CodecID_CYUV            CodecID = C.AV_CODEC_ID_CYUV
	CodecID_H264            CodecID = C.AV_CODEC_ID_H264
	CodecID_INDEO3          CodecID = C.AV_CODEC_ID_INDEO3
	CodecID_VP3             CodecID = C.AV_CODEC_ID_VP3
	CodecID_THEORA          CodecID = C.AV_CODEC_ID_THEORA
	CodecID_ASV1            CodecID = C.AV_CODEC_ID_ASV1
	CodecID_ASV2            CodecID = C.AV_CODEC_ID_ASV2
	CodecID_FFV1            CodecID = C.AV_CODEC_ID_FFV1
	CodecID_4XM             CodecID = C.AV_CODEC_ID_4XM
	CodecID_VCR1            CodecID = C.AV_CODEC_ID_VCR1
	CodecID_CLJR            CodecID = C.AV_CODEC_ID_CLJR
	CodecID_MDEC            CodecID = C.AV_CODEC_ID_MDEC
	CodecID_ROQ             CodecID = C.AV_CODEC_ID_ROQ
	CodecID_INTERPLAY_VIDEO CodecID = C.AV_CODEC_ID_INTERPLAY_VIDEO
	CodecID_XAN_WC3         CodecID = C.AV_CODEC_ID_XAN_WC3
	CodecID_XAN_WC4         CodecID = C.AV_CODEC_ID_XAN_WC4
	CodecID_RPZA            CodecID = C.AV_CODEC_ID_RPZA
	CodecID_CINEPAK         CodecID = C.AV_CODEC_ID_CINEPAK
	CodecID_WS_VQA          CodecID = C.AV_CODEC_ID_WS_VQA
	CodecID_MSRLE           CodecID = C.AV_CODEC_ID_MSRLE
	CodecID_MSVIDEO1        CodecID = C.AV_CODEC_ID_MSVIDEO1
	CodecID_IDCIN           CodecID = C.AV_CODEC_ID_IDCIN
	CodecID_8BPS            CodecID = C.AV_CODEC_ID_8BPS
	CodecID_SMC             CodecID = C.AV_CODEC_ID_SMC
	CodecID_FLIC            CodecID = C.AV_CODEC_ID_FLIC
	CodecID_TRUEMOTION1     CodecID = C.AV_CODEC_ID_TRUEMOTION1
	CodecID_VMDVIDEO        CodecID = C.AV_CODEC_ID_VMDVIDEO
	CodecID_MSZH            CodecID = C.AV_CODEC_ID_MSZH
	CodecID_ZLIB            CodecID = C.AV_CODEC_ID_ZLIB
	CodecID_QTRLE           CodecID = C.AV_CODEC_ID_QTRLE
	CodecID_TSCC            CodecID = C.AV_CODEC_ID_TSCC
	CodecID_ULTI            CodecID = C.AV_CODEC_ID_ULTI
	CodecID_QDRAW           CodecID = C.AV_CODEC_ID_QDRAW
	CodecID_VIXL            CodecID = C.AV_CODEC_ID_VIXL
	CodecID_QPEG            CodecID = C.AV_CODEC_ID_QPEG
	CodecID_PNG             CodecID = C.AV_CODEC_ID_PNG
	CodecID_PPM             CodecID = C.AV_CODEC_ID_PPM
	CodecID_PBM             CodecID = C.AV_CODEC_ID_PBM
	CodecID_PGM             CodecID = C.AV_CODEC_ID_PGM
	CodecID_PGMYUV          CodecID = C.AV_CODEC_ID_PGMYUV
	CodecID_PAM             CodecID = C.AV_CODEC_ID_PAM
	CodecID_FFVHUFF         CodecID = C.AV_CODEC_ID_FFVHUFF
	CodecID_RV30            CodecID = C.AV_CODEC_ID_RV30
	CodecID_RV40            CodecID = C.AV_CODEC_ID_RV40
	CodecID_VC1             CodecID = C.AV_CODEC_ID_VC1
	CodecID_WMV3            CodecID = C.AV_CODEC_ID_WMV3
	CodecID_LOCO            CodecID = C.AV_CODEC_ID_LOCO
	CodecID_WNV1            CodecID = C.AV_CODEC_ID_WNV1
	CodecID_AASC            CodecID = C.AV_CODEC_ID_AASC
	CodecID_INDEO2          CodecID = C.AV_CODEC_ID_INDEO2
	CodecID_FRAPS           CodecID = C.AV_CODEC_ID_FRAPS
	CodecID_TRUEMOTION2     CodecID = C.AV_CODEC_ID_TRUEMOTION2
	CodecID_BMP             CodecID = C.AV_CODEC_ID_BMP
	CodecID_CSCD            CodecID = C.AV_CODEC_ID_CSCD
	CodecID_MMVIDEO         CodecID = C.AV_CODEC_ID_MMVIDEO
	CodecID_ZMBV            CodecID = C.AV_CODEC_ID_ZMBV
	CodecID_AVS             CodecID = C.AV_CODEC_ID_AVS
	CodecID_SMACKVIDEO      CodecID = C.AV_CODEC_ID_SMACKVIDEO
	CodecID_NUV             CodecID = C.AV_CODEC_ID_NUV
	CodecID_KMVC            CodecID = C.AV_CODEC_ID_KMVC
	CodecID_FLASHSV         CodecID = C.AV_CODEC_ID_FLASHSV
	CodecID_CAVS            CodecID = C.AV_CODEC_ID_CAVS
	CodecID_JPEG2000        CodecID = C.AV_CODEC_ID_JPEG2000
	CodecID_VMNC            CodecID = C.AV_CODEC_ID_VMNC
	CodecID_VP5             CodecID = C.AV_CODEC_ID_VP5
	CodecID_VP6             CodecID = C.AV_CODEC_ID_VP6
	CodecID_VP6F            CodecID = C.AV_CODEC_ID_VP6F
	CodecID_TARGA           CodecID = C.AV_CODEC_ID_TARGA
	CodecID_DSICINVIDEO     CodecID = C.AV_CODEC_ID_DSICINVIDEO
	CodecID_TIERTEXSEQVIDEO CodecID = C.AV_CODEC_ID_TIERTEXSEQVIDEO
	CodecID_TIFF            CodecID = C.AV_CODEC_ID_TIFF
	CodecID_GIF             CodecID = C.AV_CODEC_ID_GIF
	CodecID_DXA             CodecID = C.AV_CODEC_ID_DXA
	CodecID_DNXHD           CodecID = C.AV_CODEC_ID_DNXHD
	CodecID_THP             CodecID = C.AV_CODEC_ID_THP
	CodecID_SGI             CodecID = C.AV_CODEC_ID_SGI
	CodecID_C93             CodecID = C.AV_CODEC_ID_C93
	CodecID_BETHSOFTVID     CodecID = C.AV_CODEC_ID_BETHSOFTVID
	CodecID_PTX             CodecID = C.AV_CODEC_ID_PTX
	CodecID_TXD             CodecID = C.AV_CODEC_ID_TXD
	CodecID_VP6A            CodecID = C.AV_CODEC_ID_VP6A
	CodecID_AMV             CodecID = C.AV_CODEC_ID_AMV
	CodecID_VB              CodecID = C.AV_CODEC_ID_VB
	CodecID_PCX             CodecID = C.AV_CODEC_ID_PCX
	CodecID_SUNRAST         CodecID = C.AV_CODEC_ID_SUNRAST
	CodecID_INDEO4          CodecID = C.AV_CODEC_ID_INDEO4
	CodecID_INDEO5          CodecID = C.AV_CODEC_ID_INDEO5
	CodecID_MIMIC           CodecID = C.AV_CODEC_ID_MIMIC
	CodecID_RL2             CodecID = C.AV_CODEC_ID_RL2
	CodecID_ESCAPE124       CodecID = C.AV_CODEC_ID_ESCAPE124
	CodecID_DIRAC           CodecID = C.AV_CODEC_ID_DIRAC
	CodecID_BFI             CodecID = C.AV_CODEC_ID_BFI
	CodecID_CMV             CodecID = C.AV_CODEC_ID_CMV
	CodecID_MOTIONPIXELS    CodecID = C.AV_CODEC_ID_MOTIONPIXELS
	CodecID_TGV             CodecID = C.AV_CODEC_ID_TGV
	CodecID_TGQ             CodecID = C.AV_CODEC_ID_TGQ
	CodecID_TQI             CodecID = C.AV_CODEC_ID_TQI
	CodecID_AURA            CodecID = C.AV_CODEC_ID_AURA
	CodecID_AURA2           CodecID = C.AV_CODEC_ID_AURA2
	CodecID_V210X           CodecID = C.AV_CODEC_ID_V210X
	CodecID_TMV             CodecID = C.AV_CODEC_ID_TMV
	CodecID_V210            CodecID = C.AV_CODEC_ID_V210
	CodecID_DPX             CodecID = C.AV_CODEC_ID_DPX
	CodecID_MAD             CodecID = C.AV_CODEC_ID_MAD
	CodecID_FRWU            CodecID = C.AV_CODEC_ID_FRWU
	CodecID_FLASHSV2        CodecID = C.AV_CODEC_ID_FLASHSV2
	CodecID_CDGRAPHICS      CodecID = C.AV_CODEC_ID_CDGRAPHICS
	CodecID_R210            CodecID = C.AV_CODEC_ID_R210
	CodecID_ANM             CodecID = C.AV_CODEC_ID_ANM
	CodecID_BINKVIDEO       CodecID = C.AV_CODEC_ID_BINKVIDEO
	CodecID_IFF_ILBM        CodecID = C.AV_CODEC_ID_IFF_ILBM
	CodecID_KGV1            CodecID = C.AV_CODEC_ID_KGV1
	CodecID_YOP             CodecID = C.AV_CODEC_ID_YOP
	CodecID_VP8             CodecID = C.AV_CODEC_ID_VP8
	CodecID_PICTOR          CodecID = C.AV_CODEC_ID_PICTOR
	CodecID_ANSI            CodecID = C.AV_CODEC_ID_ANSI
	CodecID_A64_MULTI       CodecID = C.AV_CODEC_ID_A64_MULTI
	CodecID_A64_MULTI5      CodecID = C.AV_CODEC_ID_A64_MULTI5
	CodecID_R10K            CodecID = C.AV_CODEC_ID_R10K
	CodecID_MXPEG           CodecID = C.AV_CODEC_ID_MXPEG
	CodecID_LAGARITH        CodecID = C.AV_CODEC_ID_LAGARITH
	CodecID_PRORES          CodecID = C.AV_CODEC_ID_PRORES
	CodecID_JV              CodecID = C.AV_CODEC_ID_JV
	CodecID_DFA             CodecID = C.AV_CODEC_ID_DFA
	CodecID_WMV3IMAGE       CodecID = C.AV_CODEC_ID_WMV3IMAGE
	CodecID_VC1IMAGE        CodecID = C.AV_CODEC_ID_VC1IMAGE
	CodecID_UTVIDEO         CodecID = C.AV_CODEC_ID_UTVIDEO
	CodecID_BMV_VIDEO       CodecID = C.AV_CODEC_ID_BMV_VIDEO
	CodecID_VBLE            CodecID = C.AV_CODEC_ID_VBLE
	CodecID_DXTORY          CodecID = C.AV_CODEC_ID_DXTORY
	CodecID_V410            CodecID = C.AV_CODEC_ID_V410
	CodecID_XWD             CodecID = C.AV_CODEC_ID_XWD
	CodecID_CDXL            CodecID = C.AV_CODEC_ID_CDXL
	CodecID_XBM             CodecID = C.AV_CODEC_ID_XBM
	CodecID_ZEROCODEC       CodecID = C.AV_CODEC_ID_ZEROCODEC
	CodecID_MSS1            CodecID = C.AV_CODEC_ID_MSS1
	CodecID_MSA1            CodecID = C.AV_CODEC_ID_MSA1
	CodecID_TSCC2           CodecID = C.AV_CODEC_ID_TSCC2
	CodecID_MTS2            CodecID = C.AV_CODEC_ID_MTS2
	CodecID_CLLC            CodecID = C.AV_CODEC_ID_CLLC
	CodecID_MSS2            CodecID = C.AV_CODEC_ID_MSS2
	CodecID_VP9             CodecID = C.AV_CODEC_ID_VP9
	CodecID_AIC             CodecID = C.AV_CODEC_ID_AIC
	CodecID_ESCAPE130       CodecID = C.AV_CODEC_ID_ESCAPE130
	CodecID_G2M             CodecID = C.AV_CODEC_ID_G2M
	CodecID_WEBP            CodecID = C.AV_CODEC_ID_WEBP
	CodecID_HNM4_VIDEO      CodecID = C.AV_CODEC_ID_HNM4_VIDEO
	CodecID_HEVC            CodecID = C.AV_CODEC_ID_HEVC
	CodecID_FIC             CodecID = C.AV_CODEC_ID_FIC
	CodecID_ALIAS_PIX       CodecID = C.AV_CODEC_ID_ALIAS_PIX
	CodecID_BRENDER_PIX     CodecID = C.AV_CODEC_ID_BRENDER_PIX
	CodecID_PAF_VIDEO       CodecID = C.AV_CODEC_ID_PAF_VIDEO
	CodecID_EXR             CodecID = C.AV_CODEC_ID_EXR
	CodecID_VP7             CodecID = C.AV_CODEC_ID_VP7
	CodecID_SANM            CodecID = C.AV_CODEC_ID_SANM
	CodecID_SGIRLE          CodecID = C.AV_CODEC_ID_SGIRLE
	CodecID_MVC1            CodecID = C.AV_CODEC_ID_MVC1
	CodecID_MVC2            CodecID = C.AV_CODEC_ID_MVC2
	CodecID_HQX             CodecID = C.AV_CODEC_ID_HQX
	CodecID_TDSC            CodecID = C.AV_CODEC_ID_TDSC
	CodecID_HQ_HQA          CodecID = C.AV_CODEC_ID_HQ_HQA
	CodecID_HAP             CodecID = C.AV_CODEC_ID_HAP
	CodecID_DDS             CodecID = C.AV_CODEC_ID_DDS
	CodecID_DXV             CodecID = C.AV_CODEC_ID_DXV
	CodecID_SCREENPRESSO    CodecID = C.AV_CODEC_ID_SCREENPRESSO
	CodecID_RSCC            CodecID = C.AV_CODEC_ID_RSCC
	CodecID_AVS2            CodecID = C.AV_CODEC_ID_AVS2
	CodecID_Y41P            CodecID = C.AV_CODEC_ID_Y41P
	CodecID_AVRP            CodecID = C.AV_CODEC_ID_AVRP
	CodecID_012V            CodecID = C.AV_CODEC_ID_012V
	CodecID_AVUI            CodecID = C.AV_CODEC_ID_AVUI
	CodecID_AYUV            CodecID = C.AV_CODEC_ID_AYUV
	CodecID_TARGA_Y216      CodecID = C.AV_CODEC_ID_TARGA_Y216
	CodecID_V308            CodecID = C.AV_CODEC_ID_V308
	CodecID_V408            CodecID = C.AV_CODEC_ID_V408
	CodecID_YUV4            CodecID = C.AV_CODEC_ID_YUV4
	CodecID_AVRN            CodecID = C.AV_CODEC_ID_AVRN
	CodecID_CPIA            CodecID = C.AV_CODEC_ID_CPIA
	CodecID_XFACE           CodecID = C.AV_CODEC_ID_XFACE
	CodecID_SNOW            CodecID = C.AV_CODEC_ID_SNOW
	CodecID_SMVJPEG         CodecID = C.AV_CODEC_ID_SMVJPEG
	CodecID_APNG            CodecID = C.AV_CODEC_ID_APNG
	CodecID_DAALA           CodecID = C.AV_CODEC_ID_DAALA
	CodecID_CFHD            CodecID = C.AV_CODEC_ID_CFHD
	CodecID_TRUEMOTION2RT   CodecID = C.AV_CODEC_ID_TRUEMOTION2RT
	CodecID_M101            CodecID = C.AV_CODEC_ID_M101
	CodecID_MAGICYUV        CodecID = C.AV_CODEC_ID_MAGICYUV
	CodecID_SHEERVIDEO      CodecID = C.AV_CODEC_ID_SHEERVIDEO
	CodecID_YLC             CodecID = C.AV_CODEC_ID_YLC
	CodecID_PSD             CodecID = C.AV_CODEC_ID_PSD
	CodecID_PIXLET          CodecID = C.AV_CODEC_ID_PIXLET
	CodecID_SPEEDHQ         CodecID = C.AV_CODEC_ID_SPEEDHQ
	CodecID_FMVC            CodecID = C.AV_CODEC_ID_FMVC
	CodecID_SCPR            CodecID = C.AV_CODEC_ID_SCPR
	CodecID_CLEARVIDEO      CodecID = C.AV_CODEC_ID_CLEARVIDEO
	CodecID_XPM             CodecID = C.AV_CODEC_ID_XPM
	CodecID_AV1             CodecID = C.AV_CODEC_ID_AV1
	CodecID_BITPACKED       CodecID = C.AV_CODEC_ID_BITPACKED
	CodecID_MSCC            CodecID = C.AV_CODEC_ID_MSCC
	CodecID_SRGC            CodecID = C.AV_CODEC_ID_SRGC
	CodecID_SVG             CodecID = C.AV_CODEC_ID_SVG
	CodecID_GDV             CodecID = C.AV_CODEC_ID_GDV
	CodecID_FITS            CodecID = C.AV_CODEC_ID_FITS
	CodecID_IMM4            CodecID = C.AV_CODEC_ID_IMM4
	CodecID_PROSUMER        CodecID = C.AV_CODEC_ID_PROSUMER
	CodecID_MWSC            CodecID = C.AV_CODEC_ID_MWSC
	CodecID_RASC            CodecID = C.AV_CODEC_ID_RASC

	//--------------------------------------
	// Audio codecs
	//--------------------------------------

	CodecID_FIRST_AUDIO      CodecID = C.AV_CODEC_ID_FIRST_AUDIO
	CodecID_PCM_S16LE        CodecID = C.AV_CODEC_ID_PCM_S16LE
	CodecID_PCM_S16BE        CodecID = C.AV_CODEC_ID_PCM_S16BE
	CodecID_PCM_U16LE        CodecID = C.AV_CODEC_ID_PCM_U16LE
	CodecID_PCM_U16BE        CodecID = C.AV_CODEC_ID_PCM_U16BE
	CodecID_PCM_U8           CodecID = C.AV_CODEC_ID_PCM_U8
	CodecID_PCM_S8           CodecID = C.AV_CODEC_ID_PCM_S8
	CodecID_PCM_MULAW        CodecID = C.AV_CODEC_ID_PCM_MULAW
	CodecID_PCM_ALAW         CodecID = C.AV_CODEC_ID_PCM_ALAW
	CodecID_PCM_S32LE        CodecID = C.AV_CODEC_ID_PCM_S32LE
	CodecID_PCM_S32BE        CodecID = C.AV_CODEC_ID_PCM_S32BE
	CodecID_PCM_U32LE        CodecID = C.AV_CODEC_ID_PCM_U32LE
	CodecID_PCM_U32BE        CodecID = C.AV_CODEC_ID_PCM_U32BE
	CodecID_PCM_S24LE        CodecID = C.AV_CODEC_ID_PCM_S24LE
	CodecID_PCM_S24BE        CodecID = C.AV_CODEC_ID_PCM_S24BE
	CodecID_PCM_U24LE        CodecID = C.AV_CODEC_ID_PCM_U24LE
	CodecID_PCM_U24BE        CodecID = C.AV_CODEC_ID_PCM_U24BE
	CodecID_PCM_S24DAUD      CodecID = C.AV_CODEC_ID_PCM_S24DAUD
	CodecID_PCM_ZORK         CodecID = C.AV_CODEC_ID_PCM_ZORK
	CodecID_PCM_S16LE_PLANAR CodecID = C.AV_CODEC_ID_PCM_S16LE_PLANAR
	CodecID_PCM_DVD          CodecID = C.AV_CODEC_ID_PCM_DVD
	CodecID_PCM_F32BE        CodecID = C.AV_CODEC_ID_PCM_F32BE
	CodecID_PCM_F32LE        CodecID = C.AV_CODEC_ID_PCM_F32LE
	CodecID_PCM_F64BE        CodecID = C.AV_CODEC_ID_PCM_F64BE
	CodecID_PCM_F64LE        CodecID = C.AV_CODEC_ID_PCM_F64LE
	CodecID_PCM_BLURAY       CodecID = C.AV_CODEC_ID_PCM_BLURAY
	CodecID_PCM_LXF          CodecID = C.AV_CODEC_ID_PCM_LXF
	CodecID_S302M            CodecID = C.AV_CODEC_ID_S302M
	CodecID_PCM_S8_PLANAR    CodecID = C.AV_CODEC_ID_PCM_S8_PLANAR
	CodecID_PCM_S24LE_PLANAR CodecID = C.AV_CODEC_ID_PCM_S24LE_PLANAR
	CodecID_PCM_S32LE_PLANAR CodecID = C.AV_CODEC_ID_PCM_S32LE_PLANAR
	CodecID_PCM_S16BE_PLANAR CodecID = C.AV_CODEC_ID_PCM_S16BE_PLANAR
	CodecID_PCM_S64LE        CodecID = C.AV_CODEC_ID_PCM_S64LE
	CodecID_PCM_S64BE        CodecID = C.AV_CODEC_ID_PCM_S64BE
	CodecID_PCM_F16LE        CodecID = C.AV_CODEC_ID_PCM_F16LE
	CodecID_PCM_F24LE        CodecID = C.AV_CODEC_ID_PCM_F24LE
	CodecID_PCM_VIDC         CodecID = C.AV_CODEC_ID_PCM_VIDC

	CodecID_ADPCM_IMA_QT      CodecID = C.AV_CODEC_ID_ADPCM_IMA_QT
	CodecID_ADPCM_IMA_WAV     CodecID = C.AV_CODEC_ID_ADPCM_IMA_WAV
	CodecID_ADPCM_IMA_DK3     CodecID = C.AV_CODEC_ID_ADPCM_IMA_DK3
	CodecID_ADPCM_IMA_DK4     CodecID = C.AV_CODEC_ID_ADPCM_IMA_DK4
	CodecID_ADPCM_IMA_SMJPEG  CodecID = C.AV_CODEC_ID_ADPCM_IMA_SMJPEG
	CodecID_ADPCM_MS          CodecID = C.AV_CODEC_ID_ADPCM_MS
	CodecID_ADPCM_4XM         CodecID = C.AV_CODEC_ID_ADPCM_4XM
	CodecID_ADPCM_XA          CodecID = C.AV_CODEC_ID_ADPCM_XA
	CodecID_ADPCM_ADX         CodecID = C.AV_CODEC_ID_ADPCM_ADX
	CodecID_ADPCM_EA          CodecID = C.AV_CODEC_ID_ADPCM_EA
	CodecID_ADPCM_G726        CodecID = C.AV_CODEC_ID_ADPCM_G726
	CodecID_ADPCM_CT          CodecID = C.AV_CODEC_ID_ADPCM_CT
	CodecID_ADPCM_SWF         CodecID = C.AV_CODEC_ID_ADPCM_SWF
	CodecID_ADPCM_YAMAHA      CodecID = C.AV_CODEC_ID_ADPCM_YAMAHA
	CodecID_ADPCM_SBPRO_4     CodecID = C.AV_CODEC_ID_ADPCM_SBPRO_4
	CodecID_ADPCM_SBPRO_3     CodecID = C.AV_CODEC_ID_ADPCM_SBPRO_3
	CodecID_ADPCM_SBPRO_2     CodecID = C.AV_CODEC_ID_ADPCM_SBPRO_2
	CodecID_ADPCM_THP         CodecID = C.AV_CODEC_ID_ADPCM_THP
	CodecID_ADPCM_IMA_AMV     CodecID = C.AV_CODEC_ID_ADPCM_IMA_AMV
	CodecID_ADPCM_EA_R1       CodecID = C.AV_CODEC_ID_ADPCM_EA_R1
	CodecID_ADPCM_EA_R3       CodecID = C.AV_CODEC_ID_ADPCM_EA_R3
	CodecID_ADPCM_EA_R2       CodecID = C.AV_CODEC_ID_ADPCM_EA_R2
	CodecID_ADPCM_IMA_EA_SEAD CodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD
	CodecID_ADPCM_IMA_EA_EACS CodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_EACS
	CodecID_ADPCM_EA_XAS      CodecID = C.AV_CODEC_ID_ADPCM_EA_XAS
	CodecID_ADPCM_EA_MAXIS_XA CodecID = C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA
	CodecID_ADPCM_IMA_ISS     CodecID = C.AV_CODEC_ID_ADPCM_IMA_ISS
	CodecID_ADPCM_G722        CodecID = C.AV_CODEC_ID_ADPCM_G722
	CodecID_ADPCM_IMA_APC     CodecID = C.AV_CODEC_ID_ADPCM_IMA_APC
	CodecID_ADPCM_VIMA        CodecID = C.AV_CODEC_ID_ADPCM_VIMA
	CodecID_ADPCM_AFC         CodecID = C.AV_CODEC_ID_ADPCM_AFC
	CodecID_ADPCM_IMA_OKI     CodecID = C.AV_CODEC_ID_ADPCM_IMA_OKI
	CodecID_ADPCM_DTK         CodecID = C.AV_CODEC_ID_ADPCM_DTK
	CodecID_ADPCM_IMA_RAD     CodecID = C.AV_CODEC_ID_ADPCM_IMA_RAD
	CodecID_ADPCM_G726LE      CodecID = C.AV_CODEC_ID_ADPCM_G726LE
	CodecID_ADPCM_THP_LE      CodecID = C.AV_CODEC_ID_ADPCM_THP_LE
	CodecID_ADPCM_PSX         CodecID = C.AV_CODEC_ID_ADPCM_PSX
	CodecID_ADPCM_AICA        CodecID = C.AV_CODEC_ID_ADPCM_AICA
	CodecID_ADPCM_IMA_DAT4    CodecID = C.AV_CODEC_ID_ADPCM_IMA_DAT4
	CodecID_ADPCM_MTAF        CodecID = C.AV_CODEC_ID_ADPCM_MTAF

	CodecID_AMR_NB          CodecID = C.AV_CODEC_ID_AMR_NB
	CodecID_AMR_WB          CodecID = C.AV_CODEC_ID_AMR_WB
	CodecID_RA_144          CodecID = C.AV_CODEC_ID_RA_144
	CodecID_RA_288          CodecID = C.AV_CODEC_ID_RA_288
	CodecID_ROQ_DPCM        CodecID = C.AV_CODEC_ID_ROQ_DPCM
	CodecID_INTERPLAY_DPCM  CodecID = C.AV_CODEC_ID_INTERPLAY_DPCM
	CodecID_XAN_DPCM        CodecID = C.AV_CODEC_ID_XAN_DPCM
	CodecID_SOL_DPCM        CodecID = C.AV_CODEC_ID_SOL_DPCM
	CodecID_SDX2_DPCM       CodecID = C.AV_CODEC_ID_SDX2_DPCM
	CodecID_GREMLIN_DPCM    CodecID = C.AV_CODEC_ID_GREMLIN_DPCM
	CodecID_MP2             CodecID = C.AV_CODEC_ID_MP2
	CodecID_MP3             CodecID = C.AV_CODEC_ID_MP3
	CodecID_AAC             CodecID = C.AV_CODEC_ID_AAC
	CodecID_AC3             CodecID = C.AV_CODEC_ID_AC3
	CodecID_DTS             CodecID = C.AV_CODEC_ID_DTS
	CodecID_VORBIS          CodecID = C.AV_CODEC_ID_VORBIS
	CodecID_DVAUDIO         CodecID = C.AV_CODEC_ID_DVAUDIO
	CodecID_WMAV1           CodecID = C.AV_CODEC_ID_WMAV1
	CodecID_WMAV2           CodecID = C.AV_CODEC_ID_WMAV2
	CodecID_MACE3           CodecID = C.AV_CODEC_ID_MACE3
	CodecID_MACE6           CodecID = C.AV_CODEC_ID_MACE6
	CodecID_VMDAUDIO        CodecID = C.AV_CODEC_ID_VMDAUDIO
	CodecID_FLAC            CodecID = C.AV_CODEC_ID_FLAC
	CodecID_MP3ADU          CodecID = C.AV_CODEC_ID_MP3ADU
	CodecID_MP3ON4          CodecID = C.AV_CODEC_ID_MP3ON4
	CodecID_SHORTEN         CodecID = C.AV_CODEC_ID_SHORTEN
	CodecID_ALAC            CodecID = C.AV_CODEC_ID_ALAC
	CodecID_WESTWOOD_SND1   CodecID = C.AV_CODEC_ID_WESTWOOD_SND1
	CodecID_GSM             CodecID = C.AV_CODEC_ID_GSM
	CodecID_QDM2            CodecID = C.AV_CODEC_ID_QDM2
	CodecID_COOK            CodecID = C.AV_CODEC_ID_COOK
	CodecID_TRUESPEECH      CodecID = C.AV_CODEC_ID_TRUESPEECH
	CodecID_TTA             CodecID = C.AV_CODEC_ID_TTA
	CodecID_SMACKAUDIO      CodecID = C.AV_CODEC_ID_SMACKAUDIO
	CodecID_QCELP           CodecID = C.AV_CODEC_ID_QCELP
	CodecID_WAVPACK         CodecID = C.AV_CODEC_ID_WAVPACK
	CodecID_DSICINAUDIO     CodecID = C.AV_CODEC_ID_DSICINAUDIO
	CodecID_IMC             CodecID = C.AV_CODEC_ID_IMC
	CodecID_MUSEPACK7       CodecID = C.AV_CODEC_ID_MUSEPACK7
	CodecID_MLP             CodecID = C.AV_CODEC_ID_MLP
	CodecID_GSM_MS          CodecID = C.AV_CODEC_ID_GSM_MS
	CodecID_ATRAC3          CodecID = C.AV_CODEC_ID_ATRAC3
	CodecID_APE             CodecID = C.AV_CODEC_ID_APE
	CodecID_NELLYMOSER      CodecID = C.AV_CODEC_ID_NELLYMOSER
	CodecID_MUSEPACK8       CodecID = C.AV_CODEC_ID_MUSEPACK8
	CodecID_SPEEX           CodecID = C.AV_CODEC_ID_SPEEX
	CodecID_WMAVOICE        CodecID = C.AV_CODEC_ID_WMAVOICE
	CodecID_WMAPRO          CodecID = C.AV_CODEC_ID_WMAPRO
	CodecID_WMALOSSLESS     CodecID = C.AV_CODEC_ID_WMALOSSLESS
	CodecID_ATRAC3P         CodecID = C.AV_CODEC_ID_ATRAC3P
	CodecID_EAC3            CodecID = C.AV_CODEC_ID_EAC3
	CodecID_SIPR            CodecID = C.AV_CODEC_ID_SIPR
	CodecID_MP1             CodecID = C.AV_CODEC_ID_MP1
	CodecID_TWINVQ          CodecID = C.AV_CODEC_ID_TWINVQ
	CodecID_TRUEHD          CodecID = C.AV_CODEC_ID_TRUEHD
	CodecID_MP4ALS          CodecID = C.AV_CODEC_ID_MP4ALS
	CodecID_ATRAC1          CodecID = C.AV_CODEC_ID_ATRAC1
	CodecID_BINKAUDIO_RDFT  CodecID = C.AV_CODEC_ID_BINKAUDIO_RDFT
	CodecID_BINKAUDIO_DCT   CodecID = C.AV_CODEC_ID_BINKAUDIO_DCT
	CodecID_AAC_LATM        CodecID = C.AV_CODEC_ID_AAC_LATM
	CodecID_QDMC            CodecID = C.AV_CODEC_ID_QDMC
	CodecID_CELT            CodecID = C.AV_CODEC_ID_CELT
	CodecID_G723_1          CodecID = C.AV_CODEC_ID_G723_1
	CodecID_G729            CodecID = C.AV_CODEC_ID_G729
	CodecID_8SVX_EXP        CodecID = C.AV_CODEC_ID_8SVX_EXP
	CodecID_8SVX_FIB        CodecID = C.AV_CODEC_ID_8SVX_FIB
	CodecID_BMV_AUDIO       CodecID = C.AV_CODEC_ID_BMV_AUDIO
	CodecID_RALF            CodecID = C.AV_CODEC_ID_RALF
	CodecID_IAC             CodecID = C.AV_CODEC_ID_IAC
	CodecID_ILBC            CodecID = C.AV_CODEC_ID_ILBC
	CodecID_OPUS            CodecID = C.AV_CODEC_ID_OPUS
	CodecID_COMFORT_NOISE   CodecID = C.AV_CODEC_ID_COMFORT_NOISE
	CodecID_TAK             CodecID = C.AV_CODEC_ID_TAK
	CodecID_METASOUND       CodecID = C.AV_CODEC_ID_METASOUND
	CodecID_PAF_AUDIO       CodecID = C.AV_CODEC_ID_PAF_AUDIO
	CodecID_ON2AVC          CodecID = C.AV_CODEC_ID_ON2AVC
	CodecID_DSS_SP          CodecID = C.AV_CODEC_ID_DSS_SP
	CodecID_CODEC2          CodecID = C.AV_CODEC_ID_CODEC2
	CodecID_FFWAVESYNTH     CodecID = C.AV_CODEC_ID_FFWAVESYNTH
	CodecID_SONIC           CodecID = C.AV_CODEC_ID_SONIC
	CodecID_SONIC_LS        CodecID = C.AV_CODEC_ID_SONIC_LS
	CodecID_EVRC            CodecID = C.AV_CODEC_ID_EVRC
	CodecID_SMV             CodecID = C.AV_CODEC_ID_SMV
	CodecID_DSD_LSBF        CodecID = C.AV_CODEC_ID_DSD_LSBF
	CodecID_DSD_MSBF        CodecID = C.AV_CODEC_ID_DSD_MSBF
	CodecID_DSD_LSBF_PLANAR CodecID = C.AV_CODEC_ID_DSD_LSBF_PLANAR
	CodecID_DSD_MSBF_PLANAR CodecID = C.AV_CODEC_ID_DSD_MSBF_PLANAR
	CodecID_4GV             CodecID = C.AV_CODEC_ID_4GV
	CodecID_INTERPLAY_ACM   CodecID = C.AV_CODEC_ID_INTERPLAY_ACM
	CodecID_XMA1            CodecID = C.AV_CODEC_ID_XMA1
	CodecID_XMA2            CodecID = C.AV_CODEC_ID_XMA2
	CodecID_DST             CodecID = C.AV_CODEC_ID_DST
	CodecID_ATRAC3AL        CodecID = C.AV_CODEC_ID_ATRAC3AL
	CodecID_ATRAC3PAL       CodecID = C.AV_CODEC_ID_ATRAC3PAL
	CodecID_DOLBY_E         CodecID = C.AV_CODEC_ID_DOLBY_E
	CodecID_APTX            CodecID = C.AV_CODEC_ID_APTX
	CodecID_APTX_HD         CodecID = C.AV_CODEC_ID_APTX_HD
	CodecID_SBC             CodecID = C.AV_CODEC_ID_SBC
	CodecID_ATRAC9          CodecID = C.AV_CODEC_ID_ATRAC9

	//--------------------------------------
	// Subtitle codecs
	//--------------------------------------

	CodecID_FIRST_SUBTITLE     CodecID = C.AV_CODEC_ID_FIRST_SUBTITLE
	CodecID_DVD_SUBTITLE       CodecID = C.AV_CODEC_ID_DVD_SUBTITLE
	CodecID_DVB_SUBTITLE       CodecID = C.AV_CODEC_ID_DVB_SUBTITLE
	CodecID_TEXT               CodecID = C.AV_CODEC_ID_TEXT
	CodecID_XSUB               CodecID = C.AV_CODEC_ID_XSUB
	CodecID_SSA                CodecID = C.AV_CODEC_ID_SSA
	CodecID_MOV_TEXT           CodecID = C.AV_CODEC_ID_MOV_TEXT
	CodecID_HDMV_PGS_SUBTITLE  CodecID = C.AV_CODEC_ID_HDMV_PGS_SUBTITLE
	CodecID_DVB_TELETEXT       CodecID = C.AV_CODEC_ID_DVB_TELETEXT
	CodecID_SRT                CodecID = C.AV_CODEC_ID_SRT
	CodecID_MICRODVD           CodecID = C.AV_CODEC_ID_MICRODVD
	CodecID_EIA_608            CodecID = C.AV_CODEC_ID_EIA_608
	CodecID_JACOSUB            CodecID = C.AV_CODEC_ID_JACOSUB
	CodecID_SAMI               CodecID = C.AV_CODEC_ID_SAMI
	CodecID_REALTEXT           CodecID = C.AV_CODEC_ID_REALTEXT
	CodecID_STL                CodecID = C.AV_CODEC_ID_STL
	CodecID_SUBVIEWER1         CodecID = C.AV_CODEC_ID_SUBVIEWER1
	CodecID_SUBVIEWER          CodecID = C.AV_CODEC_ID_SUBVIEWER
	CodecID_SUBRIP             CodecID = C.AV_CODEC_ID_SUBRIP
	CodecID_WEBVTT             CodecID = C.AV_CODEC_ID_WEBVTT
	CodecID_MPL2               CodecID = C.AV_CODEC_ID_MPL2
	CodecID_VPLAYER            CodecID = C.AV_CODEC_ID_VPLAYER
	CodecID_PJS                CodecID = C.AV_CODEC_ID_PJS
	CodecID_ASS                CodecID = C.AV_CODEC_ID_ASS
	CodecID_HDMV_TEXT_SUBTITLE CodecID = C.AV_CODEC_ID_HDMV_TEXT_SUBTITLE
	CodecID_TTML               CodecID = C.AV_CODEC_ID_TTML

	//--------------------------------------
	// Fake codecs
	//--------------------------------------

	CodecID_FIRST_UNKNOWN   CodecID = C.AV_CODEC_ID_FIRST_UNKNOWN
	CodecID_TTF             CodecID = C.AV_CODEC_ID_TTF
	CodecID_SCTE_35         CodecID = C.AV_CODEC_ID_SCTE_35
	CodecID_BINTEXT         CodecID = C.AV_CODEC_ID_BINTEXT
	CodecID_XBIN            CodecID = C.AV_CODEC_ID_XBIN
	CodecID_IDF             CodecID = C.AV_CODEC_ID_IDF
	CodecID_OTF             CodecID = C.AV_CODEC_ID_OTF
	CodecID_SMPTE_KLV       CodecID = C.AV_CODEC_ID_SMPTE_KLV
	CodecID_DVD_NAV         CodecID = C.AV_CODEC_ID_DVD_NAV
	CodecID_TIMED_ID3       CodecID = C.AV_CODEC_ID_TIMED_ID3
	CodecID_BIN_DATA        CodecID = C.AV_CODEC_ID_BIN_DATA
	CodecID_PROBE           CodecID = C.AV_CODEC_ID_PROBE
	CodecID_MPEG2TS         CodecID = C.AV_CODEC_ID_MPEG2TS
	CodecID_MPEG4SYSTEMS    CodecID = C.AV_CODEC_ID_MPEG4SYSTEMS
	CodecID_FFMETADATA      CodecID = C.AV_CODEC_ID_FFMETADATA
	CodecID_WRAPPED_AVFRAME CodecID = C.AV_CODEC_ID_WRAPPED_AVFRAME
)

type CodecParameters = C.struct_AVCodecParameters

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

// Fill the parameters struct based on the values from the supplied codec
// context.
//
// Any allocated fields in par are freed and replaced with duplicates of the
// corresponding fields in codec.
//
// Returns
//     >= 0 on success, a negative AVERROR code on failure
func (cp *CodecParameters) FromContext(cctx *CodecContext) error {
	return CheckErr(C.avcodec_parameters_from_context(cp, cctx))
}

func (cp *CodecParameters) Copy(cpDest *CodecParameters) error {
	return CheckErr(C.avcodec_parameters_copy(cp, cpDest))
}

type Dictionary = C.struct_AVDictionary
type DictionaryEntry = C.struct_AVDictionaryEntry
type DictionaryFlags = C.int

var (
	DictionaryFlags_MATCH_CASE      DictionaryFlags = C.AV_DICT_MATCH_CASE
	DictionaryFlags_IGNORE_SUFFIX   DictionaryFlags = C.AV_DICT_IGNORE_SUFFIX
	DictionaryFlags_DONT_STRDUP_KEY DictionaryFlags = C.AV_DICT_DONT_STRDUP_KEY
	DictionaryFlags_DONT_STRDUP_VAL DictionaryFlags = C.AV_DICT_DONT_STRDUP_VAL
	DictionaryFlags_DONT_OVERWRITE  DictionaryFlags = C.AV_DICT_DONT_OVERWRITE
	DictionaryFlags_APPEND          DictionaryFlags = C.AV_DICT_APPEND
	DictionaryFlags_MULTIKEY        DictionaryFlags = C.AV_DICT_MULTIKEY
)

// Copy entries from one AVDictionary struct into another.
//
// Parameters
//   dst	pointer to a pointer to a AVDictionary struct. If *dst is NULL, this function will allocate a struct for you and put it in *dst
//   src	pointer to source AVDictionary struct
//   flags	flags to use when setting entries in *dst
//
// Note
//   metadata is read using the AV_DICT_IGNORE_SUFFIX flag
//
// Returns
//   0 on success, negative AVERROR code on failure. If dst was allocated by this function, callers should free the associated memory.
func (d *Dictionary) Copy(flags DictionaryFlags) (*Dictionary, error) {
	var dst *Dictionary
	err := CheckErr(C.av_dict_copy(&dst, d, C.int(flags)))
	return dst, err
}

func (d *Dictionary) AsMap() map[string]string {
	m := make(map[string]string)
	for _, k := range d.Keys() {
		if v, ok := d.Get(k); ok {
			m[k] = v
		}
	}
	return m
}

func (d *Dictionary) Keys() []string {
	var keys []string

	// gotcha: can't use our CString() wrapper
	// because that'll give us a nil pointer
	key_ := C.CString("")
	defer FreeString(key_)

	var entry *DictionaryEntry = nil
	for {
		entry = C.av_dict_get(d, key_, entry, DictionaryFlags_IGNORE_SUFFIX)
		if entry == nil {
			break
		}
		keys = append(keys, entry.Key())
	}
	return keys
}

func (d *Dictionary) Get(key string) (string, bool) {
	key_ := CString(key)
	defer FreeString(key_)

	var entry *DictionaryEntry
	entry = C.av_dict_get(d, key_, nil, 0)

	if entry == nil {
		return "", false
	}
	return entry.Value(), true
}

// Set the given entry in *pm, overwriting an existing entry.
//
// Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set, these
// arguments will be freed on error.
//
// Warning: Adding a new entry to a dictionary invalidates all existing entries
// previously returned with av_dict_get.
//
// Parameters
//     pm pointer to a pointer to a dictionary struct. If *pm is NULL a dictionary
//     struct is allocated and put in *pm.
//     key entry key to add to *pm (will either be av_strduped or added as a new key
//     depending on flags)
//     value entry value to add to *pm (will be av_strduped or added as a new key
//     depending on flags). Passing a NULL value will cause an existing entry to be
//     deleted.
//
// Returns
//     >= 0 on success otherwise an error code <0
func (d *Dictionary) Set(key string, value string) (*Dictionary, error) {
	key_ := CString(key)
	defer FreeString(key_)

	value_ := CString(value)
	defer FreeString(value_)

	err := CheckErr(C.av_dict_set(&d, key_, value_, 0))
	return d, err
}

// Free all the memory allocated for an AVDictionary struct and all keys and
// values.
func (d *Dictionary) Free() {
	C.av_dict_free(&d)
}

func (e *DictionaryEntry) Key() string {
	return C.GoString(e.key)
}

func (e *DictionaryEntry) Value() string {
	return C.GoString(e.value)
}

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

func FormatAllocContext() *FormatContext {
	return C.avformat_alloc_context()
}

func (ctx *FormatContext) OpenInput(url string, format *InputFormat, options **Dictionary) error {
	url_ := CString(url)
	defer FreeString(url_)

	ret := C.avformat_open_input(&ctx, url_, format, options)
	return CheckErr(ret)
}

func (ctx *FormatContext) OpenReader(r *Reader, format *InputFormat, options **Dictionary) error {
	ctx.SetPB(r.ctx)
	ret := C.avformat_open_input(&ctx, nil, format, options)
	return CheckErr(ret)
}

func (ctx *FormatContext) SetWriter(w *Writer) {
	ctx.SetPB(w.ctx)
}

type SeekFlag int

const (
	// Seek backward
	SeekFlagBackward SeekFlag = C.AVSEEK_FLAG_BACKWARD
	// Seeking based on position in bytes
	SeekFlagByte SeekFlag = C.AVSEEK_FLAG_BYTE
	// Seek to any frame, even non-keyframes
	SeekFlagAny SeekFlag = C.AVSEEK_FLAG_ANY
	// Seeking based on frame number
	SeekFlagFrame SeekFlag = C.AVSEEK_FLAG_FRAME
)

func (ctx *FormatContext) Seek(streamIndex int, timestamp Timing, flags SeekFlag) error {
	return CheckErr(C.av_seek_frame(ctx, C.int(streamIndex), timestamp, C.int(flags)))
}

func (ctx *FormatContext) PB() *IOContext {
	return ctx.pb
}

func (ctx *FormatContext) SetPB(pb *IOContext) {
	ctx.pb = pb
}

func (ctx *FormatContext) InputFormat() *InputFormat {
	return ctx.iformat
}

func (ctx *FormatContext) SetInputFormat(ifmt *InputFormat) {
	ctx.iformat = ifmt
}

func (ctx *FormatContext) OutputFormat() *OutputFormat {
	return ctx.oformat
}

func (ctx *FormatContext) SetOutputFormat(ofmt *OutputFormat) {
	ctx.oformat = ofmt
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

// Print detailed information about the input or output format, such as
// duration, bitrate, streams, container, programs, metadata, side data, codec
// and time base.
//
// Parameters
//   ic	the context to analyze
//   index	index of the stream to dump information about
//   url	the URL to print, such as source or destination file
//   is_output Select whether the specified context is an input(0) or output(1)
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

// Allocate the stream private data and write the stream header to an output media file.
//
// Parameters
//   s Media file handle, must be allocated with avformat_alloc_context(). Its
//   oformat field must be set to the desired output format; Its pb field must be
//   set to an already opened AVIOContext.
//   options An AVDictionary filled with AVFormatContext and muxer-private
//   options. On return this parameter will be destroyed and replaced with a dict
//   containing options that were not found. May be NULL.
//
// Returns
//   AVSTREAM_INIT_IN_WRITE_HEADER on success if the codec had not already been
//   fully initialized in avformat_init, AVSTREAM_INIT_IN_INIT_OUTPUT on success
//   if the codec had already been fully initialized in avformat_init, negative
//   AVERROR on failure.
//
// See Also
//   av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
func (ctx *FormatContext) WriteHeader(options **Dictionary) error {
	return CheckErr(C.avformat_write_header(ctx, options))
}

// Write the stream trailer to an output media file and free the file private data.
//
// May only be called after a successful call to avformat_write_header.
//
// Parameters
//     s	media file handle
//
// Returns
//     0 if OK, AVERROR_xxx on error
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

type Frame = C.struct_AVFrame

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

func (frame *Frame) SetPts(pts Timing) {
	frame.pts = pts
}

// frame timestamp estimated using various heuristics, in stream time base
func (frame *Frame) BestEffortTimestamp() Timing {
	return frame.best_effort_timestamp
}

// DTS copied from the AVPacket that triggered returning this frame.
func (frame *Frame) PktDts() Timing {
	return frame.pkt_dts
}

// duration of the corresponding packet, expressed in AVStream->time_base units, 0 if unknown.
func (frame *Frame) PktDuration() Timing {
	return frame.pkt_duration
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

func (frame *Frame) PictType() PictureType {
	return PictureType(frame.pict_type)
}

func (frame *Frame) SetPictType(pt PictureType) {
	frame.pict_type = uint32(pt)
}

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

type Reader struct {
	inner io.ReadSeeker
	size  int64
	ctx   *IOContext
	id    int32
	freed bool
	err   error
}

type Writer struct {
	inner io.WriteSeeker
	ctx   *IOContext
	id    int32
	freed bool
	err   error
}

var bridgeSeed int32 = 1
var readers sync.Map
var writers sync.Map

func reserveReaderId(obj *Reader) {
	obj.id = atomic.AddInt32(&bridgeSeed, 1)
	readers.Store(obj.id, obj)
}

func reserveWriterId(obj *Writer) {
	obj.id = atomic.AddInt32(&bridgeSeed, 1)
	writers.Store(obj.id, obj)
}

func freeReaderId(id int32) {
	readers.Delete(id)
}

func freeWriterId(id int32) {
	writers.Delete(id)
}

func finalizeReader(obj *Reader) {
	obj.Free()
}

func finalizeWriter(obj *Writer) {
	obj.Free()
}

func NewReader(inner io.ReadSeeker, size int64) *Reader {
	r := &Reader{
		inner: inner,
		size:  size,
	}
	runtime.SetFinalizer(r, finalizeReader)
	reserveReaderId(r)

	bufferSize := 4 * 1024 // 4KB, see doc for `avio_alloc_context`
	buffer := C.av_malloc(C.size_t(bufferSize))
	r.ctx = C.avio_alloc_context(
		(*C.uchar)(buffer),            // buffer
		C.int(bufferSize),             // buffer_size
		0,                             // write_flag
		unsafe.Pointer(uintptr(r.id)), // opaque
		CCallback(C.goff_reader_read_packet_trampoline), // read_packet
		nil,                                      // write_packet
		CCallback(C.goff_reader_seek_trampoline), // seek
	)
	return r
}

func NewWriter(inner io.WriteSeeker) *Writer {
	w := &Writer{
		inner: inner,
	}
	runtime.SetFinalizer(w, finalizeWriter)
	reserveWriterId(w)

	bufferSize := 4 * 1024 // 4KB, see doc for `avio_alloc_context`
	buffer := C.av_malloc(C.size_t(bufferSize))
	w.ctx = C.avio_alloc_context(
		(*C.uchar)(buffer), // buffer
		C.int(bufferSize),  // buffer_size
		1,                  // write_flag
		unsafe.Pointer(uintptr(w.id)),
		nil, // read_packet
		CCallback(C.goff_writer_write_packet_trampoline), // write_packet
		CCallback(C.goff_writer_seek_trampoline),         // seek
	)
	return w
}

// Free deallocates all associated resoures with a Reader
func (r *Reader) Free() {
	if r.freed {
		return
	}

	r.freed = true
	freeReaderId(r.id)
	C.av_free(unsafe.Pointer(r.ctx.buffer))
	r.ctx.buffer = nil
	C.avio_context_free(&r.ctx)
}

// Free deallocates all associated resoures with a Reader
func (w *Writer) Free() error {
	if w.freed {
		return nil
	}

	w.freed = true
	freeWriterId(w.id)
	C.av_free(unsafe.Pointer(w.ctx.buffer))
	w.ctx.buffer = nil
	C.avio_context_free(&w.ctx)
	return nil
}

// Error returns the last error of a reader, if any
func (r *Reader) Error() error {
	return r.err
}

// Error returns the last error of a writer, if any
func (w *Writer) Error() error {
	return w.err
}

//export goff_reader_read_packet
func goff_reader_read_packet(opaque unsafe.Pointer, buf unsafe.Pointer, bufSize int) int {
	id := int32(uintptr(opaque))
	p, ok := readers.Load(id)
	if !ok {
		return C.AVERROR_EXTERNAL
	}

	r, ok := (p).(*Reader)
	if !ok {
		return C.AVERROR_EXTERNAL
	}

	if r.err != nil {
		return C.AVERROR_EXTERNAL
	}

	h := reflect.SliceHeader{
		Data: uintptr(buf),
		Cap:  bufSize,
		Len:  bufSize,
	}
	goBuf := *(*[]byte)(unsafe.Pointer(&h))

	readBytes, err := r.inner.Read(goBuf)
	if err != nil {
		if err == io.EOF {
			if readBytes == 0 {
				return C.AVERROR_EOF
			} else {
				return readBytes
			}
		}

		r.err = err
		return C.AVERROR_EXTERNAL
	}

	return readBytes
}

//export goff_reader_seek
func goff_reader_seek(opaque unsafe.Pointer, offset int64, whence int) int64 {
	id := int32(uintptr(opaque))
	p, ok := readers.Load(id)
	if !ok {
		return -1
	}

	r, ok := (p).(*Reader)
	if !ok {
		return -1
	}

	if whence&C.AVSEEK_SIZE > 0 {
		// don't seek, just return size
		return r.size
	}
	// ignore AVSEEK_FORCE

	newOffset, err := r.inner.Seek(offset, io.SeekStart)
	if err != nil {
		r.err = err
		return -1
	}

	return newOffset
}

//export goff_writer_write_packet
func goff_writer_write_packet(opaque unsafe.Pointer, buf unsafe.Pointer, bufSize int) int {
	id := int32(uintptr(opaque))
	p, ok := writers.Load(id)
	if !ok {
		return C.AVERROR_EXTERNAL
	}

	w, ok := (p).(*Writer)
	if !ok {
		return C.AVERROR_EXTERNAL
	}

	if w.err != nil {
		return C.AVERROR_EXTERNAL
	}

	h := reflect.SliceHeader{
		Data: uintptr(buf),
		Cap:  bufSize,
		Len:  bufSize,
	}
	goBuf := *(*[]byte)(unsafe.Pointer(&h))

	writtenBytes, err := w.inner.Write(goBuf)
	if err != nil {
		w.err = err
		return C.AVERROR_EXTERNAL
	}

	return writtenBytes
}

//export goff_writer_seek
func goff_writer_seek(opaque unsafe.Pointer, offset int64, whence int) int64 {
	id := int32(uintptr(opaque))
	p, ok := writers.Load(id)
	if !ok {
		return -1
	}

	w, ok := (p).(*Writer)
	if !ok {
		return -1
	}

	if whence&C.AVSEEK_SIZE > 0 {
		// FIXME: for now, we don't support AVSEEK_SIZE for writers
		return -1
	}
	// ignore AVSEEK_FORCE

	newOffset, err := w.inner.Seek(offset, io.SeekStart)
	if err != nil {
		w.err = err
		return -1
	}

	return newOffset
}

// Use only if used with IOOpen
func (ctx *IOContext) Close() error {
	return CheckErr(C.avio_close(ctx))
}

func (ctx *IOContext) PutStr(s string) error {
	s_ := CString(s)
	defer FreeString(s_)

	return CheckErr(C.avio_put_str(ctx, s_))
}

var log_callback LogCallback
var log_max_level LogLevel

type LogCallback func(ptr uintptr, level LogLevel, line string, printPrefix bool)
type CCallback = *[0]byte

func LogSetCallback(maxLevel LogLevel, lc LogCallback) {
	log_callback = lc
	log_max_level = maxLevel

	if log_callback == nil {
		C.av_log_set_callback(CCallback(C.av_log_default_callback))
	} else {
		C.av_log_set_callback(CCallback(C.goff_log_callback_trampoline))
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
func goff_send_log_to_go(ptr uintptr, level C.int, line *C.char, _printPrefix C.int) {
	printPrefix := false
	if _printPrefix == 1 {
		printPrefix = true
	}

	if log_callback != nil {
		log_callback(uintptr(ptr), LogLevel(level), C.GoString(line), printPrefix)
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

type SearchFlags = C.int

var (
	SearchFlags_CHILDREN SearchFlags = C.AV_OPT_SEARCH_CHILDREN
)

func OptSet(obj unsafe.Pointer, name string, val string, searchFlags SearchFlags) error {
	name_ := CString(name)
	defer FreeString(name_)

	val_ := CString(val)
	defer FreeString(val_)

	return CheckErr(C.av_opt_set(obj, name_, val_, C.int(searchFlags)))
}

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

func (pkt *Packet) Unref() {
	C.av_packet_unref(pkt)
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

type PictureType C.enum_AVPictureType

var (
	// Undefined
	PictureType_None PictureType = C.AV_PICTURE_TYPE_NONE
	// Intra
	PictureType_I PictureType = C.AV_PICTURE_TYPE_I
	// Predicted
	PictureType_P PictureType = C.AV_PICTURE_TYPE_P
	// Bi-dir predicted
	PictureType_B PictureType = C.AV_PICTURE_TYPE_B
	// S(GMC)-VOP MPEG-4
	PictureType_S PictureType = C.AV_PICTURE_TYPE_S
	// Switching Intra.
	PictureType_SI PictureType = C.AV_PICTURE_TYPE_SI
	// Switching Predicted
	PictureType_SP PictureType = C.AV_PICTURE_TYPE_SP
	// BI TYPE
	PictureType_BI PictureType = C.AV_PICTURE_TYPE_BI
)

func (pt PictureType) String() string {
	c := byte(C.av_get_picture_type_char(pt.C()))
	return string([]byte{c})
}

func (pt PictureType) C() C.enum_AVPictureType {
	return C.enum_AVPictureType(pt)
}

type PixelFormat C.enum_AVPixelFormat
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

var (
	PixelFormat_NONE        PixelFormat = C.AV_PIX_FMT_NONE
	PixelFormat_YUV420P     PixelFormat = C.AV_PIX_FMT_YUV420P
	PixelFormat_YUYV422     PixelFormat = C.AV_PIX_FMT_YUYV422
	PixelFormat_RGB24       PixelFormat = C.AV_PIX_FMT_RGB24
	PixelFormat_YUV422P     PixelFormat = C.AV_PIX_FMT_YUV422P
	PixelFormat_YUV444P     PixelFormat = C.AV_PIX_FMT_YUV444P
	PixelFormat_YUV410P     PixelFormat = C.AV_PIX_FMT_YUV410P
	PixelFormat_GRAY8       PixelFormat = C.AV_PIX_FMT_GRAY8
	PixelFormat_MONOWHITE   PixelFormat = C.AV_PIX_FMT_MONOWHITE
	PixelFormat_MONOBLACK   PixelFormat = C.AV_PIX_FMT_MONOBLACK
	PixelFormat_PAL8        PixelFormat = C.AV_PIX_FMT_PAL8
	PixelFormat_YUVJ420P    PixelFormat = C.AV_PIX_FMT_YUVJ420P
	PixelFormat_YUVJ422P    PixelFormat = C.AV_PIX_FMT_YUVJ422P
	PixelFormat_YUVJ444P    PixelFormat = C.AV_PIX_FMT_YUVJ444P
	PixelFormat_UYVY422     PixelFormat = C.AV_PIX_FMT_UYVY422
	PixelFormat_UYYVYY411   PixelFormat = C.AV_PIX_FMT_UYYVYY411
	PixelFormat_BGR8        PixelFormat = C.AV_PIX_FMT_BGR8
	PixelFormat_BGR4        PixelFormat = C.AV_PIX_FMT_BGR4
	PixelFormat_BGR4_BYTE   PixelFormat = C.AV_PIX_FMT_BGR4_BYTE
	PixelFormat_RGB8        PixelFormat = C.AV_PIX_FMT_RGB8
	PixelFormat_RGB4        PixelFormat = C.AV_PIX_FMT_RGB4
	PixelFormat_RGB4_BYTE   PixelFormat = C.AV_PIX_FMT_RGB4_BYTE
	PixelFormat_NV12        PixelFormat = C.AV_PIX_FMT_NV12
	PixelFormat_NV21        PixelFormat = C.AV_PIX_FMT_NV21
	PixelFormat_ARGB        PixelFormat = C.AV_PIX_FMT_ARGB
	PixelFormat_RGBA        PixelFormat = C.AV_PIX_FMT_RGBA
	PixelFormat_ABGR        PixelFormat = C.AV_PIX_FMT_ABGR
	PixelFormat_BGRA        PixelFormat = C.AV_PIX_FMT_BGRA
	PixelFormat_GRAY16BE    PixelFormat = C.AV_PIX_FMT_GRAY16BE
	PixelFormat_GRAY16LE    PixelFormat = C.AV_PIX_FMT_GRAY16LE
	PixelFormat_YUV440P     PixelFormat = C.AV_PIX_FMT_YUV440P
	PixelFormat_YUVJ440P    PixelFormat = C.AV_PIX_FMT_YUVJ440P
	PixelFormat_YUVA420P    PixelFormat = C.AV_PIX_FMT_YUVA420P
	PixelFormat_RGB48BE     PixelFormat = C.AV_PIX_FMT_RGB48BE
	PixelFormat_RGB48LE     PixelFormat = C.AV_PIX_FMT_RGB48LE
	PixelFormat_RGB565BE    PixelFormat = C.AV_PIX_FMT_RGB565BE
	PixelFormat_RGB565LE    PixelFormat = C.AV_PIX_FMT_RGB565LE
	PixelFormat_RGB555BE    PixelFormat = C.AV_PIX_FMT_RGB555BE
	PixelFormat_RGB555LE    PixelFormat = C.AV_PIX_FMT_RGB555LE
	PixelFormat_VAAPI_MOCO  PixelFormat = C.AV_PIX_FMT_VAAPI_MOCO
	PixelFormat_VAAPI_IDCT  PixelFormat = C.AV_PIX_FMT_VAAPI_IDCT
	PixelFormat_VAAPI_VLD   PixelFormat = C.AV_PIX_FMT_VAAPI_VLD
	PixelFormat_VAAPI       PixelFormat = C.AV_PIX_FMT_VAAPI
	PixelFormat_YUV420P16LE PixelFormat = C.AV_PIX_FMT_YUV420P16LE
	PixelFormat_YUV420P16BE PixelFormat = C.AV_PIX_FMT_YUV420P16BE
	PixelFormat_YUV422P16LE PixelFormat = C.AV_PIX_FMT_YUV422P16LE
	PixelFormat_YUV422P16BE PixelFormat = C.AV_PIX_FMT_YUV422P16BE
	PixelFormat_DXVA2_VLD   PixelFormat = C.AV_PIX_FMT_DXVA2_VLD
	PixelFormat_RGB444LE    PixelFormat = C.AV_PIX_FMT_RGB444LE
	PixelFormat_RGB444BE    PixelFormat = C.AV_PIX_FMT_RGB444BE
	PixelFormat_BGR444LE    PixelFormat = C.AV_PIX_FMT_BGR444LE
	PixelFormat_BGR444BE    PixelFormat = C.AV_PIX_FMT_BGR444BE
	PixelFormat_YA8         PixelFormat = C.AV_PIX_FMT_YA8
	PixelFormat_Y400A       PixelFormat = C.AV_PIX_FMT_Y400A
	PixelFormat_GRAY8A      PixelFormat = C.AV_PIX_FMT_GRAY8A
	PixelFormat_BGR48BE     PixelFormat = C.AV_PIX_FMT_BGR48BE
	PixelFormat_BGR48LE     PixelFormat = C.AV_PIX_FMT_BGR48LE
	PixelFormat_YUV420P9BE  PixelFormat = C.AV_PIX_FMT_YUV420P9BE
	PixelFormat_YUV420P9LE  PixelFormat = C.AV_PIX_FMT_YUV420P9LE
	PixelFormat_YUV420P10BE PixelFormat = C.AV_PIX_FMT_YUV420P10BE
	PixelFormat_YUV420P10LE PixelFormat = C.AV_PIX_FMT_YUV420P10LE
	PixelFormat_YUV422P10BE PixelFormat = C.AV_PIX_FMT_YUV422P10BE
	PixelFormat_YUV422P10LE PixelFormat = C.AV_PIX_FMT_YUV422P10LE
	PixelFormat_YUV444P9BE  PixelFormat = C.AV_PIX_FMT_YUV444P9BE
	PixelFormat_YUV444P9LE  PixelFormat = C.AV_PIX_FMT_YUV444P9LE
	PixelFormat_YUV444P10BE PixelFormat = C.AV_PIX_FMT_YUV444P10BE
	PixelFormat_YUV444P10LE PixelFormat = C.AV_PIX_FMT_YUV444P10LE
	PixelFormat_YUV422P9BE  PixelFormat = C.AV_PIX_FMT_YUV422P9BE
	PixelFormat_YUV422P9LE  PixelFormat = C.AV_PIX_FMT_YUV422P9LE

	PixelFormat_GBRP     PixelFormat = C.AV_PIX_FMT_GBRP
	PixelFormat_GBR24P   PixelFormat = C.AV_PIX_FMT_GBR24P
	PixelFormat_GBRP9BE  PixelFormat = C.AV_PIX_FMT_GBRP9BE
	PixelFormat_GBRP9LE  PixelFormat = C.AV_PIX_FMT_GBRP9LE
	PixelFormat_GBRP10BE PixelFormat = C.AV_PIX_FMT_GBRP10BE
	PixelFormat_GBRP10LE PixelFormat = C.AV_PIX_FMT_GBRP10LE
	PixelFormat_GBRP16BE PixelFormat = C.AV_PIX_FMT_GBRP16BE
	PixelFormat_GBRP16LE PixelFormat = C.AV_PIX_FMT_GBRP16LE

	PixelFormat_YUVA422P     PixelFormat = C.AV_PIX_FMT_YUVA422P
	PixelFormat_YUVA444P     PixelFormat = C.AV_PIX_FMT_YUVA444P
	PixelFormat_YUVA420P9BE  PixelFormat = C.AV_PIX_FMT_YUVA420P9BE
	PixelFormat_YUVA420P9LE  PixelFormat = C.AV_PIX_FMT_YUVA420P9LE
	PixelFormat_YUVA422P9BE  PixelFormat = C.AV_PIX_FMT_YUVA422P9BE
	PixelFormat_YUVA422P9LE  PixelFormat = C.AV_PIX_FMT_YUVA422P9LE
	PixelFormat_YUVA444P9BE  PixelFormat = C.AV_PIX_FMT_YUVA444P9BE
	PixelFormat_YUVA444P9LE  PixelFormat = C.AV_PIX_FMT_YUVA444P9LE
	PixelFormat_YUVA420P10BE PixelFormat = C.AV_PIX_FMT_YUVA420P10BE
	PixelFormat_YUVA420P10LE PixelFormat = C.AV_PIX_FMT_YUVA420P10LE
	PixelFormat_YUVA422P10BE PixelFormat = C.AV_PIX_FMT_YUVA422P10BE
	PixelFormat_YUVA422P10LE PixelFormat = C.AV_PIX_FMT_YUVA422P10LE
	PixelFormat_YUVA444P10BE PixelFormat = C.AV_PIX_FMT_YUVA444P10BE
	PixelFormat_YUVA444P10LE PixelFormat = C.AV_PIX_FMT_YUVA444P10LE
	PixelFormat_YUVA420P16BE PixelFormat = C.AV_PIX_FMT_YUVA420P16BE
	PixelFormat_YUVA420P16LE PixelFormat = C.AV_PIX_FMT_YUVA420P16LE
	PixelFormat_YUVA422P16BE PixelFormat = C.AV_PIX_FMT_YUVA422P16BE
	PixelFormat_YUVA422P16LE PixelFormat = C.AV_PIX_FMT_YUVA422P16LE
	PixelFormat_YUVA444P16BE PixelFormat = C.AV_PIX_FMT_YUVA444P16BE
	PixelFormat_YUVA444P16LE PixelFormat = C.AV_PIX_FMT_YUVA444P16LE

	PixelFormat_VDPAU   PixelFormat = C.AV_PIX_FMT_VDPAU
	PixelFormat_XYZ12LE PixelFormat = C.AV_PIX_FMT_XYZ12LE
	PixelFormat_XYZ12BE PixelFormat = C.AV_PIX_FMT_XYZ12BE

	PixelFormat_NV16        PixelFormat = C.AV_PIX_FMT_NV16
	PixelFormat_NV20LE      PixelFormat = C.AV_PIX_FMT_NV20LE
	PixelFormat_NV20BE      PixelFormat = C.AV_PIX_FMT_NV20BE
	PixelFormat_RGBA64BE    PixelFormat = C.AV_PIX_FMT_RGBA64BE
	PixelFormat_RGBA64LE    PixelFormat = C.AV_PIX_FMT_RGBA64LE
	PixelFormat_BGRA64BE    PixelFormat = C.AV_PIX_FMT_BGRA64BE
	PixelFormat_BGRA64LE    PixelFormat = C.AV_PIX_FMT_BGRA64LE
	PixelFormat_YVYU422     PixelFormat = C.AV_PIX_FMT_YVYU422
	PixelFormat_YA16BE      PixelFormat = C.AV_PIX_FMT_YA16BE
	PixelFormat_YA16LE      PixelFormat = C.AV_PIX_FMT_YA16LE
	PixelFormat_GBRAP       PixelFormat = C.AV_PIX_FMT_GBRAP
	PixelFormat_GBRAP16BE   PixelFormat = C.AV_PIX_FMT_GBRAP16BE
	PixelFormat_GBRAP16LE   PixelFormat = C.AV_PIX_FMT_GBRAP16LE
	PixelFormat_QSV         PixelFormat = C.AV_PIX_FMT_QSV
	PixelFormat_MMAL        PixelFormat = C.AV_PIX_FMT_MMAL
	PixelFormat_D3D11VA_VLD PixelFormat = C.AV_PIX_FMT_D3D11VA_VLD
	PixelFormat_CUDA        PixelFormat = C.AV_PIX_FMT_CUDA
	PixelFormat_0RGB        PixelFormat = C.AV_PIX_FMT_0RGB
	PixelFormat_RGB0        PixelFormat = C.AV_PIX_FMT_RGB0
	PixelFormat_0BGR        PixelFormat = C.AV_PIX_FMT_0BGR
	PixelFormat_BGR0        PixelFormat = C.AV_PIX_FMT_BGR0

	PixelFormat_YUV420P12BE PixelFormat = C.AV_PIX_FMT_YUV420P12BE
	PixelFormat_YUV420P12LE PixelFormat = C.AV_PIX_FMT_YUV420P12LE
	PixelFormat_YUV420P14BE PixelFormat = C.AV_PIX_FMT_YUV420P14BE
	PixelFormat_YUV420P14LE PixelFormat = C.AV_PIX_FMT_YUV420P14LE
	PixelFormat_YUV422P12BE PixelFormat = C.AV_PIX_FMT_YUV422P12BE
	PixelFormat_YUV422P12LE PixelFormat = C.AV_PIX_FMT_YUV422P12LE
	PixelFormat_YUV422P14BE PixelFormat = C.AV_PIX_FMT_YUV422P14BE
	PixelFormat_YUV422P14LE PixelFormat = C.AV_PIX_FMT_YUV422P14LE
	PixelFormat_YUV444P12BE PixelFormat = C.AV_PIX_FMT_YUV444P12BE
	PixelFormat_YUV444P12LE PixelFormat = C.AV_PIX_FMT_YUV444P12LE
	PixelFormat_YUV444P14BE PixelFormat = C.AV_PIX_FMT_YUV444P14BE
	PixelFormat_YUV444P14LE PixelFormat = C.AV_PIX_FMT_YUV444P14LE

	PixelFormat_GBRP12BE PixelFormat = C.AV_PIX_FMT_GBRP12BE
	PixelFormat_GBRP12LE PixelFormat = C.AV_PIX_FMT_GBRP12LE
	PixelFormat_GBRP14BE PixelFormat = C.AV_PIX_FMT_GBRP14BE
	PixelFormat_GBRP14LE PixelFormat = C.AV_PIX_FMT_GBRP14LE

	PixelFormat_YUVJ411P PixelFormat = C.AV_PIX_FMT_YUVJ411P

	PixelFormat_BAYER_BGGR8    PixelFormat = C.AV_PIX_FMT_BAYER_BGGR8
	PixelFormat_BAYER_RGGB8    PixelFormat = C.AV_PIX_FMT_BAYER_RGGB8
	PixelFormat_BAYER_GBRG8    PixelFormat = C.AV_PIX_FMT_BAYER_GBRG8
	PixelFormat_BAYER_GRBG8    PixelFormat = C.AV_PIX_FMT_BAYER_GRBG8
	PixelFormat_BAYER_BGGR16LE PixelFormat = C.AV_PIX_FMT_BAYER_BGGR16LE
	PixelFormat_BAYER_BGGR16BE PixelFormat = C.AV_PIX_FMT_BAYER_BGGR16BE
	PixelFormat_BAYER_RGGB16LE PixelFormat = C.AV_PIX_FMT_BAYER_RGGB16LE
	PixelFormat_BAYER_RGGB16BE PixelFormat = C.AV_PIX_FMT_BAYER_RGGB16BE
	PixelFormat_BAYER_GBRG16LE PixelFormat = C.AV_PIX_FMT_BAYER_GBRG16LE
	PixelFormat_BAYER_GBRG16BE PixelFormat = C.AV_PIX_FMT_BAYER_GBRG16BE
	PixelFormat_BAYER_GRBG16LE PixelFormat = C.AV_PIX_FMT_BAYER_GRBG16LE
	PixelFormat_BAYER_GRBG16BE PixelFormat = C.AV_PIX_FMT_BAYER_GRBG16BE

	PixelFormat_XVMC PixelFormat = C.AV_PIX_FMT_XVMC

	PixelFormat_YUV440P10LE PixelFormat = C.AV_PIX_FMT_YUV440P10LE
	PixelFormat_YUV440P10BE PixelFormat = C.AV_PIX_FMT_YUV440P10BE
	PixelFormat_YUV440P12LE PixelFormat = C.AV_PIX_FMT_YUV440P12LE
	PixelFormat_YUV440P12BE PixelFormat = C.AV_PIX_FMT_YUV440P12BE

	PixelFormat_AYUV64LE PixelFormat = C.AV_PIX_FMT_AYUV64LE
	PixelFormat_AYUV64BE PixelFormat = C.AV_PIX_FMT_AYUV64BE

	PixelFormat_VIDEOTOOLBOX PixelFormat = C.AV_PIX_FMT_VIDEOTOOLBOX

	PixelFormat_P010LE PixelFormat = C.AV_PIX_FMT_P010LE
	PixelFormat_P010BE PixelFormat = C.AV_PIX_FMT_P010BE

	PixelFormat_GBRAP12BE PixelFormat = C.AV_PIX_FMT_GBRAP12BE
	PixelFormat_GBRAP12LE PixelFormat = C.AV_PIX_FMT_GBRAP12LE
	PixelFormat_GBRAP10BE PixelFormat = C.AV_PIX_FMT_GBRAP10BE
	PixelFormat_GBRAP10LE PixelFormat = C.AV_PIX_FMT_GBRAP10LE

	PixelFormat_MEDIACODEC PixelFormat = C.AV_PIX_FMT_MEDIACODEC

	PixelFormat_GRAY12BE PixelFormat = C.AV_PIX_FMT_GRAY12BE
	PixelFormat_GRAY12LE PixelFormat = C.AV_PIX_FMT_GRAY12LE
	PixelFormat_GRAY10BE PixelFormat = C.AV_PIX_FMT_GRAY10BE
	PixelFormat_GRAY10LE PixelFormat = C.AV_PIX_FMT_GRAY10LE

	PixelFormat_P016LE PixelFormat = C.AV_PIX_FMT_P016LE
	PixelFormat_P016BE PixelFormat = C.AV_PIX_FMT_P016BE

	PixelFormat_D3D11 PixelFormat = C.AV_PIX_FMT_D3D11

	PixelFormat_GRAY9BE PixelFormat = C.AV_PIX_FMT_GRAY9BE
	PixelFormat_GRAY9LE PixelFormat = C.AV_PIX_FMT_GRAY9LE

	PixelFormat_GBRPF32BE  PixelFormat = C.AV_PIX_FMT_GBRPF32BE
	PixelFormat_GBRPF32LE  PixelFormat = C.AV_PIX_FMT_GBRPF32LE
	PixelFormat_GBRAPF32BE PixelFormat = C.AV_PIX_FMT_GBRAPF32BE
	PixelFormat_GBRAPF32LE PixelFormat = C.AV_PIX_FMT_GBRAPF32LE

	PixelFormat_DRM_PRIME PixelFormat = C.AV_PIX_FMT_DRM_PRIME

	PixelFormat_GRAY14BE  PixelFormat = C.AV_PIX_FMT_GRAY14BE
	PixelFormat_GRAY14LE  PixelFormat = C.AV_PIX_FMT_GRAY14LE
	PixelFormat_GRAYF32BE PixelFormat = C.AV_PIX_FMT_GRAYF32BE
	PixelFormat_GRAYF32LE PixelFormat = C.AV_PIX_FMT_GRAYF32LE
)

type (
	SampleFormat C.enum_AVSampleFormat
)

func (sf SampleFormat) C() C.enum_AVSampleFormat {
	return C.enum_AVSampleFormat(sf)
}

var (
	SampleFormat_NONE SampleFormat = C.AV_SAMPLE_FMT_NONE
	// unsigned 8 bits
	SampleFormat_U8 SampleFormat = C.AV_SAMPLE_FMT_U8
	// signed 16 bits
	SampleFormat_S16 SampleFormat = C.AV_SAMPLE_FMT_S16
	// signed 32 bits
	SampleFormat_S32 SampleFormat = C.AV_SAMPLE_FMT_S32
	// float
	SampleFormat_FLT SampleFormat = C.AV_SAMPLE_FMT_FLT
	// double
	SampleFormat_DBL SampleFormat = C.AV_SAMPLE_FMT_DBL
	// unsigned 8 bits, planar
	SampleFormat_U8P SampleFormat = C.AV_SAMPLE_FMT_U8P
	// signed 16 bits, planar
	SampleFormat_S16P SampleFormat = C.AV_SAMPLE_FMT_S16P
	// signed 32 bits, planar
	SampleFormat_S32P SampleFormat = C.AV_SAMPLE_FMT_S32P
	// float, planar
	SampleFormat_FLTP SampleFormat = C.AV_SAMPLE_FMT_FLTP
	// double, planar
	SampleFormat_DBLP SampleFormat = C.AV_SAMPLE_FMT_DBLP
	// signed 64 bits
	SampleFormat_S64 SampleFormat = C.AV_SAMPLE_FMT_S64
	// signed 64 bits, planar
	SampleFormat_S64P SampleFormat = C.AV_SAMPLE_FMT_S64P
)

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

func (s *Stream) AvgFrameRate() Rational {
	return s.avg_frame_rate
}

func (s *Stream) CodecParameters() *CodecParameters {
	return s.codecpar
}

func (s *Stream) SetCodecParameters(cp *CodecParameters) {
	s.codecpar = cp
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

func (swctx *SwsContext) Free() {
	C.sws_freeContext(swctx)
}

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

func (r Rational) Mul(r2 Rational) Rational {
	return C.av_mul_q(r, r2)
}

func (r Rational) Div(r2 Rational) Rational {
	return C.av_div_q(r, r2)
}
