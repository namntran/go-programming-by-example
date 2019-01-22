// File name: ...\s02\21_utf8_3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	s := "𠜎"

	bytes := []byte(s)

	fmt.Println(bytes)
	fmt.Printf("%U \n", bytes) //%U unicode

	for _, b := range bytes { // 'range' returns an 'index' and a 'value in each iteration. Underscore is used to discard the values returned in each iteration as the 'index'.
		fmt.Printf("%d ", b) //%d is decimal
	}
	fmt.Println()
	for _, b := range bytes {
		fmt.Printf("%o ", b) //%o is octo
	}
	fmt.Println()
	for _, b := range bytes {
		fmt.Printf("%x ", b) // %x is hexadeximal
	}

	fmt.Println()
	fmt.Println("\xf0\xa0\x9c\x8e")
	fmt.Println("\360\240\234\216")

	ch := '𠜎'
	fmt.Printf("%q \n", ch) //q is character
	fmt.Printf("%+q \n", ch) //+q is unicode

	fmt.Println("\U0002070e") // \U prints character

	fmt.Println(string(0x0002070e)) 
}
