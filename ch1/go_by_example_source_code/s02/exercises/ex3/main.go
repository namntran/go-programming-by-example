// s02/exercises/ex3
// write a program to calculate area of a cirle

package main

import "fmt"

func main() {
	fmt.Println("Program to calculate area of a circle")
	fmt.Println("Enter the radius of the circle: ")
	circle()
}
func circle() {
	const PI float64 = 3.14159265359
	var r float64
	var area float64
	fmt.Scanln(&r)
	area = PI * r * r
	fmt.Println("Area of a circle is: ", area)
}
