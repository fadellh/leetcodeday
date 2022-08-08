package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	tail := newList

	for list1 != nil && list2 != nil {
		if list1 != nil && list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else if list2 != nil {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return newList.Next
}
