package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value, nil}
}

func PrintList[T any](node *Node[T]) {
	if node == nil {
		return
	}
	for node != nil {
		fmt.Printf("%v -> ", node.val)
		node = node.next
	}
	fmt.Print("<nil> \n")
}
