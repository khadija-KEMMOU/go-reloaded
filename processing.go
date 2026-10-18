package main

import (
	"strings"
)

// processes line by line and word by word
func processText(sample string) string {
	// seperating lines
	lines := strings.Split(sample, "\n")

	for i := 0; i < len(lines); i++ {
		// seperating words
		words := strings.Fields(lines[i])
		words = splitCommandPunctuation(words)

		for j := 0; j < len(words); j++ {
			// the conversion part
			switch words[j] {
			case "(hex)":
				words, j = processConversion(words, j, hexToDecimal)
			case "(bin)":
				words, j = processConversion(words, j, binToDecimal)
			case "(up)":
				words, j = processTransformation(words, j, lowToUp)
			case "(up,":
				words, j = processTransformWithNbr(words, j, lowToUp)
			case "(low)":
				words, j = processTransformation(words, j, upToLow)
			case "(low,":
				words, j = processTransformWithNbr(words, j, upToLow)
			case "(cap)":
				words, j = processTransformation(words, j, capital)
			case "(cap,":
				words, j = processTransformWithNbr(words, j, capital)
			}
		}
		words = processQuotes(words)

		words = processPunctuation(words)

		words = processArticle(words)
		lines[i] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}
