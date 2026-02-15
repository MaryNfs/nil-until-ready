package happy_number

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct{
		input int
		expected bool
	}{
		{1,true},
		{19, true},
		{2,false},
		{7,true},
		{32,true},
	}
	for _,test := range tests{
		res := isHappy(test.input)
		if res != test.expected {
			t.Errorf("isPalindrome(%v) = %v; expected %v", test.input, res, test.expected)
		}
	}
}
