package longest_subarray

import "testing"

func TestLongestSubarray(t *testing.T) {
	cases := []struct{
		nums []int
		out int
	}{
		{[]int{1,1,0,1},3},
		{[]int{0,1,1,1,0,1,1,0,1},5},
		{[]int{1,1,1},2},
		{[]int{0,1},1},
	}
	for _,c := range cases {
		res := longestSubarray(c.nums)
		if res != c.out {
			t.Errorf("longestSubarray(%v)=%v, but expected %v",c.nums,res,c.out)
		}
	}
}
