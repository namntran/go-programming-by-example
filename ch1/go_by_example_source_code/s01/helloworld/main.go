// line comment
/*
multi-line comment
*/

package main // if you need to write a program and execute - needs package 'main'. Not needed for a library

import "fmt"

func main() { // entry point into any executable file (code)
	// fmt.Println("Hello World") // from the 'fmt' (format) library, call the 'Println' function
	// fmt.Println("Chinese 漢字")
	// fmt.Println("Arabic مرحبا.")
	// fmt.Println("Russian Стой на месте")
	helloworld()
	firstName()
	concepts()
	variables()
	tab()
}
func helloworld() {
	fmt.Println("Hello World") // from the 'fmt' (format) library, call the 'Println' function
	fmt.Println("Chinese 漢字")
	fmt.Println("Arabic مرحبا")
	fmt.Println("Russian Стой на месте")
}
func firstName() {
	firstName := "Jim\n\n"
	fmt.Println("Name: ", firstName)
}
func concepts() {
	k := 10
	K := 14
	fmt.Println("Print k:", k, "Print K:", K)
	fmt.Println(4 + 2*3)
}
func variables() {
	var i = 20
	i++
	fmt.Println(i)
}
func tab() {
	fmt.Println("John\tJack\nTim")
	i := 10
	fmt.Printf("%d %x %T\n", i, i, i)
}
