// File name: ...\s03\04_for_1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	// to repeat the same operations
	fmt.Println(2, 20)
	fmt.Println(3, 30)
	fmt.Println(4, 40)
	fmt.Println(5, 50)
	fmt.Println(6, 60)
	fmt.Println(7, 70)
	fmt.Println(8, 80)
	fmt.Println(9, 90)
	fmt.Println(10, 100)

	fmt.Println()
	i := 2
	for i < 11 {
		fmt.Println(i, i*10)
		i++
	}
}
