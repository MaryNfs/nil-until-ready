package create_target_array

func createTargetArray(nums []int, index []int) []int {
	length := len(nums)
	t := make([]int, 0, length)
	for i := 0; i < length; i++ {
		t = append(t, 0)
	}
	for i := 0; i < length; i++ {
		// 2. Shift elements to the right to make room
		copy(t[index[i]+1:], t[index[i]:])

		// 3. Insert the new value at the index
		t[index[i]] = nums[i]
	}
	return t
}