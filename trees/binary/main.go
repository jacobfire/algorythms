package main

import "fmt"

type Node struct {
	key   string
	value int
	left  *Node
	right *Node
}

func New(k string, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

func insert(node *Node, key string, value int) {
	if key < node.key {
		if node.left == nil {
			node.left = New(key, value)
		} else {
			insert(node.left, key, value)
		}
	}
	if key > node.key {
		if node.right == nil {
			node.right = New(key, value)
		} else {
			insert(node.right, key, value)
		}
	}
}

func main() {
	root := New("first", 1)
	insert(root, "second", 2)

	fmt.Printf("%+v", root)
}
