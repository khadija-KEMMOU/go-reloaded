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
	letters := []rune(cap)
	first := strings.ToUpper(string(letters[0:1]))
	rest := letters[1:]
	result := string(first) + string(rest)
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

// specifies if a word starts with a vowel
func startWithVowel(word string) bool {
	if word == "" {
		return false
	}
	if word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' || word[0] == 'h' || word[0] == 'A' || word[0] == 'E' || word[0] == 'I' || word[0] == 'O' || word[0] == 'U' || word[0] == 'H' {
		return true
	}
	return false
}