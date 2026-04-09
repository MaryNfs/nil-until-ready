package is_valid

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct{
		s string
		out bool
	}{
		{"()",true},
		{"()[]{}",true},
		{"(]",false},
		{"([)]",false},
		{"]",false},
	}
	for _,c := range cases{
		res := isValid(c.s)
		if res != c.out{
			t.Errorf("isValid(%v)=%v, but expected %v",c.s,res,c.out)
		}
	}
}
