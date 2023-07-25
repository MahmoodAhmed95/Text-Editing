package ReloadFuncs

import (
	"strconv"
	"strings"
)

func HexNumToDec(input string) string {
	words1 := strings.Fields(input)
	for i := 1; i < len(words1); i++ {
		if words1[i] == "(hex)" {
			hexNumber := words1[i-1]
			decimalNumber, err := strconv.ParseInt(hexNumber, 16, 64)
			if err == nil {
				words1[i-1] = strconv.FormatInt(decimalNumber, 10)
			}
			words1 = append(words1[:i], words1[i+1:]...)
			i--
		}
	}
	out_str := strings.Join(words1, " ")
	return out_str
}
