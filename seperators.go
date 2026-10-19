package main

import (
	"strconv"
	"strings"
)

// separates the punctuation stuck after a ")": "(up)," becomes "(up)" and ","
func splitCommandPunctuation(words []string) []string {
	result := []string{}
	for _, word := range words {
		end := strings.LastIndex(word, ")")
		if end != -1 && end < len(word)-1 && isPunctuation(word[end+1:]) {
			result = append(result, word[:end+1], word[end+1:])
		} else {
			result = append(result, word)
		}
	}
	return result
}

// separates the quotes stuck to the start or end of a word: "'hi'" becomes "'", "hi", "'"
func splitQuotes(words []string) []string {
	result := []string{}
	for _, word := range words {
		core := strings.Trim(word, "'")

		// the word is only quotes, or has none: keep each quote alone
		if core == "" {
			for range word {
				result = append(result, "'")
			}
			continue
		}

		lead := len(word) - len(strings.TrimLeft(word, "'"))
		trail := len(word) - len(strings.TrimRight(word, "'"))

		for k := 0; k < lead; k++ {
			result = append(result, "'")
		}
		result = append(result, core)
		for k := 0; k < trail; k++ {
			result = append(result, "'")
		}
	}
	return result
}

// every command the switch in processText knows
var commands = []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)", "(up,", "(low,", "(cap,"}

// cuts one word around the first command found, then repeats on both sides
func cutCommands(word string) []string {
	for _, command := range commands {
		pos := strings.Index(word, command)
		if pos == -1 {
			continue
		}
		parts := []string{}
		if pos > 0 {
			parts = append(parts, cutCommands(word[:pos])...)
		}
		parts = append(parts, command)
		rest := word[pos+len(command):]
		if rest != "" {
			parts = append(parts, cutCommands(rest)...)
		}
		return parts
	}
	return []string{word}
}

// cuts the commands out of every word: "hello(up)k" becomes "hello", "(up)", "k"
func splitCommands(words []string) []string {
	result := []string{}
	for _, word := range words {
		result = append(result, cutCommands(word)...)
	}
	return result
}

// cuts what is stuck after the number's ")": "2)k" becomes "2)" and "k"
func splitAfterNumber(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		result = append(result, words[i])

		isCommand := words[i] == "(up," || words[i] == "(low," || words[i] == "(cap,"
		if !isCommand || i+1 >= len(words) {
			continue
		}

		next := words[i+1]
		end := strings.Index(next, ")")
		if end == -1 || end == len(next)-1 {
			continue // no ")", or nothing stuck after it
		}
		if _, err := strconv.Atoi(next[:end]); err != nil {
			continue // what comes before ")" is not a number
		}

		result = append(result, next[:end+1], next[end+1:])
		i++ // the next word is already added
	}
	return result
}
