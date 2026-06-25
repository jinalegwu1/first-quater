package main

import (
	"strconv"
	"strings"
)

func CommandNWord(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if i+1 < len(word) && (word[i] == "(up," || word[i] == "(cap," || word[i] == "(low,") {
			numsr := strings.TrimSuffix(word[i+1], ")")
			n, _ := strconv.Atoi(numsr)
			start := i - n
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				if word[i] == "(up," {
					word[j] = strings.ToUpper(word[j])
				} else if word[i] == "(low," {
					word[j] = strings.ToLower(word[j])
				} else if word[i] == "(cap," {
					if len(word[j]) > 0{
					word[j] = strings.ToUpper(word[j][:1]) + word[j][1:]
					}
				}
			}
			word = append(word[:i], word[i+2:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}
