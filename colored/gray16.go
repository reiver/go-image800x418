package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewGray16 returns a new 800×418 Gray16 image that has been pre-colored with the provided color.
func NewGray16(c color.Color) *image.Gray16 {
	img := img800x418.NewGray16()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
