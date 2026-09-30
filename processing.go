package main

import (
	"strings"
	"strconv"
)

func processConversion(words []string, converted []bool, j int, convert func(string) (int, error)) {

}

func processText(sample string) string {
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
	return strings.Join(lines, "\n")
}