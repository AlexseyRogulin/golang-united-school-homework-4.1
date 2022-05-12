package reverse_string

import "strings"

func ReverseString(input string) string {
	var sb strings.Builder
	for i := len([]rune(input)) - 1; i >= 0; i-- {
		sb.WriteRune([]rune(input)[i])
	}
	return sb.String()
}
