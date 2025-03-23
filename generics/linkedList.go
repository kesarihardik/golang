package main

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value, nil}
}
