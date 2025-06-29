package main

import (
	"flag"
	"fmt"
	"os"
)

func parser() (string, string, int, int, int) {
	var input = flag.String("i", "", "Input image file")
	var output = flag.String("o", "out.txt", "Output text file")
	var xFlag = flag.Int("x", 10, "Character font width")
	var yFlag = flag.Int("y", 14, "Character font height")
	var cores = flag.Int("t", 16, "Number of threads to use")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *xFlag <= 0 || *yFlag <= 0 {
		fmt.Fprintln(os.Stderr, "Font width and height be greater than 0")
		flag.Usage()
		os.Exit(1)
	}
	if *cores <= 0 {
		fmt.Fprintln(os.Stderr, "Number of threads must be greater than 0")
		flag.Usage()
		os.Exit(1)
	}
	return *input, *output, *xFlag, *yFlag, *cores
}
