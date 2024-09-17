package main

import "fmt"

type Student struct {
	name string
	age  int
}

//a method is a function with special receiver argument. receiver type should be in same package.
//receiver can be non-struct
//method associates with a type

func (s Student) printStudentDetails() {
	fmt.Println(s.age, s.name)
}

func printStudentName(s Student) {
	fmt.Println(s.name)
}

func main() {
	s := Student{name: "John", age: 23}
	printStudentName(s)     //function
	s.printStudentDetails() //method
}
