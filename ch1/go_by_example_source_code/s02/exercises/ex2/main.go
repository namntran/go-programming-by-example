// s02/ex2/fahrenheit_to_celsius

package main

import "fmt"

func main() {
	fmt.Println("Program to calculate Fahrenheit into Celsius")
	fmt.Println("Enter the temperature in Fahrenheit: ")
	celsius()
}
func celsius() {
	var f float32
	var c float32
	fmt.Scanf("%f", &f)
	c = (f - 32) * 5 / 9
	fmt.Println("Fahrenheit: ", f, "\nCelsius:", c)
}

// Fahrenheit to Celsius
// (32°F − 32) × 5/9 = 0°C
// Deduct 32, then multiply by 5, then divide by 9
