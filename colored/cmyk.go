package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewCMYK returns a new 800×418 CMYK image that has been pre-colored with the provided color.
func NewCMYK(c color.Color) *image.CMYK {
	img := img800x418.NewCMYK()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
