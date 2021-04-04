package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "jpg", "provide the output image format")
	flag.Parse()
	p := flag.Args()[0]
	file, err := os.Open(p)
	if err != nil {
		log.Fatalf("failed to open %s: %v", file.Name(), err)
	}

	img, kind, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode %s: %v", p, err)
	}

	log.Printf("input image format is %v and intended format is %v", kind, format)
	if kind == format {
		log.Fatalf("%s is already in %s format\n", p, kind)
	}

	outfile := p[:len(p)-len(filepath.Ext(p))] + "." + format
	out, err := os.Create(outfile)
	if err != nil {
		log.Printf("failed to create %s: %v", outfile, err)
		return
	}

	switch format {
	case "jpeg":
		err = toJPEG(img, out)
	case "png":
		err = toPNG(img, out)
	case "gif":
		err = toGIF(img, out)
	}
	if err != nil {
		log.Fatalf("error %s: %v", kind, err)
	}

	file.Close()
	out.Close()
}

//!+ toJPEG

// toJPEG converts the image into a jpg/jpeg format
func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//!- toJPEG

//!+ toPNG

// toPNG converts the input image ti png format
func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

//!- toPNG

//!+ toGIF

// toGIF function converts the input image to gif format
func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, nil)
}

//!- toGIF
