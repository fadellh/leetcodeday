package linkedlist

import (
	"fmt"
)

// Basic of Linked List is pointer. So, you must understand how pointer work in memory

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructors() MyLinkedList {
	return MyLinkedList{
		head: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head

	idx := 0
	for cur != nil {
		if idx == index {
			return cur.val
		}
		cur = cur.next
		idx++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val:  val,
		next: this.head,
	}
	this.head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newTail := &Node{
		val:  val,
		next: nil,
	}

	cur := this.head

	if cur == nil {
		this.head = newTail
	}

	for cur != nil {
		if cur.next == nil {
			cur.next = newTail
			return
		}
		cur = cur.next
	}

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.head
	// cur := dummy

	if index == 0 {
		newNode := &Node{
			val:  val,
			next: this.head,
		}
		this.head = newNode
		return
	}

	idx := 1
	for cur != nil {
		if index == idx {
			newNode := &Node{
				val:  val,
				next: cur.next,
			}
			cur.next = newNode
			return
		}
		cur = cur.next
		idx++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head
	if cur == nil {
		return
	}
	if index == 0 {
		if cur.next != nil {
			newNode := &Node{
				val:  cur.next.val,
				next: cur.next.next,
			}
			this.head = newNode
			return
		}
		this.head = nil
		return
	}

	idx := 0
	prev := this.head
	for cur != nil {
		if index == idx {
			prev.next = cur.next
		}
		prev = cur
		cur = cur.next
		idx++
	}
}

func (this *MyLinkedList) PrintIndex() {
	idx := 0
	cur := this.head

	fmt.Println("===Looping===")

	for cur != nil {
		fmt.Println("===Index:", (idx))

		fmt.Printf("Memory address of cur: %p\n", cur)
		fmt.Println("val", cur.val)
		fmt.Printf("Memory address of cur next: %p\n", cur.next)
		cur = cur.next
		idx++
	}
}

func Implemen() {
	obj := Constructors()
	// obj.AddAtHead(3)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	// obj.AddAtTail(5)
	obj.AddAtTail(1)
	// obj.AddAtIndex(3, 100)
	// obj.DeleteAtIndex(0)

	// res := obj.Get(3)
	// fmt.Println(res)
	obj.PrintIndex()
	// obj.AddAtTail(2)
	// obj.AddAtTail(3)
	resi := obj.Get(0)
	fmt.Println(resi)
	// res = obj.Get(2)
	// fmt.Println(res)
	// res = obj.Get(3)
	// fmt.Println(res)
	// res = obj.Get(5)
	// fmt.Println(res)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
