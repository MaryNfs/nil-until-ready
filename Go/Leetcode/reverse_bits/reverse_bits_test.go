package reverse_bits

import "testing"

func TestReverseBits(t *testing.T) {
	cases := []struct {
		num int
		out int
	}{
		{43261596, 964176192},
		{2147483644, 1073741822},
	}
	for _, c := range cases {
		res := reverseBits(c.num)
		if res != c.out {
			t.Errorf("reverseBits(%v)=%v, but expected %v", c.num, res, c.out)
		}
	}
}
