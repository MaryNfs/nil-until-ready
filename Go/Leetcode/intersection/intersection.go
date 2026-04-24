package intersection

func intersection(nums1 []int, nums2 []int) []int {
	res := make(map[int]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				_, ok := res[nums1[i]]
				if !ok {
					res[nums1[i]]++
				}
			}
		}
	}
    final := make([]int,0)
    for k,_ := range res{
        final= append(final,k)
    }
	return final
}
