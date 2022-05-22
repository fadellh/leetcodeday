package twopointer

func RemoveNode(head *ListNode, n int) *ListNode {
	t := 0
	arr := [100]ListNode{}

	for head != nil {
		arr[t] = *head
		head = (*head).Next
		t++
	}
	arr2 := [100]ListNode{}
	for idx := range arr {
		if idx == t-2 {
			bef := arr[idx-1]
			aft := arr[idx+1]
			bef.Next = &aft
			arr2[idx] = bef
			break
		}

		arr2[idx] = arr[idx]
	}
	return &arr2[0]
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
