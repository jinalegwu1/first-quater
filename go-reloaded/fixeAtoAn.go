package main

import (
	"strings"
)

func FixAToAn(s string) string{
	w := strings.Fields(s)
	vowel := "aeiouhAEIOUH"
	for i := 0; i < len(w); i++ {
		if w[i] == "a" && strings.ContainsAny(w[i+1][:1], vowel) {
			w[i] = "an"
		} else if w[i] == "A" && strings.ContainsAny(w[i+i][:1], vowel) {
			w[i] = "An"
		} else if w[i] == "an" && !strings.ContainsAny(w[i+1][:1], vowel) {
			w[i] = "a"
		} else if w[i] == "An" && !strings.ContainsAny(w[i+1][:1], vowel) {
			w[i] = "A"
		}
	}
	return strings.Join(w, " ")
}