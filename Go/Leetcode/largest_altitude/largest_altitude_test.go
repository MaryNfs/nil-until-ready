package largest_altitude

import "testing"

func TestLargestAltitude(t *testing.T) {
	cases := []struct{
		gain []int
		out int
	}{
		{[]int{-5,1,5,0,-7},1},
		{[]int{-4,-3,-2,-1,4,3,2},0},
	}

	for _,c := range cases{
		res := largestAltitude(c.gain)
		if res != c.out {
			t.Errorf("largestAltitude(%v)=%v, but expected %v",c.gain,res,c.out)
		}
	}
}
