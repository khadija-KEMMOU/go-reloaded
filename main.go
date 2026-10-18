package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// checking the args numbers
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input.txt> <output.txt>")
		return
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// checking the paths before touching any file
	err := checkPaths(inputPath, outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// storing the content to a var
	sample, err := os.ReadFile(os.Args[1])
	// data is byte
	if err != nil {
		fmt.Println("Error: cannot read the input file")
		return
	}

	result := processText(string(sample))

	// writing the results
	err = os.WriteFile(outputPath, []byte(result), 0o644)
	if err != nil {
		fmt.Println("Error: cannot write the output file")
		return
	}

	fmt.Println("The answer has been written in", outputPath, "successfully!")
}

// makes sure the output can't erase a file that isn't a result file
func checkPaths(inputPath, outputPath string) error {
	if !strings.HasSuffix(strings.ToLower(outputPath), ".txt") {
		return errors.New("the output file must end with .txt")
	}

	absInput, _ := filepath.Abs(inputPath)
	absOutput, _ := filepath.Abs(outputPath)
	if absInput == absOutput {
		return errors.New("the input and output must be different files")
	}
	return nil
}
