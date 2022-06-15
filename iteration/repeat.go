package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {

		repeated = repeated + character
	}

	return repeated
}

func Join(list []string) string {

	return strings.Join(list, "")
}
