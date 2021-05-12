// fizzBuzz :
// div by 3 -> fizz
// div by 5 -> buzz
// div by both (15) -> fizzbuzz

package main

import (
	"fmt"
)

func main() {

	for x := 1; x < 21; x++ {
		flag := 0
		if x%3 == 0 {
			fmt.Print("fizz")
			flag = 1
		}
		if x%5 == 0 {
			fmt.Print("buzz")
			flag = 1
		}
		if flag != 1 {
			fmt.Println(x)
		} else {
			fmt.Println()
		}
	}
}
