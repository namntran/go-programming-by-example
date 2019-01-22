// File name: ...\s02\10_logical_operators\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var b1 = true
	var b2 = false

	fmt.Println(b1 || b2) // or = true
	fmt.Println(b1 && b2) // and = false
	fmt.Println(b1 && !b2) // and not b2 = true
	fmt.Println(true && !!false) // and not not false = false
}
