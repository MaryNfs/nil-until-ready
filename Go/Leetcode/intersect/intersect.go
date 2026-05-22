package intersect

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, 0)
	final := make([]int, 0)
	for _, n := range nums1 {
		hash[n]++
	}
	for _, n := range nums2 {
		_, ok := hash[n]
		if ok && hash[n] > 0 {
			hash[n]--
			final = append(final, n)
		}
	}
	return final
}
