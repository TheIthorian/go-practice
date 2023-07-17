package main

import (
	"io"
	"os"
	"strings"
)

var (
	alphabet        string = "abcdefghijklmnopqrstuvwxyz"
	captialAlphabet string = strings.ToUpper(alphabet)
)

func rot(char byte) byte {
	charArray := &alphabet
	if char >= 'A' && char <= 'Z' {
		charArray = &captialAlphabet
	}

	var index int = -1
	for i, letter := range *charArray {
		if rune(char) == letter {
			index = i
		}
	}

	if index == -1 {
		return byte(char)
	}

	rotIndex := (index + 13) % len(*charArray)
	return (*charArray)[rotIndex]
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(buffer []byte) (int, error) {
	n, err := r.r.Read(buffer)
	if err != nil {
		return n, err
	}

	for i, val := range buffer {
		buffer[i] = rot(val)
	}

	return len(buffer), nil
}

func main() {
	inputString := ""
	if len(os.Args) > 1 {
		inputString = os.Args[1]
	}

	s := strings.NewReader(inputString + "\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
