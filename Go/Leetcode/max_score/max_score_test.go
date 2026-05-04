package max_score

import "testing"

func TestMaxScore(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		k     int
		out   int64
	}{
		{[]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3, 12},
		{[]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1, 30},
	}

	for _, c := range cases {
		res := maxScore(c.nums1, c.nums2, c.k)

		if res != c.out {
			t.Errorf("maxScore(%v,%v,%v)=%v but expected %v", c.nums1, c.nums2, c.k, res, c.out)
		}
	}
}
