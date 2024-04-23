package vips

// #include "foreign.h"
import "C"
import (
	"unsafe"
)

// JpegExportParams are options when exporting a JPEG to file or buffer
type JpegExportParams struct {
	StripMetadata      bool
	Quality            int
	Interlace          bool
	OptimizeCoding     bool
	SubsampleMode      SubsampleMode
	TrellisQuant       bool
	OvershootDeringing bool
	OptimizeScans      bool
	QuantTable         int
}

// NewJpegExportParams creates default values for an export of a JPEG image.
// By default, vips creates interlaced JPEGs with a quality of 80/100.
func NewJpegExportParams() *JpegExportParams {
	return &JpegExportParams{
		Quality:   80,
		Interlace: true,
	}
}

// PngExportParams are options when exporting a PNG to file or buffer
type PngExportParams struct {
	StripMetadata bool
	Compression   int
	Filter        PngFilter
	Interlace     bool
	Quality       int
	Palette       bool
	Dither        float64
	Bitdepth      int
	Profile       string // TODO: Use this param during save
}

// NewPngExportParams creates default values for an export of a PNG image.
// By default, vips creates non-interlaced PNGs with a compression of 6/10.
func NewPngExportParams() *PngExportParams {
	return &PngExportParams{
		Compression: 6,
		Filter:      PngFilterNone,
		Interlace:   false,
		Palette:     false,
	}
}

// WebpExportParams are options when exporting a WEBP to file or buffer
type WebpExportParams struct {
	StripMetadata   bool
	Quality         int
	Lossless        bool
	NearLossless    bool
	ReductionEffort int
	IccProfile      string
}

// NewWebpExportParams creates default values for an export of a WEBP image.
// By default, vips creates lossy images with a quality of 75/100.
func NewWebpExportParams() *WebpExportParams {
	return &WebpExportParams{
		Quality:         75,
		Lossless:        false,
		NearLossless:    false,
		ReductionEffort: 4,
	}
}

// HeifExportParams are options when exporting a HEIF to file or buffer
type HeifExportParams struct {
	Quality  int
	Lossless bool
}

// NewHeifExportParams creates default values for an export of a HEIF image.
func NewHeifExportParams() *HeifExportParams {
	return &HeifExportParams{
		Quality:  80,
		Lossless: false,
	}
}

// TiffExportParams are options when exporting a TIFF to file or buffer
type TiffExportParams struct {
	StripMetadata bool
	Quality       int
	Compression   TiffCompression
	Predictor     TiffPredictor
}

// NewTiffExportParams creates default values for an export of a TIFF image.
func NewTiffExportParams() *TiffExportParams {
	return &TiffExportParams{
		Quality:     80,
		Compression: TiffCompressionLzw,
		Predictor:   TiffPredictorHorizontal,
	}
}

// GifExportParams  are options when exporting an GIF to file or buffer.
type GifExportParams struct {
	StripMetadata bool
	// 使用 quality 设置输出 GIF 的质量（1-100 之间的整数）。
	Quality int
	// 使用 dither 设置 Floyd-Steinberg 抖动程度
	Dither float64
	// 使用 effort 控制 CPU 工作量（1 是最快，10 是最慢，7 是默认值）。
	Effort int
	// 使用 bitdepth 设置输出 GIF 的位深度（1-8 之间的整数）。
	Bitdepth int
	// 使用 interframe_maxerror 设置阈值，低于该阈值的像素被视为相等。帧与帧之间不发生变化的像素可以变得透明，从而提高压缩率。默认 0。
	InterframeMaxerror int
	// 如果 reuse 为 TRUE，则 GIF 将使用从 in 中的元数据获取的单个全局调色板保存，并且不会进行新的调色板优化。
	Reuse bool
	// 如果 interlace 为 TRUE，则 GIF 文件将为隔行扫描（逐行 GIF）。这些文件可能更适合通过慢速网络连接显示，但需要更多内存进行编码。
	Interlace bool
	// 使用 interpalette_maxerror 设置阈值，低于该阈值将重用先前生成的调色板。
	InterpaletteMaxerror int
}

// NewGifExportParams creates default values for an export of a GIF image.
func NewGifExportParams() *GifExportParams {
	return &GifExportParams{
		Quality:  75,
		Effort:   7,
		Bitdepth: 8,
	}
}

// AvifExportParams are options when exporting an AVIF to file or buffer.
type AvifExportParams struct {
	StripMetadata bool
	Quality       int
	Lossless      bool
	Speed         int
}

// NewAvifExportParams creates default values for an export of an AVIF image.
func NewAvifExportParams() *AvifExportParams {
	return &AvifExportParams{
		Quality:  80,
		Lossless: false,
		Speed:    5,
	}
}

// Jp2kExportParams are options when exporting an JPEG2000 to file or buffer.
type Jp2kExportParams struct {
	Quality       int
	Lossless      bool
	TileWidth     int
	TileHeight    int
	SubsampleMode SubsampleMode
}

// NewJp2kExportParams creates default values for an export of an JPEG2000 image.
func NewJp2kExportParams() *Jp2kExportParams {
	return &Jp2kExportParams{
		Quality:    80,
		Lossless:   false,
		TileWidth:  512,
		TileHeight: 512,
	}
}

func vipsSaveJPEGToBuffer(in *C.VipsImage, params JpegExportParams) ([]byte, error) {
	p := C.create_save_params(C.JPEG)
	p.inputImage = in
	p.stripMetadata = C.int(boolToInt(params.StripMetadata))
	p.quality = C.int(params.Quality)
	p.interlace = C.int(boolToInt(params.Interlace))
	p.jpegOptimizeCoding = C.int(boolToInt(params.OptimizeCoding))
	p.jpegSubsample = C.VipsForeignJpegSubsample(params.SubsampleMode)
	p.jpegTrellisQuant = C.int(boolToInt(params.TrellisQuant))
	p.jpegOvershootDeringing = C.int(boolToInt(params.OvershootDeringing))
	p.jpegOptimizeScans = C.int(boolToInt(params.OptimizeScans))
	p.jpegQuantTable = C.int(params.QuantTable)

	return vipsSaveToBuffer(p)
}

func vipsSavePNGToBuffer(in *C.VipsImage, params PngExportParams) ([]byte, error) {
	p := C.create_save_params(C.PNG)
	p.inputImage = in
	p.quality = C.int(params.Quality)
	p.stripMetadata = C.int(boolToInt(params.StripMetadata))
	p.interlace = C.int(boolToInt(params.Interlace))
	p.pngCompression = C.int(params.Compression)
	p.pngFilter = C.VipsForeignPngFilter(params.Filter)
	p.pngPalette = C.int(boolToInt(params.Palette))
	p.pngDither = C.double(params.Dither)
	p.pngBitdepth = C.int(params.Bitdepth)

	return vipsSaveToBuffer(p)
}

func vipsSaveWebPToBuffer(in *C.VipsImage, params WebpExportParams) ([]byte, error) {
	p := C.create_save_params(C.WEBP)
	p.inputImage = in
	p.stripMetadata = C.int(boolToInt(params.StripMetadata))
	p.quality = C.int(params.Quality)
	p.webpLossless = C.int(boolToInt(params.Lossless))
	p.webpNearLossless = C.int(boolToInt(params.NearLossless))
	p.webpReductionEffort = C.int(params.ReductionEffort)

	if params.IccProfile != "" {
		p.webpIccProfile = C.CString(params.IccProfile)
		defer C.free(unsafe.Pointer(p.webpIccProfile))
	}

	return vipsSaveToBuffer(p)
}

func vipsSaveTIFFToBuffer(in *C.VipsImage, params TiffExportParams) ([]byte, error) {
	p := C.create_save_params(C.TIFF)
	p.inputImage = in
	p.stripMetadata = C.int(boolToInt(params.StripMetadata))
	p.quality = C.int(params.Quality)
	p.tiffCompression = C.VipsForeignTiffCompression(params.Compression)

	return vipsSaveToBuffer(p)
}

func vipsSaveHEIFToBuffer(in *C.VipsImage, params HeifExportParams) ([]byte, error) {
	p := C.create_save_params(C.HEIF)
	p.inputImage = in
	p.outputFormat = C.HEIF
	p.quality = C.int(params.Quality)
	p.heifLossless = C.int(boolToInt(params.Lossless))

	return vipsSaveToBuffer(p)
}

func vipsSaveAVIFToBuffer(in *C.VipsImage, params AvifExportParams) ([]byte, error) {
	p := C.create_save_params(C.AVIF)
	p.inputImage = in
	p.outputFormat = C.AVIF
	p.quality = C.int(params.Quality)
	p.heifLossless = C.int(boolToInt(params.Lossless))
	p.avifSpeed = C.int(params.Speed)

	return vipsSaveToBuffer(p)
}

func vipsSaveJP2KToBuffer(in *C.VipsImage, params Jp2kExportParams) ([]byte, error) {
	p := C.create_save_params(C.JP2K)
	p.inputImage = in
	p.outputFormat = C.JP2K
	p.quality = C.int(params.Quality)
	p.jp2kLossless = C.int(boolToInt(params.Lossless))
	p.jp2kTileWidth = C.int(params.TileWidth)
	p.jp2kTileHeight = C.int(params.TileHeight)
	p.jpegSubsample = C.VipsForeignJpegSubsample(params.SubsampleMode)

	return vipsSaveToBuffer(p)
}

func vipsSaveGIFToBuffer(in *C.VipsImage, params GifExportParams) ([]byte, error) {
	p := C.create_save_params(C.GIF)
	p.inputImage = in
	p.quality = C.int(params.Quality)
	p.gifDither = C.double(params.Dither)
	p.gifEffort = C.int(params.Effort)
	p.gifBitdepth = C.int(params.Bitdepth)
	p.gitInterframeMaxerror = C.int(params.InterframeMaxerror)
	p.gitReuse = C.int(boolToInt(params.Reuse))
	p.gitInterlace = C.int(boolToInt(params.Interlace))
	p.gitInterpaletteMaxerror = C.int(params.InterpaletteMaxerror)

	return vipsSaveToBuffer(p)
}

func vipsSaveToBuffer(params C.struct_SaveParams) ([]byte, error) {
	if err := C.save_to_buffer(&params); err != 0 {
		if params.outputBuffer != nil {
			gFreePointer(params.outputBuffer)
		}
		return nil, handleVipsError()
	}

	buf := C.GoBytes(params.outputBuffer, C.int(params.outputLen))
	defer gFreePointer(params.outputBuffer)

	return buf, nil
}
