package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
)

var ascii = []string{" ", ".", "-", ":", "=", "+", "*", "#", "░", "▒", "▓", "█"}

func main() {
	characterWidth, characterHeight, cores, input, output := parser()
	ch := make(chan int, cores)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	in, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding image: %v\n", err)
		os.Exit(1)
	}

	grayImg := image.NewGray(in.Bounds())
	for x := in.Bounds().Min.X; x < in.Bounds().Max.X; x++ {
		for y := in.Bounds().Min.Y; y < in.Bounds().Max.Y; y++ {
			grayImg.Set(x, y, in.At(x, y))
		}
	}

	outputArray := make([]string, (grayImg.Bounds().Max.Y-grayImg.Bounds().Min.Y)/characterHeight+1)
	for t := grayImg.Bounds().Min.Y; t < grayImg.Bounds().Max.Y; t += characterHeight {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
			for x := grayImg.Bounds().Min.X; x < grayImg.Bounds().Max.X; x += characterWidth {
				var sum, count int
				for n := t; n < t+characterHeight && n < grayImg.Bounds().Max.Y; n++ {
					for m := x; m < x+characterWidth; m++ {
						if m < grayImg.Bounds().Min.X || m >= grayImg.Bounds().Max.X || n < grayImg.Bounds().Min.Y || n >= grayImg.Bounds().Max.Y {
							continue
						}
						sum += int(grayImg.GrayAt(m, n).Y)
						count++
					}
				}
				index := sum / count * (len(ascii) - 1) / 255
				mu.Lock()
				outputArray[t/characterHeight] += ascii[index]
				mu.Unlock()
			}
			<-ch
		}()
	}

	wg.Wait()
	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	outputString := ""
	for _, line := range outputArray {
		outputString += line + "\n"
	}
	if _, err := outputFile.WriteString(outputString); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to output file: %v\n", err)
		os.Exit(1)
	}
}
