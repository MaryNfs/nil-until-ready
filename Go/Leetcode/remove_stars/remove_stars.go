package remove_stars

import (
	"slices"
	"strings"
)

// Timelimit solution
func removeStars(s string) string {
	arr := strings.Split(s, "")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "*" {
			arr = slices.Delete(arr, i-1, i+1)
			i = i - 2
		}
	}
	return strings.Join(arr, "")
}
