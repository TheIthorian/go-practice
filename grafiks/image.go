package main

import (
	"fmt"
	"image"
	"image/color"
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
			color := color.GrayModel.Convert(resizedImage.At(x, y)).(color.Gray)
			buffer[x+y*width] = color.Y
		}
	}

	return buffer
}
