// File name: ...\s03\06_for_3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	i := 2
	for { //infinite looop
		if i == 8 {
			break //get out
		}

		if i == 4 || i == 6 {
			i++
			continue //go back to the top/beginnning of the for loop
		}

		// fmt.Print(++i)	//compiler error
		// fmt.Print(i++)	//compiler error

		fmt.Print(i, " ")
		i++
	}

}
