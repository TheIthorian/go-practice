package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func getImageBuffer(filePath string, width int, height int) []uint8 {
	reader, err := os.Open(filePath)
	defer reader.Close()
	if err != nil {
		log.Fatal(err)
	}

	myImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	resizedImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(
		resizedImage,
		resizedImage.Rect,
		myImage,
		myImage.Bounds(),
		draw.Over,
		nil,
	)

	buffer := make([]uint8, width*height)

	bounds := resizedImage.Bounds()
	fmt.Println("Resized image to", bounds.Max.X, bounds.Max.Y)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := resizedImage.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 8 reduces this to the range [0, 255].
			grayscaleFactor := uint(0.3*float32(r) + 0.6*float32(g) + 0.1*float32(b))
			fmt.Println(grayscaleFactor >> 8)

			buffer[x+y*width] = uint8(grayscaleFactor >> 8)
		}
	}

	return buffer
}
