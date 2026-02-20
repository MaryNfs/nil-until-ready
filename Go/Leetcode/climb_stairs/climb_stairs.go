package climb_stairs

// Time Limit Exceeded
func climbStairs(n int) int {
    if n ==1 || n==2 || n==0 {
        return n
    }
    res := climbStairs(n-2)+climbStairs(n-1)

    return res
}