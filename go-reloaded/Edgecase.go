package main

import (
	"regexp"
	"strings"
)

func EdgeCases(s string) string{
	word := strings.Join(strings.Fields(s), " ")
	// fixed Quote
	word = regexp.MustCompile(`'\s+([^']*?)\s+'`).ReplaceAllString(word, `'$'`)
	word = regexp.MustCompile(`"\s+([^"]*?)\s+"`).ReplaceAllString(word, `"$"`)
	// fixed puntuations
	word = regexp.MustCompile(`\s+([:;.,!?]+)`).ReplaceAllString(word, `$1`)
	word = regexp.MustCompile(`([:;.,!?])([a-zA-Z0-9])`).ReplaceAllString(word, `$1 $2`)
	return word
}