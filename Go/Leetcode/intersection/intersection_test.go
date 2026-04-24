package intersection

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		out   []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, c := range cases {
		res := intersection(c.nums1, c.nums2)
		if !slices.Equal(res, c.out) {
			t.Errorf("intersection(%v,%v)=%v but expected %v", c.nums1, c.nums2, res, c.out)
		}
	}
}
