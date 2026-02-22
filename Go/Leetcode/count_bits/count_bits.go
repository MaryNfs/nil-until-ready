package countBits

// Ugly
func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		binary := binary_Calc(i)
		res = append(res, binary)
	}
	return res
}

func binary_Calc(n int) int {
	res := make([]int, 0)
	for n >= 2 {
		mod := n % 2
		res = append(res, mod)
		n = n / 2

	}
	res = append(res, n)
	count := 0
	for i := 0; i < len(res); i++ {
		if res[i] == 1 {
			count++
		}
	}
	return count
}
