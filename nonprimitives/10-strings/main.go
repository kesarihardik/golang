package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	//string is read only(immutable) sequence of bytes not necessarily UTF-8 encoded.
	s1 := string("Hello, 世界")
	fmt.Printf("bytes in s1: %d, runes in s1: %d\n", len(s1), utf8.RuneCountInString(s1))

	//range iterates over runes in string
	for i, ch := range s1 {
		fmt.Printf("\n index = %v, value = %v, character = %c, Unicode = %U", i, ch, ch, ch)
	}

	/* string is read only slice of bytes.
	bytes(1 byte) can handle ASCII(1 byte/UTF-8) but can't handle unicode(UTF-32).
	by default go reads string using rune.
	*/
	fmt.Print("\n")
	var s = []byte("Hey there! 😃") //change to rune to read Unicode points correctly
	for i, v := range s {
		fmt.Printf("\nindex - %d , character = %c, Unicode = %U\n", i, v, v)
	}

	//package strings
	// fmt.Println(strings.ToLower(s1))
	// fmt.Println(strings.HasPrefix("abc", "t"))
	// fmt.Println(strings.Trim("!!hello!!", "!"))
	// fmt.Println(strings.TrimSpace("  abcd  "))
	// fmt.Println(strings.Contains(s1, "ell"))
	// fmt.Println(strings.Index(s1, "l"))
	// fmt.Println(strings.Replace("Hello there", "there", "hi", 1))
	// fmt.Println(strings.Count(s1, "l"))

	fmt.Println(strconv.Itoa(123))
	if i, err := strconv.Atoi("123"); err == nil {
		fmt.Println(i)
	}

	w := strings.Fields(s1) //strings.Split(s1," ")
	fmt.Println(w)
	fmt.Println(strings.Join(w, ","))

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
