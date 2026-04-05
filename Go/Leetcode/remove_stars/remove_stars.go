package remove_stars

import (
	"slices"
	"strings"
)

// Timelimit solution
// func removeStars(s string) string {
// 	arr := strings.Split(s, "")
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == "*" {
// 			arr = slices.Delete(arr, i-1, i+1)
// 			i = i - 2
// 		}
// 	}
// 	return strings.Join(arr, "")
// }

// accepted solution
// define stack array with type byte
// no need to convert slice to string and vice versa
func removeStars(s string) string {
	stack := make([]byte,0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		}else{
            stack = append(stack,s[i])
        }
	}
	return string(stack)
}