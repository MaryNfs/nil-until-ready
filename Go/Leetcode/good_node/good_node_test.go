package good_node

import "testing"

func TesGoodNodes(t *testing.T) {
	cases := []struct {
		root []any
		out  int
	}{
		{[]any{3, 1, 4, 3, nil, 1, 5}, 4},
		{[]any{3, 3, nil, 4, 2}, 3},
		{[]any{1}, 1},
	}
	for _, c := range cases {
		root := buildTreeLevelOrder(c.root)
		res := goodNodes(root)
		if res != c.out {
			t.Errorf("goodNodes(%v)=%v, expected %v.", c.root, res, c.out)
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
