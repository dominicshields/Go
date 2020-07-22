package main

import (
	"fmt"
)

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type ItemBinarySearchTree struct {
	root *Node
}

var bst ItemBinarySearchTree

func main() {

	bst.Insert(1, 1)
	bst.Insert(87, 5)
	bst.Insert(22, 5)
	bst.Insert(88, 5)
	bst.Insert(13, 5)
	bst.Insert(4, 5)
	bst.Insert(2, 2)
	bst.Insert(45, 5)
	bst.Insert(7, 3)
	bst.Insert(3, 5)
	bst.Insert(11, 5)
	bst.Insert(98, 5)
	bst.Insert(90, 5)
	bst.Insert(89, 5)
	bst.Insert(8, 5)
	bst.String()
}

func (bst *ItemBinarySearchTree) Insert(key int, value int) {
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
	bst.String()
}

func insertNode(node, newNode *Node) {
fmt.Printf("Newnode.key = %v, Node.key = %v\n",newNode.key , node.key)
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *ItemBinarySearchTree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}
