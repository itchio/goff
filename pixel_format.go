package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/pixdesc.h>
import "C"

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
