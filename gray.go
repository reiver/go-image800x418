package image800x418

import (
	"image"
)

// NewGray returns a new 800×418 Gray image.
func NewGray() *image.Gray {
	return image.NewGray(Bounds())
}
