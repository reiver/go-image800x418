package img800x418

import (
	"image"
)

// NewNRGBA64 returns a new 800×418 NRGBA64 image.
func NewNRGBA64() *image.NRGBA64 {
	return image.NewNRGBA64(Bounds())
}
