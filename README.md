# AsciiImage

AsciiImage is a simple command-line tool written in Go for converting images into ASCII art.

## Features

- Convert images (JPEG, PNG, etc.) to ASCII art
- Adjustable output width and character set
- Parallel executions

## Installation

Clone the repository and build the project:

```sh
git clone https://github.com/yourusername/AsciiImage.git
cd AsciiImage
go build .
```

## Usage
From the directory where the binary is located, you can run the tool with the following command:
```sh
./AsciiImage -i input.jpg
```
Available options are:
- `-i`: Input image file (required)
- `-o`: Output file (optional, defaults to `./out.txt)
- `-x`: Width of the character font (optional, defaults to 10)
- `-y`: Height of the character font (optional, defaults to 14)
- `-t`: Number of threads to use (optional, defaults to 16)

### Example:
```sh
./AsciiImage -i input.jpg -o output.txt -x 5 -y 7 -t 20 # Bigger image
```

```sh
./AsciiImage -i input.jpg -o output.txt -x 5 -y 7 -t 8 # Smaller image
```

## License
This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.