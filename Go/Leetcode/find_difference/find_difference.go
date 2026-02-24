package find_difference

// Slowest possible way? :)))
func findDifference(nums1 []int, nums2 []int) [][]int {
	res1 := make([]int, 0)
	res2 := make([]int, 0)

	map1 := map[int]int{}
	map2 := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				map1[nums1[i]]++
			}
		}
	}

	for i := 0; i < len(nums1); i++ {
		_, ok := map1[nums1[i]]
		if !ok {
			map1[nums1[i]]++
			res1 = append(res1, nums1[i])
		}
	}

	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums2[i] == nums1[j] {
				map2[nums2[i]]++
			}
		}
	}

	for i := 0; i < len(nums2); i++ {
		_, ok := map2[nums2[i]]
		if !ok {
			map2[nums2[i]]++
			res2 = append(res2, nums2[i])
		}
	}

	return [][]int{res1, res2}
}
