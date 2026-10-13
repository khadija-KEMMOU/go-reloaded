package main

import (
	"strconv"
	"strings"
)

// handles only hex and bin functions
func processConversion(words []string, j int, convert func(string) (int, error)) ([]string, int) {
	if j > 0 {
		number := words[j-1]
		result, err := convert(number)

		if err == nil {
			words[j-1] = strconv.Itoa(result)
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
		for i := j-1; i >= 0 && n > 0; i-- {
			words[i] = transform(words[i])
			n--
		}

		words = append(words[:j], words[j+2:]...)
		j -= 2
	}
	return words, j
}

//handles punctuation except the quotes ""
func processPunctuation(words []string) []string {
	parts := []string{}
	for i := 0; i < len(words); i++ {
		current := ""
		punctuation := ""
		inPunctuation := false
		for _, char := range words[i] {
			if isPunctuationChar(char) {
				if !inPunctuation {
					parts = append(parts, current)
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
	
	for i := 0; i < len(words); i++ {
		if i > 0 && isPunctuation(words[i]) {
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}

// handles the quotes
func processQuotes(words []string) []string {
	parts := []string{}
	inQuote := false
	for i := 0; i < len(words); i++ {
		handled := false
		for j := 0; j < len(words[i]); j++ {
			if isQuote(words[i], j) {
				// handles the single quotes
				if words[i] == "'" {
					if inQuote == false {
						inQuote = true
						if i+1 < len(words) {
							parts = append(parts, "'" + words[i+1])
							i++
							handled = true
						}
					} else {
						inQuote = false
						parts[len(parts)-1] += "'"
					}
				// handles quotes attached to words
				} else {
					if j == len(words[i])-1 && !inQuote {
						parts = append(parts, words[i][:j])
						handled = true
						inQuote = true
						if i+1 < len(words) {
							parts = append(parts, "'"+words[i+1])
							i++
							break
						}
					}
					if j == 0 && inQuote {
						parts[len(parts)-1] +="'"
						parts = append(parts, words[i][1:])
						handled = true
						inQuote = false
						break
					}
				}
			}
		}
		if !handled && words[i] != "'" {
			parts = append(parts, words[i])
		}
	}
	return parts
}