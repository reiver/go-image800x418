package image800x418

import (
	"image"
)

// NewCMYK returns a new 800×418 CMYK image.
func NewCMYK() *image.CMYK {
	return image.NewCMYK(Bounds())
}
