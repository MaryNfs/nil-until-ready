package is_subsequence

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	sar := strings.Split(s, "")
	tar := strings.Split(t, "")
	j := 0
	if len(sar) == 0 {
		return true
	}
	for i := 0; i < len(tar); i++ {
		if j < len(sar) {
			if sar[j] == tar[i] {
				j++
			}
		}
	}
	if j == len(sar) {
		return true
	}
	return false
}

