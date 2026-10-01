package main

import (
	"strings"
	"strconv"
)

func processConversion(words []string, converted []bool, j int, convert func(string) (int, error)) ([]string, []bool, int) {
	if j > 0 && !converted[j-1] {
		number := words[j-1]
		result, err := convert(number)

		if err == nil {
			words[j-1] = strconv.Itoa(result)
			converted[j-1] = true
		}
	}
	words = append(words[:j], words[j+1:]...)
	converted = append(converted[:j], converted[j+1:]...)
	j--
	return words, converted, j
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

			//the conversion part
			switch words[j] {
			case "(hex)" :
				words, converted, j = processConversion(words, converted, j, hexToDecimal)
			case "(bin)" :
				words, converted, j = processConversion(words, converted, j, binToDecimal)
			}
			
		}
		lines[i] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}