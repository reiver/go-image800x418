package image800x418

import (
	"image"
	"image/color"
)

// NewPaletted returns a new 800×418 Paletted image using the provided palette.
func NewPaletted(palette color.Palette) *image.Paletted {
	return image.NewPaletted(Bounds(), palette)
}
