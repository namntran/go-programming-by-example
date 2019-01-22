// File name: ...\s02\23_utf8_5\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	s := "こんにちは"

	fmt.Printf("% x\n\n", s)

	r := []rune(s) //rune is an alias for int32 which is 4 bytes
	// rune is used to distinguish integer values from character values
	fmt.Printf("%x\n", r)

	fmt.Println(string(r))
}
