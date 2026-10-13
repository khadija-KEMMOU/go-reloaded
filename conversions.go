package main

import (
	"strconv"
	"strings"
)

// converts hexadecimals to decimals
func hexToDecimal(hex string) (int, error) {
	number, err := strconv.ParseInt(hex, 16, 64)
	return int(number), err
}

// converts binary to decimals
func binToDecimal(bin string) (int, error) {
	number, err := strconv.ParseInt(bin, 2, 64)
	return int(number), err
}

// transforms lowercase to uppercase
func lowToUp(low string) string {
	return strings.ToUpper(low)
}

// transforms uppercase to lowercase
func upToLow(up string) string {
	return strings.ToLower(up)
}

// capitalizes
func capital(cap string) string {
	if cap == "" {
		return cap
	}
	cap = strings.ToLower(cap)
	first := strings.ToUpper(cap[0:1])
	rest := cap[1:]
	result := first + rest
	return result
}

func isPunctuationChar(char rune) bool {
	return char == '!' || char == '.' || char == ',' || char == '?' || char == ';' || char == ':'
}

// tells if a word is a group of puncutations or has a letter in it
func isPunctuation(word string) bool {
	if word == "" {
		return false
	}
	for _, char := range word {
		if !isPunctuationChar(char) {
			return false
		}
	}
	return true
}

// tells if its a quote
func isQuote(words string, i int) string {
	if words == "'" {
		return "single"
	}
	if words[i] == '\'' {
		if i == 0 {
			return "opening"
		} else if i == len(words)-1 {
			return "closing"
		} else {
			return "none"
		}
	}
	return "none"
}