package main

import (
	"strings"
	"unicode"
)

func GetWords(str string) []string {
	splitWords := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	}
	return strings.FieldsFunc(str, splitWords)
}

func WordCount(str string) int {
	return len(GetWords(str))
}

func GetSentences(str string) []string {
	// could implement quoting
	splitSent := func(c rune) bool {
		return c == '.' || c == '!' || c == '?' || c == '\n'
	}
	return strings.FieldsFunc(str, splitSent)

}
