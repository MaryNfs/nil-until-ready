package path_sum

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	cases := []struct {
		root      []any
		targetSum int
		out       int
	}{
		{[]any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}, 8, 3},
		{[]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}, 22, 3},
	}
	for _, c := range cases {
		fmt.Println("f")
		root := buildTreeLevelOrder(c.root)
		res := pathSum(root, c.targetSum)
		if res != c.out {
			t.Errorf("pathSum(%v,%v)=%v, expected %v.", c.root, c.targetSum, res, c.out)
		}
	}
}

func buildTreeLevelOrder(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	q := []*TreeNode{root}
	i := 1

	for i < len(vals) && len(q) > 0 {
		cur := q[0]
		q = q[1:]

		// left
		if i < len(vals) && vals[i] != nil {
			cur.Left = &TreeNode{Val: vals[i].(int)}
			q = append(q, cur.Left)
		}
		i++

		// right
		if i < len(vals) && vals[i] != nil {
			cur.Right = &TreeNode{Val: vals[i].(int)}
			q = append(q, cur.Right)
		}
		i++
	}

	return root
}
