package image800x418

import (
	"image"
)

// NewNRGBA returns a new 800×418 NRGBA image.
func NewNRGBA() *image.NRGBA {
	return image.NewNRGBA(Bounds())
}
