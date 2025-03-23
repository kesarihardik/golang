package main

import "fmt"

func main() {

	//slice is dynamically sized, flexible view into the elements of an array.
	//slice capacity = original array length

	vowels := [5]rune{'a', 'e', 'i', 'o', 'u'} //array literal
	vowelSlice := vowels[0:3]                  // slice from index 0 to 2
	fmt.Printf("%v , len: %d,capacity: %d", vowelSlice, len(vowelSlice), cap(vowelSlice))

	// They don't store data. They are a reference.
	//slices are referrence to array. they modify original array
	vowelSlice[0] = 'b'
	fmt.Printf("%c", vowels[0])

	//slices are array types without length
	v := []int{1, 2, 3, 4, 5} //array allocated, modified to slice,reference allocated to v
	fmt.Printf("\n%T", v)

	a := make([]int, 5, 6) //slice                make(type,length, capacity)
	a = a[:3]              //reslice - set length 3 [0,3)
	printSlice(a)

	//nil slice
	var s []int
	printSlice(s)
	s = append(s, 1, 2, 3) //automatically allocates bigger array in case array is full.
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("\nlength=%d capacity=%d %v", len(s), cap(s), s)
}
