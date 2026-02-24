package find_difference

import (
	"slices"
	"testing"
)

func TestFindDifference(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		result [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, c := range cases {
		res := findDifference(c.nums1, c.nums2)
		equal := slices.EqualFunc(res, c.result, func(s1, s2 []int) bool {
			return slices.Equal(s1, s2)
		})
		if !equal{
			t.Errorf("findDifference(%v,%v)=%v but expected: %v",c.nums1,c.nums2,res,c.result)
		}
	}
}
