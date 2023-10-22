package main

import "flag"

type Options struct {
	mode            *string
	sourceImagePath *string
}

func getOptions() *Options {
	mode := flag.String("mode", "static", "Mode of program. One of [sim, image, camera, static]")
	sourceImagePath := flag.String("source", "assets/cat.jpg", "Source of image")

	flag.Parse()

	return &Options{
		mode,
		sourceImagePath,
	}
}
