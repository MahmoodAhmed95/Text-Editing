package ReloadFuncs

import "strings"

func ModQuote(match string) string {
	inside_quote := match[1 : len(match)-1]
	inside_quote = strings.Trim(inside_quote, " ")

	return "'" + inside_quote + "'"
}
