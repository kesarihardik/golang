package main

import "fmt"

//generic function
func findIndex[T comparable](a []T, val T) int {
	for i, v := range a {
		if v == val {
			return i
		}
	}
	return -1
}

func main() {
	var c []any = []any{12, "afb"} //generic array
	fmt.Println(c)

	a := []int{12, 14, 1, 32, 4}
	fmt.Println(findIndex(a, 32))

	b := []string{"dad", "mom"}
	fmt.Println(findIndex(b, "dad"))

}
