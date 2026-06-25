package main

import (
	"strconv"
	"strings"
)

func Convert(con string)(string){
	word := strings.Fields(con)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0{
			hex, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				return err.Error()
			}
			word[i-1] = strconv.FormatInt(hex, 10)
			word = append(word[:i],word[i+1:]... )
		i--
		} else if word[i] == "(bin)" && i > 0{
			bin, err := strconv.ParseInt(word[i-1], 2, 64)
			if err != nil {
				return err.Error()
			}
			word[i-1] = strconv.FormatInt(bin, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(up)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i],word[i+1:]... )
			i--
		} else if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(cap)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1][:1]) + word[i-1][1:]
			word = append(word[:i], word[i+1:]...)
			i--
		} 
		
	}
	return strings.Join(word, " ")
	
}