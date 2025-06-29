package main

import (
	"flag"
)

func parser() (string, string, int, int, int) {
	var input = flag.String("i", "", "Input image file")
	var output = flag.String("o", "out.txt", "Output text file")
	var xFlag = flag.Int("x", 10, "Character font width")
	var yFlag = flag.Int("y", 14, "Character font height")
	var threads = flag.Int("t", 16, "Number of threads to use")

	flag.Parse()
	return *input, *output, *xFlag, *yFlag, *threads
}
