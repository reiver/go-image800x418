package image800x418

import (
	"image"
)

// NewRGBA64 returns a new 800×418 RGBA64 image.
func NewRGBA64() *image.RGBA64 {
	return image.NewRGBA64(Bounds())
}
