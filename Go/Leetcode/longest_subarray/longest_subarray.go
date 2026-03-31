package longest_subarray

func longestSubarray(nums []int) int {
	count, res, zer := 0, 0, 0
	l := -1
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zer = 1
			count = r - l - 1
			l = r
		} else {
			count++
		}
		res = max(res, count)
	}
	if zer == 0 {
		res = len(nums) - 1
	}
	return res
}