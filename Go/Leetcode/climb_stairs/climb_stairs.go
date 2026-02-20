package climb_stairs

// Time Limit Exceeded
// func climbStairs(n int) int {
//     if n ==1 || n==2 || n==0 {
//         return n
//     }
//     res := climbStairs(n-2)+climbStairs(n-1)

//     return res
// }

// Magic way
// https://www.youtube.com/shorts/x4epBH6HF10
// func climbStairs(n int) int {
// 	one := 1
// 	two := 1

// 	for i := 0; i < n-1; i++ {
// 		one, two = two, two+one
// 	}
// 	return two
// }

// Third try
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	one, two := 1, 2
	var res int
	
	for i := 2; i <= n-1; i++ {
		res = one + two
		one = two
		two = res
	}
	return res
}
