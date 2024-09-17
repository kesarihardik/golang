package main

import "fmt"

func main() {
	// var mp map[int]int        //nil map
	var mp1 map[string]uint8 = make(map[string]uint8) //map[key]value
	fmt.Println(mp1)

	//unlike struct, field names are mandatory in map literals if they are not type names(struct, string etc)
	mp := map[string]uint16{"John": 34, "Jane": 23}
	mp["Venice"] = 33
	delete(mp, "John")
	fmt.Println(mp)

	age, ok := mp["John"]
	if ok {
		fmt.Println(age)
	}

	for name := range mp {
		fmt.Printf("\n name: %s, age: %d", name, mp[name])
	}
}
