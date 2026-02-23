package count_bits

// Ugly
// func countBits(n int) []int {
// 	res := make([]int, 0)
// 	for i := 0; i <= n; i++ {
// 		binary := binary_Calc(i)
// 		res = append(res, binary)
// 	}
// 	return res
// }

// func binary_Calc(n int) int {
// 	res := make([]int, 0)
// 	for n >= 2 {
// 		mod := n % 2
// 		res = append(res, mod)
// 		n = n / 2

// 	}
// 	res = append(res, n)
// 	count := 0
// 	for i := 0; i < len(res); i++ {
// 		if res[i] == 1 {
// 			count++
// 		}
// 	}
// 	return count
// }

// DP : https://www.youtube.com/watch?v=RyBM56RIWrM
func countBits(n int) []int {
	dp := make([]int, n+1)
	off := 1
	for i := 1; i < n+1; i++ {
		if off*2 == i {
			off = i
		}
		dp[i] = 1 + dp[i-off]
	}
	return dp
}
