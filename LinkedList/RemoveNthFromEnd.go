package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	left := dummy
	right := head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	//delete Node
	left.Next = left.Next.Next
	return dummy.Next

}
