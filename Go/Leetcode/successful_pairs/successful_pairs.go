package successful_pairs

import "slices"

// O(n*m)
// func successfulPairs(spells []int, potions []int, success int64) []int {
// 	res := make([]int, 0)
// 	for _, s := range spells {
// 		r := 0
// 		for _, p := range potions {
// 			num := s * p
// 			if int64(num) >= success {
// 				r++
// 			}
// 		}
// 		res = append(res, r)
// 	}
// 	return res
// }

// O(nlogm+mlogm) binary search
func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, 0)
	slices.Sort(potions)
	for _, s := range spells {
		index := len(potions)
		l, r := 0, len(potions)-1

		for l <= r {
			m := (r + l) / 2
			if int64(s)*int64(potions[m]) >= success {
				r = m - 1
				index = m
			} else {
				l = m + 1
			}

		}

		res = append(res, len(potions)-index)
	}
	return res
}
