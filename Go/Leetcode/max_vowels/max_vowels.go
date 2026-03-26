package max_vowels

import "strings"

// TimeLimit
func maxVowels(s string, k int) int {
	arr := strings.Split(s, "")
	vows := make(map[string]int, 0)
	vows = map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}
	count := 0
	max := 0
	j := 0
	for j <= len(arr)-k {
		for i := j; i < j+k; i++ {
			_, ok := vows[arr[i]]
			if ok {
				count++
			}
		}
		if count > max {
			max = count
		}
		count = 0
		j++
	}
	return max
}
