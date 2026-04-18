package tribonacci

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		sum := make([]int, n+1)
		sum[0] = 0
		sum[1] = 1
		sum[2] = 1

		for i := 3; i <= n; i++ {
			sum[i] = sum[i-1] + sum[i-2] + sum[i-3]
		}
		return sum[n]
	}

}