package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aa", true},
		{" ", true},
		{"race a car", false},
		{"A man, a plan, a canal: Panama", true},
	}
	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}

}
