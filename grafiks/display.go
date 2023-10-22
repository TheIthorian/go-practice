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

	return display{buffer, width, height}
}

func (d *display) drawPoint(p *Point, alpha int) {
	if p.X >= d.width || p.X < 0 || p.Y >= d.height || p.Y < 0 {
		return
	}

	char := ' '
	if alpha > 200 {
		char = '#'
	}

	d.buffer[p.Y][p.X] = char
}

func (d *display) clear() {
	fmt.Print("\033[H\033[2J")
	for Y := 0; Y < d.height; Y++ {
		for X := 0; X < d.width; X++ {
			d.drawPoint(&Point{X, Y}, 0)
		}
	}
}

func (d *display) render() {
	for _, row := range d.buffer {
		fmt.Println(toString(row))
	}
}
