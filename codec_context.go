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
import "unsafe"

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
