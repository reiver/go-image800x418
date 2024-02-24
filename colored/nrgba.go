package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewNRGBA returns a new 800×418 NRGBA image that has been pre-colored with the provided color.
func NewNRGBA(c color.Color) *image.NRGBA {
	img := img800x418.NewNRGBA()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
