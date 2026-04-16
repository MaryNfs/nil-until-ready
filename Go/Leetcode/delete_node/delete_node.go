package delete_node

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil {
		if root.Val == key {
			if root.Right == nil {
				return root.Left
			} else if root.Left == nil {
				return root.Right
			}
			// find the min
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			root.Val = cur.Val
			root.Right = deleteNode(root.Right, root.Val)
		} else if root.Val > key {
			root.Left = deleteNode(root.Left, key)
		} else if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		}
		return root
	}
	return nil
}
