package twopointer

func RemoveNode(head *ListNode, n int) *ListNode {
	var count int
	curr := head
	for curr != nil {
		curr = curr.Next
		count++
	}

	curr = head
	for i := 1; i < count-n; i++ {
		curr = curr.Next
	}

	if count == n {
		head = head.Next
	} else if curr.Next != nil {
		curr.Next = curr.Next.Next
	}

	return head
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	idx := 0
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		idx++
	}

	idx = idx - n
	i := 0
	var result ListNode = *head
	for head != nil {
		if i == idx {
			temp = head
			result.Next = temp
		}
		head = head.Next
		i++
	}
	return &result
}
