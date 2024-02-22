# go-img800x418

Package **img800x418** provides tools for working with images whose width and height are 800×418, for the Go programming language.

800x418 images are popular with what are something called "cards" or "link preview images".
I.e., the images used with OpenGraph `"og:image"`, Twitter Card `"twitter:image"`, and Farcaster Frames (i.e., Frame Protocol) `"fc:frame:image"`.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-img800x418

[![GoDoc](https://godoc.org/github.com/reiver/go-img800x418?status.svg)](https://godoc.org/github.com/reiver/go-img800x418)

## Examples

Here is an example of creating an 800×418 image using package **img800x418**.
In this particular example a paletted image is created.

```golang
import "github.com/reiver/go-img800x418"

// ...

img := img800x418.NewPaletted(palette)
```

## Import

To import package **img800x418** use `import` code like the follownig:
```
import "github.com/reiver/go-img800x418"
```

## Installation

To install package **img800x418** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-img800x418
```

## Author

Package **img800x418** was written by [Charles Iliya Krempeaux](http://changelog.ca)
