package img800x418

import (
	"image"
)

// NewGray16 returns a new 800×418 Gray16 image.
func NewGray16() *image.Gray16 {
	return image.NewGray16(Bounds())
}
