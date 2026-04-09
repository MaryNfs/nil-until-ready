package is_valid

func isValid(s string) bool {
	var stack []byte

	open := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	close := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		_, ok1 := open[s[i]]
		if ok1 {
			stack = append(stack, s[i])
		}
		c, ok2 := close[s[i]]
		if ok2 {
			if len(stack) > 0 && c == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}