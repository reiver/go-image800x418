package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewRGBA returns a new 800×418 RGBA image that has been pre-colored with the provided color.
func NewRGBA(c color.Color) *image.RGBA {
	img := img800x418.NewRGBA()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
