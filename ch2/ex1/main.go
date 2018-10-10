package main

import "fmt"

func main() {
	concepts()
}

func concepts() {
	firstName := "John"
	firstName = `Nam` //back ticks not '(next to 1 to the left)
	fmt.Println(firstName)
}
