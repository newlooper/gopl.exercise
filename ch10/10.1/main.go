package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "", "Output formtat: png|jpg|gif")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "usage: program -format=<png|jpg|gif> < INPUT > OUTPUT")
		os.Exit(1)
	}

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	format = strings.ToLower(format)
	switch format {
	case "jpg", "jpeg":
		err = toJPEG(img, os.Stdout)
	case "png":
		err = toPNG(img, os.Stdout)
	case "gif":
		err = toGIF(img, os.Stdout)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, nil)
}
