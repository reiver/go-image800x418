package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewGray returns a new 800×418 Gray image that has been pre-colored with the provided color.
func NewGray(c color.Color) *image.Gray {
	img := img800x418.NewGray()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
