package main

import (
	"errors"
	"fmt"
)

func Index[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	head := NewNode(1)
	head.next = NewNode(2)
	head.next.next = NewNode(3)
	PrintList(head)
}
