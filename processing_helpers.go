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
	prefix := "" // opening quote(s) waiting to be glued to the next word

	// handles ONE quote mark
	mark := func(forceClose bool) {
		if !inQuote && !forceClose {
			// opening quote: wait for the next word
			prefix += "'"
			inQuote = true
			return
		}
		// closing quote
		inQuote = false
		if prefix != "" {
			// nothing between the quotes: keep them together as ''
			parts = append(parts, prefix+"'")
			prefix = ""
		} else if len(parts) > 0 {
			// glue to the end of the previous word
			parts[len(parts)-1] += "'"
		} else {
			parts = append(parts, "'")
		}
	}

	for _, word := range words {
		core := strings.Trim(word, "'")

		// the word is only quotes, like "'"
		if core == "" {
			for range word {
				mark(false)
			}
			continue
		}

		lead := len(word) - len(strings.TrimLeft(word, "'"))
		trail := len(word) - len(strings.TrimRight(word, "'"))

		// quotes at the start of the word
		for k := 0; k < lead; k++ {
			// a quote followed only by punctuation ('.) must be a closing one
			mark(!inQuote && isPunctuation(core) && len(parts) > 0)
		}

		parts = append(parts, prefix+core)
		prefix = ""

		// quotes at the end of the word
		for k := 0; k < trail; k++ {
			mark(false)
		}
	}

	// an opening quote with nothing after it
	if prefix != "" {
		parts = append(parts, prefix)
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