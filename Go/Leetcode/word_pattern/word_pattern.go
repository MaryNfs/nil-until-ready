package word_pattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sarr := strings.Split(s, " ")
	parr := strings.Split(pattern, "")
	if len(parr) != len(sarr) {
		return false
	}

	fmap := make(map[string]string, 0)
	for i := 0; i < len(sarr); i++ {
		value, ok := fmap[sarr[i]]
		if !ok {
			fmap[sarr[i]] = parr[i]
		} else if value != parr[i] {
			return false
		}
	}
	kmap := make(map[string]string, 0)
	for i := 0; i < len(sarr); i++ {
		value, ok := kmap[parr[i]]
		if !ok {
			kmap[parr[i]] = sarr[i]
		} else if value != sarr[i] {
			return false
		}
	}

	return true
}
