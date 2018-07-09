package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{v, nil, nil}
	} else if v < root.value {
		root.left = insert(root.left, v)
	} else {
		root.right = insert(root.right, v)
	}
	return root
}

//add pre-order and post-order traversals

func inTraverse(root *Node) {
	if root == nil {
		return
	}
	inTraverse(root.left)
	fmt.Printf("%d ", root.value)
	inTraverse(root.right)
}

func main() {
	var treeRoot *Node
	rand.Seed(time.Now().UnixNano())
	n := 6
	var a [6]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(20) + 1
	}
	fmt.Println("Array of integer: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		treeRoot = insert(treeRoot, a[i])
	}
	inTraverse(treeRoot)
	fmt.Println()
}
