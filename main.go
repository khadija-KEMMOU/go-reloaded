package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
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

	//seperating words
	words := strings.Fields(sample)
	
	//looking for the hex to convert
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			hexNumber := words[i-1]
			result := hexToDecimal(hexNumber)
			words[i-1] = strconv.Itoa(result)
			fmt.Println(hexNumber)
		}
	}

	//writing the results
	err = os.WriteFile(os.Args[2], []byte(sample), 0644)
	if err != nil {
		return
	}
	////result := hexToDecimal("1E")

	fmt.Println(words)
}