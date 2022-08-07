package twopointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleOfLinkedList(head *ListNode) *ListNode {
	t := 0
	arr := [100]ListNode{}

	for head != nil {
		arr[t] = *head
		head = (*head).Next
		t++
	}

	half := int(t / 2)
	return &arr[half]
}

//MidleNode another solution.
//Jadi fast berguna untuk meningkat kecepatanya 2x lipat.
//Dan slow akan terus di overwrite sampe pada titik dia di tengah
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	return slow
}
