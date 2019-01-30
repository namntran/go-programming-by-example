package main

import "fmt"

func main() {
	table()
}
func table() {
	i := 2
	fmt.Println("i \t ix2 \t ix3 ")
	fmt.Println("--- \t --- \t ---")
	fmt.Printf("%d\t %d \t %d \t \n", i, i*2, i*3)
	i++
	fmt.Printf("%d\t %d \t %d \t \n", i, i*2, i*3)
	i++
	fmt.Printf("%d\t %d \t %d \t \n", i, i*2, i*3)
}
