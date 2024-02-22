package image800x418

import (
	"image"
)

// NewNYCbCrA returns a new 800×418 NYCbCrA image with the given subsample ratio.
func NewNYCbCrA(subsampleRatio image.YCbCrSubsampleRatio) *image.NYCbCrA {
	return image.NewNYCbCrA(Bounds(), subsampleRatio)
}
