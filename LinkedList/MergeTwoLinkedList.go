package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	ptr1, ptr2 := list1, list2
	result := new(ListNode)
	tail := result

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			tail.Next = ptr1
			tail = tail.Next
			ptr1 = ptr1.Next
			continue
		}
		tail.Next = ptr2
		tail = tail.Next
		ptr2 = ptr2.Next
	}

	for ptr1 != nil {
		tail.Next = ptr1
		tail = tail.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		tail.Next = ptr2
		tail = tail.Next
		ptr2 = ptr2.Next
	}

	result = result.Next
	return result
}
