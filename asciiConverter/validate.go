package asciiConverter

import "fmt"

func validate(input string, characterWidth int, characterHeight int, threads int) error {
	if input == "" {
		return fmt.Errorf("input file cannot be empty")
	}
	if characterWidth <= 0 || characterHeight <= 0 {
		return fmt.Errorf("character width and height must be greater than 0")
	}
	if threads <= 0 {
		return fmt.Errorf("number of threads must be greater than 0")
	}
	return nil
}
