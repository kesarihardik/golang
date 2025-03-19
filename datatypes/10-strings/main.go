package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//string is a sequence of bytes. len() counts number of bytes in string.
	s1 := string("Hello, 世界")
	fmt.Printf("bytes in s1: %d, runes in s1: %d", len(s1), utf8.RuneCountInString(s1))

	//these characters can not be represented by 8 bits.
	// var r byte = byte('世')                  //overflow error
	fmt.Printf("\n%T  %d", '世', '世')

	// string is read only slice of bytes
	var s = []rune("Résumé")
	var indexed = s[5]
	fmt.Printf("value: %v, character: %c , type: %T , Unicode format: %U", indexed, indexed, indexed, indexed)

	for i, v := range s {
		fmt.Println(i, v)
	}

	//strings are immmutable in go

	// var strSlice = []string{"s", "t", "r"}
	// var strBuilder strings.Builder
	// for i := range strSlice {
	// 	strBuilder.WriteString(strSlice[i])
	// }
	// var concatedStr = strBuilder.String()
	// fmt.Println(concatedStr)

	// sentence := "How was your day?"
	// words := strings.Fields(sentence)
	// fmt.Println(words)
}
