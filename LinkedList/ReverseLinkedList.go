package linkedlist

// func reverseList(head *ListNode) *ListNode {

// 	// Insert First LinkedList
// 	insert := &ListNode{
// 		Val:  9,
// 		Next: head,
// 	}

// 	head = insert
// 	cur := head
// 	var result *ListNode

// 	index := 0
// 	for cur != nil {
// 		//Find LinkedList
// 		if cur.Val == 2 {
// 			fmt.Println(cur.Val)
// 		}
// 		// Insert in the middle
// 		if index == 1 {
// 			tmp := cur.Next
// 			ins := &ListNode{
// 				Val:  8,
// 				Next: tmp,
// 			}
// 			fmt.Println(ins)
// 			cur.Next = ins
// 			break
// 		}

// 		tmps := cur.Next
// 		// result.Val = cur.Val
// 		// result.Next = cur.Next
// 		// result.Next = cur
// 		// result = cur
// 		cur = tmps
// 		index++

// 		// tmp := cur.Next
// 		// cur.Next = prev
// 		// prev = cur
// 		// cur = tmp
// 	}

// 	newNode := cur
// 	head.Next = newNode
// 	result = head

// 	return result
// }

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
