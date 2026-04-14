package decode_string

import "testing"

func TestDecodeString(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"111[bc]", "bcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbc"},
	}

	for _, c := range cases {
		res := decodeString(c.input)
		if res != c.output {
			t.Errorf("decodeString(%v)=%v but expected %v", c.input, res, c.output)
		}
	}
}
