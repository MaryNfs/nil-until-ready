package num_squares

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i:=0;i<=n;i++{
		dp[i]=n
	}
	dp[0]=0

	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			sq:= j*j
			if i-sq <0 {
				break
			}
			dp[i]=min(dp[i],dp[i-sq]+1)
		}
	}

	return dp[n]
}
