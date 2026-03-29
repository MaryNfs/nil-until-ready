package longest_ones

func longestOnes(nums []int, k int) int {
	l, count, res, flip := 0, 0, 0, 0
	for r := 0; r < len(nums); r++ {
        if nums[r]==0 {
            flip++
        }
        count ++

		if flip > k {
			if nums[l] == 0 {
				flip--
			}
			count--
			l++
		}
		res = max(res, count)
	}
	return res
}