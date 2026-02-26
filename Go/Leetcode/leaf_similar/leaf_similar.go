package leaf_similar

import (
	"slices"
)

// https://www.youtube.com/watch?v=Nr8dbnL0_cM
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := make([]int, 0)
	res2 := make([]int, 0)

	res1 = dfs(root1, res1)
	res2 = dfs(root2, res2)

	if slices.Equal(res1, res2) {
		return true
	}
	return false
}

func dfs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
	}
	res = dfs(root.Left, res)
	res = dfs(root.Right, res)

	return res
}
