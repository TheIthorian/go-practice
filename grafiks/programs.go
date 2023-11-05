package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

func toString(r []rune) string {
	str := ""
	for _, s := range r {
		if s == 0 {
			s = ' '
		}
		str += string(s)
	}
	return str
}

func getDisplay() display {
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Println("Unable to load terminal")
		return makeDisplay(0, 0)
	}

	return makeDisplay(width, height)
}

func renderImage(sourceImagePath string) {
	display := getDisplay()
	buffer := getImageBuffer(sourceImagePath, display.width, display.height)

	for X := 0; X < display.width; X += 1 {
		for Y := 0; Y < display.height; Y += 1 {
			p := Point{X, Y}
			display.drawPoint(&p, int(buffer[X+Y*display.width]))
		}
	}

	display.render()
}

func alternatingPoints() {
	myDisplay := getDisplay()
	for X := 0; X < myDisplay.width; X += 2 {
		for Y := 0; Y < myDisplay.height; Y += 2 {
			p := Point{X, Y}
			myDisplay.drawPoint(&p, 255)
		}
	}
	myDisplay.render()
}

func bouncingPoint() {
	myDisplay := getDisplay()
	myDisplay.render()
	sim := makeSim()

	for {
		time.Sleep(time.Second / 5)
		sim.step(0.5)

		myDisplay.clear()

		for _, particle := range sim.particles {
			myDisplay.drawPoint(
				&Point{X: int(particle.position.X), Y: int(particle.position.Y)},
				255,
			)
		}

		myDisplay.render()
	}
}
