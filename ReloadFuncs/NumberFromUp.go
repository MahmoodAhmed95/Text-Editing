package ReloadFuncs

import (
	"strconv"
	"strings"
)

func NumberFromUp(string_list []string, idx int) {
	num_str := strings.Split(string_list[idx+1], ")")
	num, _ := strconv.Atoi(num_str[0])

	for i := idx; i > (idx-num) && i > 0; i-- {
		string_list[i-1] = strings.ToUpper(string_list[i-1])
	}

	string_list[idx] = ""
	string_list[idx+1] = ""
}
