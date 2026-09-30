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

	//seperating lines
	lines := strings.Split(sample, "\n")

	for i := 0; i < len(lines); i++ {
		//seperating words
		words := strings.Fields(lines[i])
		converted := make([]bool, len(words))

		//looking for the hex to convert
		for j := 0; j < len(words); j++ {

			//looking specifically for hex to convert
			if words[j] == "(hex)" {
				if j > 0 && !converted[j-1] {
					hexNumber := words[j-1]
					result, err := hexToDecimal(hexNumber)
					
					if err == nil {
						words[j-1] = strconv.Itoa(result)
						converted[j-1] = true
					}
				}
				words = append(words[:j], words[j+1:]...)
				converted = append(converted[:j], converted[j+1:]...)
				j--
			} else if words[j] == "(bin)" {   //looking specificaly for bin to convert
				if j > 0 && !converted[j-1] {
					binNumber := words[j-1]
					result, err := binToDecimal(binNumber)

					if err == nil {
						words[j-1] = strconv.Itoa(result)
						converted[j-1] = true
					}
				}
				words = append(words[:j], words[j+1:]...)
				converted = append(converted[:j], converted[j+1:]...)
				j--
			}
			
		}
		lines[i] = strings.Join(words, " ")
	}

	result := strings.Join(lines, "\n")

	//writing the results
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		return
	}

	fmt.Println("The answer have been written in the file result.txt successfully!")
}