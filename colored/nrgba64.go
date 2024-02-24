package img800x418colored

import (
	"image"
	"image/color"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// NewNRGBA64 returns a new 800×418 NRGBA64 image that has been pre-colored with the provided color..
func NewNRGBA64(c color.Color) *image.NRGBA64 {
	img := img800x418.NewNRGBA64()
        if nil == img {
                return nil
        }

        imgdye.Dye(img, c)

        return img
}
