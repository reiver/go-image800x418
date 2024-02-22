package img800x418

import (
	"image"
)

// NewRGBA returns a new 800×418 RGBA image.
func NewRGBA() *image.RGBA {
	return image.NewRGBA(Bounds())
}
