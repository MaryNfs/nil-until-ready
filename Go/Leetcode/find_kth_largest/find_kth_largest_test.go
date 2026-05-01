package find_kth_largest

import "testing"

func TestFindKthLargest(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		out  int
	}{
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
	}

	for _, c := range cases {
		res := findKthLargest(c.nums,c.k)
		if res != c.out {
			t.Errorf("findKthLargest(%v,%v)=%v but expected %v",c.nums,c.k,res,c.out)
		}
	}
}
