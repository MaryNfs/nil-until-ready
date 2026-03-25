package maxOperations

import "testing"

func TesMaxOperations(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		out  int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 5, 6},
		{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 4, 2},
	}

	for _,c := range cases{
		res := maxOperations(c.nums,c.k)
		if res != c.out{
			t.Errorf("maxOperations(%v,%v)=%v, but expected %v",c.nums,c.k,res,c.out)
		}
	}
}
