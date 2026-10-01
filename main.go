package main

import (
	"os"
	"fmt"
)

func main() {
	//checking the args numbers
	if len(os.Args) != 3 {
		return
	}

	//storing the content to a var
	data, err := os.ReadFile(os.Args[1])
	//data is byte
	if err != nil {
		return
	}
	sample := string(data)

	result := processText(sample)

	//writing the results
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		return
	}

	fmt.Println("The answer have been written in the file result.txt successfully!")
}