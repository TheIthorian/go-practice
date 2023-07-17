package main

import (
	"fmt"
)

const (
	WIDTH  int = 11
	HEIGHT int = 11
)

type point struct {
	X int
	Y int
}

type display struct {
	Buffer [HEIGHT][WIDTH]rune
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

func toString(r [WIDTH]rune) string {
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
	fmt.Println("Hello chloe")
	myDisplay := display{}

	drawAlternatingPoints(&myDisplay)
	myDisplay.render()
}
