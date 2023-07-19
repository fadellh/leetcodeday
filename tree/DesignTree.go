package tree

import (
	"fmt"
)

type Node struct {
	val        int
	leftChild  *Node
	rightChild *Node
}

type Tree struct {
	root *Node
}

func Constructors() Tree {
	return Tree{
		root: nil,
	}
}

func (this *Tree) Insert(val int) {
	if this.root == nil {
		newNode := &Node{
			val:        val,
			leftChild:  nil,
			rightChild: nil,
		}
		this.root = newNode
		return
	}

	current := this.root
	for current != nil {
		if current.val > val {
			if current.leftChild == nil {
				current.leftChild = &Node{
					val:        val,
					leftChild:  nil,
					rightChild: nil,
				}
				break
			}
			current = current.leftChild
			continue
		}

		if current.rightChild == nil {
			current.rightChild = &Node{
				val:        val,
				leftChild:  nil,
				rightChild: nil,
			}
			break
		}

		current = current.rightChild
		continue
	}
}

func (this *Tree) find(val int) bool {
	current := this.root

	for current != nil {
		if current.val == val {
			return true
		}

		if val > current.val {
			current = current.rightChild
			continue
		}
		current = current.leftChild
	}
	return false
}

func (this *Tree) treversePreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	this.treversePreOrder(root.leftChild)
	this.treversePreOrder(root.rightChild)
}

func (this *Tree) treverseInOrder(root *Node) {
	if root == nil {
		return
	}
	this.treverseInOrder(root.leftChild)
	fmt.Println(root.val)
	this.treverseInOrder(root.rightChild)
}

func (this *Tree) treversePostOrder(root *Node) {
	if root == nil {
		return
	}
	this.treversePostOrder(root.leftChild)
	this.treversePostOrder(root.rightChild)
	fmt.Println(root.val)
}

func (this *Tree) Print() {
	// fmt.Println("Root: ", this.root.val)
	// return

	// left := this.root.leftChild
	// for left != nil {
	// 	fmt.Println(this.root.leftChild.val)
	// 	left = left.leftChild
	// }

	// right := this.root.rightChild
	// for right != nil {
	// 	fmt.Println(this.root.rightChild.val)
	// 	right = right.rightChild
	// }
}

func Implemen() {
	obj := Constructors()
	obj.Insert(9)
	obj.Insert(7)
	obj.Insert(6)
	obj.Insert(8)
	obj.Insert(5)
	obj.Insert(4)
	obj.Insert(3)
	obj.Insert(10)
	obj.Insert(19)
	// isFind := obj.find(13)
	// isFinds := obj.find(10)
	obj.treversePostOrder(obj.root)

	// fmt.Println(isFind)
	// fmt.Println(isFinds)

	// obj.Print()
}
