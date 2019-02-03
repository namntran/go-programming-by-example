//program to calculate what day of the week it will be in future
package main

import "fmt"

func main() {
	today := 0 // Sun = 0, Mon = 1, ..., Sat = 6
	weeklength := 9
	futureday := (today + weeklength) % 7
	fmt.Printf("today=%d , weeklength=%d, future day=%d\n", today, weeklength, futureday)

}
