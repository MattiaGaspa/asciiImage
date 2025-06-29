# AsciiImage

AsciiImage is a simple command-line tool written in Go for converting images into ASCII art.

## Features

- Convert images (JPEG, PNG, etc.) to ASCII art
- Adjustable output width and character set
- Parallel executions

## Usage

### CLI Usage
Clone the repository and build the project:

```sh
git clone https://github.com/yourusername/asciiImage.git
cd asciiImage
go build .
```

From the directory where the binary is located, you can run the tool with the following command:
```sh
./asciiImage -i input.jpg
```
Available options are:
- `-i`: Input image file (required)
- `-o`: Output file (optional, defaults to `./out.txt)
- `-x`: Width of the character font (optional, defaults to 10)
- `-y`: Height of the character font (optional, defaults to 14)
- `-t`: Number of threads to use (optional, defaults to 16)

#### Examples:
```sh
./asciiImage -i input.jpg -o output.txt -x 5 -y 7 -t 20 # Bigger image, 20 threads
```
```sh
./asciiImage -i input.jpg -o output.txt -x 5 -y 7 -t 8 # Smaller image, 8 threads
```

### Go Package Usage
You can also use asciiImage as a Go package in your own projects. Import the package and use the `ConvertToASCII` function:

``` go
package main

import (
    "os"
    "fmt"
    "github.com/MattiaGaspa/asciiImage/asciiConverter"
)

func main() {
    inputFile := "input.jpg"
    characterWidth := 10
    characterHeight := 14
    threads := 16

    outputString, err := asciiConverter.ConvertToASCII(inputFile, characterWidth, characterHeight, threads)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error converting image to ASCII: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println(outputString)
}
```

## License
This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.