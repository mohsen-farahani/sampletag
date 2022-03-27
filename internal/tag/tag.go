package tag

import "strings"

func Convertor(str string) []string {
	if strings.Contains(str, "|") {
		return strings.Split(str, "|")
	}

	return strings.Split(str, ",")
}
