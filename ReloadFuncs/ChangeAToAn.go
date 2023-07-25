package ReloadFuncs

import (
	"strings"
)

func ChangeAToAn(input string) string {
	text := strings.Fields(input)
	for i := 1; i < len(text); i++ {
		word := text[i]
		letter := word[0]
		if letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h' {
			if text[i-1] == "a" {
				text[i-1] = "an"
			} else if text[i-1] == "A" {
				text[i-1] = "An"
			}
		} else if letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H' {
			if text[i-1] == "a" {
				text[i-1] = "an"
			} else if text[i-1] == "A" {
				text[i-1] = "An"
			}
		}
	}

	out_str := strings.Join(text, " ")
	return out_str
}
