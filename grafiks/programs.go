package main

import (
	"fmt"
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
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println("Unable to load terminal")
		return makeDisplay(0, 0)
	}

	return makeDisplay(width, height)
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
