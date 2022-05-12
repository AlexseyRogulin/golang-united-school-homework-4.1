package reverse_string

import "strings"

func ReverseString(input string) string {
	var sb strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		sb.WriteByte(input[i])
	}
	return sb.String()
}
