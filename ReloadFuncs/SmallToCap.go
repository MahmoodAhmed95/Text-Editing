package ReloadFuncs

import (
	"strings"
	"unicode"
)

func SmallToCap(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			runes := []rune(words[i-1])
			runes[0] = unicode.ToUpper(runes[0])
			words[i-1] = string(runes)
			words[i] = ""
		}
	}

	return strings.Join(words, " ")
}
