package decode_string

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var stack []rune
	for _, c := range s {
		if c != ']' {
			stack = append(stack, c)
		} else {
			substr := ""
			for stack[len(stack)-1] != '[' {
				substr = string(stack[len(stack)-1]) + substr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] //for removing [
			k := ""
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				k = string(stack[len(stack)-1]) + k
				stack = stack[:len(stack)-1]
			}
			val, _ := strconv.Atoi(k)
			repeated := strings.Repeat(substr, val)
			for _, r := range repeated {
				stack = append(stack, r)
			}

		}
	}
	return string(stack)
}
