package main

import (
	"fmt"
	"math"
	"strings"
)

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

var (
	// levels  []string = []string{" ", "░", "▒", "▓", "█"}
	levels []string = strings.Split(
		"$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft()1{}[]?-_+~<>i!lI;:,^`'. ",
		"",
	)
)

func (d *display) drawPoint(p *Point, alpha int) {
	if p.X >= d.width || p.X < 0 || p.Y >= d.height || p.Y < 0 {
		return
	}

	charIndex := alpha * len(levels) / 255
	intCharIndex := math.Min(float64(charIndex), float64(len(levels)-1))
	char := levels[int(intCharIndex)]
	d.buffer[p.Y][p.X] = rune(char[0])
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
