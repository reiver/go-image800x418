package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewPaletted returns a new 800×418 Paletted image using the provided palette that has been pre-colored with the provided color.
func NewPaletted(c color.Color, palette color.Palette) *image.Paletted {
	img := img800x418.NewPaletted(palette)
	if nil == img {
		return nil
	}

	imgdye.Dye(img, c)

	return img
}

// NewPaletted returns a new 800×418 Paletted image using a palette of just one color that has been pre-colored with the provided color.
func NewPalettedSolo(c color.Color) *image.Paletted {
	var palette color.Palette = color.Palette{
		c,
	}

	return NewPaletted(c, palette)
}
