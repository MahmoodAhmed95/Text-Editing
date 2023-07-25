package ReloadFuncs

import (
	"strings"
)

func LowToUp(input string) string {
	words := strings.Fields(input)
	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
