package reverse_words

import (
	"strings"
)

func reverseWords(s string) string {
	res := strings.Split(strings.TrimSpace(s), " ")
	f := ""
	for i := len(res) - 1; i >= 0; i-- {
		if strings.TrimSpace(res[i]) != "" {
			if i == 0 {
				f += res[i]
			} else {
				f += res[i] + " "
			}
		}
	}
	return f
}
