package utils

import (
	"strings"
	"unicode"
)

// tokenize returns a slice of tokens for the given text.
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// analyze analyzes the text and returns a slice of tokens.
func analyze(text string) []string {
	tokens := tokenize(text)         //converts ["he is a good boy"] to ["he", "is", "a", "good", "boy"]
	tokens = lowercaseFilter(tokens) //converts to lower case
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
