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
