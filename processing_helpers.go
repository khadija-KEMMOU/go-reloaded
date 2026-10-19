package main

import (
	"strconv"
	"strings"
)

// handles only hex and bin functions
func processConversion(words []string, j int, convert func(string) (string, error)) ([]string, int) {
	if j > 0 {
		number := words[j-1]
		result, err := convert(number)

		if err == nil {
			words[j-1] = result
			// to eliminate the command between (...)
			words = append(words[:j], words[j+1:]...)
			j--
		}
	}
	return words, j
}

// handles only up, low and cap functions without numbers in it
func processTransformation(words []string, j int, transform func(string) string) ([]string, int) {
	if j > 0 {
		targetWord := words[j-1]
		result := transform(targetWord)
		words[j-1] = result
		words = append(words[:j], words[j+1:]...)
		j--
	}
	return words, j
}

// handles only up, low and cap functions with numbers in it
func processTransformWithNbr(words []string, j int, transform func(string) string) ([]string, int) {
	if j == 0 || j+1 >= len(words) {
		return words, j
	}

	number := strings.TrimSuffix(words[j+1], ")")
	n, err := strconv.Atoi(number)
	if err == nil {
		for i := j - 1; i >= 0 && n > 0; i-- {
			words[i] = transform(words[i])
			n--
		}

		words = append(words[:j], words[j+2:]...)
		j -= 2
	}
	return words, j
}

// handles punctuation except the quotes ""
func processPunctuation(words []string) []string {
	parts := []string{}
	// this loop seperate the words and punctuation each alone
	for i := 0; i < len(words); i++ {
		current := ""
		punctuation := ""
		inPunctuation := false
		for _, char := range words[i] {
			if isPunctuationChar(char) {
				if !inPunctuation {
					if current != "" {
						parts = append(parts, current)
					}
					current = ""
					inPunctuation = true
				}
				punctuation += string(char)
			} else {
				if inPunctuation {
					parts = append(parts, punctuation)
					punctuation = ""
					inPunctuation = false
				}
				current += string(char)
			}
		}
		if punctuation != "" {
			parts = append(parts, punctuation)
		}
		if current != "" {
			parts = append(parts, current)
		}
	}

	words = parts

	// this loop glues the punctuations to the word before it
	for i := 0; i < len(words); i++ {
		if i > 0 && isPunctuation(words[i]) {
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}

// glues each lone quote: the 1st sticks to the next word, the 2nd to the previous one
func processQuotes(words []string) []string {
	words = splitQuotes(words)
	parts := []string{}
	inQuote := false

	for i := 0; i < len(words); i++ {
		if words[i] != "'" {
			parts = append(parts, words[i])
			continue
		}

		if !inQuote {
			// opening quote: glue it to the next word
			inQuote = true
			if i+1 >= len(words) {
				parts = append(parts, "'")
			} else if words[i+1] == "'" {
				// closes right away: empty quotes
				parts = append(parts, "''")
				inQuote = false
				i++
			} else {
				parts = append(parts, "'"+words[i+1])
				i++
			}
		} else {
			// closing quote: glue it to the previous word
			inQuote = false
			if len(parts) > 0 {
				parts[len(parts)-1] += "'"
			} else {
				parts = append(parts, "'")
			}
		}
	}
	return parts
}

// processes the a and an article
func processArticle(words []string) []string {
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			if startWithVowel(words[i+1]) {
				switch words[i] {
				case "a":
					words[i] = "an"
				case "A":
					words[i] = "An"
				}
			}
		}
	}
	return words
}
