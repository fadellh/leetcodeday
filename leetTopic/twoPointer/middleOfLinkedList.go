package twopointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleOfLinkedList(head *ListNode) *ListNode {
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
