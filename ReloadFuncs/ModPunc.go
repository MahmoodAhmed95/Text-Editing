package ReloadFuncs

import "strings"

func ModPunc(match string) string {
	if strings.ContainsAny(match, "(){}&^'") {
		return match
	}
	PuncIsol := strings.Trim(match, " ")

	return PuncIsol + " "
}
