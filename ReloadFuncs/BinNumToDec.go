package ReloadFuncs

import (
	"strconv"
	"strings"
)

func BinNumToDec(numbr string) string {
	words := strings.Fields(numbr)
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			binNumber := words[i-1]
			decimalNumber, err := strconv.ParseInt(binNumber, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimalNumber, 10)
			}
			words = append(words[:i], words[i+1:]...)
		}
	}
	out_str := strings.Join(words, " ")
	return out_str
}
