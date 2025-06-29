package main

import (
	"fmt"
	"github.com/MattiaGaspa/asciiImage/asciiConverter"
	"os"
)

func main() {
	input, output, characterWidth, characterHeight, cores := parser()

	outputString, err := asciiConverter.ConvertToASCII(input, characterWidth, characterHeight, cores)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting image to ASCII: %v\n", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	if _, err := outputFile.WriteString(outputString); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to output file: %v\n", err)
		os.Exit(1)
	}
}
