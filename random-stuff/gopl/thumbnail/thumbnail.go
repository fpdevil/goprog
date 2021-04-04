package thumbnail

// The thumnail package produces a thumbnail size of image
// from larger image provided.

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"io"
)

// Image function returns a thumnail sized version of input
func Image(src image.Image) image.Image {
	// calculate the thumbnail size, preserving the aspect ratio
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // make portrait
	} else {
		height = int(128 / aspect) // make landscape
	}

	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	// scaling
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and writes a thumbnailed
// version of the same to w
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile reads an image from in and writes a thumbnail sized
// version of the same to the same directory.
// It returns the generated file name eg., "test.thumb.jpeg"
func ImageFile(in string) (string, error) {
	ext := filepath.Ext(in)
	out := strings.TrimSuffix(in, ext) + ".thumb" + ext
	return out, ImageFileX(out, in)
}

// ImageFileX reads an image from in and writes a thumbnail
// sized version of the same to out
func ImageFileX(out, in string) (err error) {
	infile, err := os.Open(in)
	if err != nil {
		return
	}
	defer infile.Close()

	outfile, err := os.Create(out)
	if err != nil {
		return
	}

	if err := ImageStream(outfile, infile); err != nil {
		outfile.Close()
		return fmt.Errorf("scaling %s to %s: %s", in, out, err)
	}
	return outfile.Close()
}
