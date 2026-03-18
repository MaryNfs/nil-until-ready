package reverse_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newhead *ListNode
	if head == nil {
		return newhead
	}
	tail := &ListNode{Val: 0, Next: nil}
	newhead = tail
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i := len(res) - 1; i >= 0; i-- {
		tail.Val = res[i]
		if i != 0 {
			tail.Next = &ListNode{Val: 0, Next: nil}
			tail = tail.Next
		}
	}

	return newhead
}
