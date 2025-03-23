package main

import "fmt"

//array,slice,func,pointer are non-primitive data types
func main() {
	arr := [5]int{1, 2, 4, 3, 5} //  var arr [5]int = [5]int{1, 2, 4, 3, 5}
	fmt.Println(arr)

	for _, v := range arr { // for i,v := range arr ;; i-index, v-copy of index value
		fmt.Println(v)
	}

	//arr:= [5]int{}              // arr initialized with value 0
}
