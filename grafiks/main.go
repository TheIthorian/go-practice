package main

import (
	"fmt"

	"golang.org/x/term"
)

var (
	WIDTH  int
	HEIGHT int
)

type point struct {
	X int
	Y int
}

type display struct {
	Buffer [][]rune
}

func makeDisplay(width, height int) display {
	var buffer [][]rune
	for i := 0; i < height; i++ {
		buffer = append(buffer, make([]rune, height))
	}

	return display{Buffer: buffer}
}

func coordToIndex(p *point) int {
	return p.Y*HEIGHT + p.X
}

func indexToCoord(index int) point {
	X := index % HEIGHT
	Y := (index) / WIDTH
	return point{X, Y}
}

func (d *display) drawPoint(p *point, a int) {
	d.Buffer[p.Y][p.X] = '#'
}

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

func (d *display) render() {
	for _, row := range d.Buffer {
		fmt.Println(toString(row))
	}
}

func drawAlternatingPoints(d *display) {
	for i := 0; i < WIDTH*HEIGHT; i += 2 {
		p := indexToCoord(i)
		d.drawPoint(&p, 255)
	}
}

func main() {
	width, height, err := term.GetSize(0)
	fmt.Printf("The terminal is of size: %v, %v ", width, height, err)
	myDisplay := makeDisplay(width, height)
	HEIGHT = height
	WIDTH = width

	drawAlternatingPoints(&myDisplay)
	myDisplay.render()
}
