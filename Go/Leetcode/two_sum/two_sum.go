package two_sum


func twoSum(nums []int, target int) []int {
	res := make([]int, 0)

	// O(n^2)
	// for i := 0; i < len(nums); i++ {
	// 	x := target - nums[i]
	// 	for j := i+1; j < len(nums); j++ {
	// 		if x == nums[j] {
	// 			res = append(res, i, j)
	// 			return res
	// 		}
	// 	}
	// }

	// O(n)
	mapping := make(map[int]int,0)
	var x int
	for i:=0;i<len(nums);i++{
		mapping[nums[i]] = i
	}
	for i:=0;i<len(nums);i++{
		x = target - nums[i]
		if x>=0 {
			key, ok := mapping[x]
			if ok && (x != nums[i] || key !=i) {
				return []int{i,key}
			}
		}
	}

	return res
}

