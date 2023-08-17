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

func drawAlternatingPoints(d *display) {
	for X := 0; X < d.width; X += 1 {
		for Y := 0; Y < d.height; Y += 1 {
			fmt.Println(X, Y)
			p := Point{X, Y}
			d.drawPoint(&p, 255)
		}
	}
}

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println("Unable to load terminal")
	}

	fmt.Printf("The terminal is of size: %v, %v\n", width, height)

	myDisplay := makeDisplay(width, height)
	// drawAlternatingPoints(&myDisplay)
	myDisplay.render()

	sim := makeSim()

	for {
		time.Sleep(time.Second / 5)
		sim.step(0.5)

		myDisplay.clear()

		for _, particle := range sim.particles {
			//	println(particle.position.X, particle.position.Y)
			myDisplay.drawPoint(
				&Point{X: int(particle.position.X), Y: int(particle.position.Y)},
				255,
			)
		}

		myDisplay.render()
	}
}
