// File name: ...\s02\22_utf8_4\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello𠜎ち"
	fmt.Println(len(s)) //len(s) returns the number of bytes in string 's' (hello has 5 bytes, japanese characters has 4 bytes and 3 bytes)
	fmt.Println(utf8.RuneCountInString(s)) //RuneCountInString(s) returns the number of runes in string 's'
}
