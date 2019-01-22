// File name: ...\s02\20_utf8_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	s := "こんにちは"

	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)

	}

	fmt.Println([]byte(s))

	fmt.Println()
	s = "𠜎ABCD"
	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}
	fmt.Println([]byte(s))
}
/* rune is an alias for int32(4 bytes)
rune is used to distinguish character values from integer values
utf8 uses 1-4 bytes. Go uses utf8 (Unicode Transformation Format) as its text encoding scheme
rune uses 4 bytes
*/