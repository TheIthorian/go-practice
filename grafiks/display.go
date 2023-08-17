package main

import "fmt"

type display struct {
	buffer [][]rune
	width  int
	height int
}

func makeDisplay(width, height int) display {
	var buffer [][]rune
	for i := 0; i < height; i++ {
		buffer = append(buffer, make([]rune, width))
	}

	println(len(buffer))
	println(len(buffer[0]))
	return display{buffer, width, height}
}

func (d *display) drawPoint(p *point, a int) {
	d.buffer[p.Y][p.X] = '#'
}

func (d *display) render() {
	for _, row := range d.buffer {
		fmt.Println(toString(row))
	}
}
