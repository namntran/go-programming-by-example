// s02/ex2/fahrenheit_to_celsius

package main

import "fmt"

func main() {
	fmt.Println("Program for Fahrenheit into Celsius")
	var Fahrenheit float32 = 86
	var Celsius float32
	Celsius = (Fahrenheit - 32) * 5 / 9
	fmt.Println("Fahrenheit: ", Fahrenheit, "\nCelsius:", Celsius)

}

// Fahrenheit to Celsius
// (32°F − 32) × 5/9 = 0°C
// Deduct 32, then multiply by 5, then divide by 9
