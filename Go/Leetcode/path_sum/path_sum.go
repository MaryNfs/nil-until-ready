package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://www.youtube.com/watch?v=VDTZiggKlAE

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	lookup := make(map[int]int, 0)
	lookup[targetSum] = 1
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		count += lookup[sum]
		lookup[sum+targetSum] += 1
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		lookup[sum+targetSum] -= 1
	}
	dfs(root, 0)

	return count
}
