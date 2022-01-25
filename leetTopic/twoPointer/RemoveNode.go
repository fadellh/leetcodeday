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
