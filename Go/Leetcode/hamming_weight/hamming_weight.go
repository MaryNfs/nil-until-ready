package hamming_weight


func hammingWeight(n int) int {
	var x int
	res := make([]int, 0)
	for n >= 2 {
		x = n % 2
		res = append(res, x)
		n = n / 2
	}
	res = append(res, n)
	count :=0
	for i:=0;i<len(res);i++{
		if res[i] == 1 {
			count++
		}
	}
	return count
}

