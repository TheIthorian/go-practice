package main

func main() {
	options := getOptions()

	if *options.mode == "sim" {
		bouncingPoint()
		return
	}

	if *options.mode == "image" {
		renderImage(*options.sourceImagePath)
		return
	}

	alternatingPoints()
}
