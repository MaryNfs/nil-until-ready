package valid_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	arr := strings.Split(s, "")
	if len(arr) < 1 {
		return true
	}
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
