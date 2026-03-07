package reverse_vowels

import (
	"strings"
)

func reverseVowels(s string) string {
	v := map[string]int{
		"a": 1,
		"A": 1,
		"e": 1,
		"E": 1,
		"i": 1,
		"I": 1,
		"o": 1,
		"O": 1,
		"u": 1,
		"U": 1,
	}
	sAr := strings.Split(s, "")
	j := 0
	rv := make([]string, 0)
	for i := 0; i < len(sAr); i++ {
		_, ok := v[sAr[i]]
		if ok {
			rv = append(rv, sAr[i])
			j++
		}
	}
	j = len(rv) - 1
	for i := 0; i < len(sAr); i++ {
		_, ok := v[sAr[i]]
		if ok {
			sAr[i] = rv[j]
			j--
		}
	}
	res := strings.Join(sAr, "")
	return res
}
