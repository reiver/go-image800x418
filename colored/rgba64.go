package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewRGBA64 returns a new 800×418 RGBA64 image that has been pre-colored with the provided color.
func NewRGBA64(c color.Color) *image.RGBA64 {
	img := img800x418.NewRGBA64()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
