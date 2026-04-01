package reverse_bits

// Not cool Question
func reverseBits(n int) int {
	res := 0
	for _ = range 32 {
        // left shift res by 1
		res = res << 1
        // add n%2 to res
		bit := n % 2
		res += bit
        // right shift n by 1
		n = n >> 1
	}
	return res
}