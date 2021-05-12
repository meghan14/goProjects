// how many even ended nos. result from multiplying two 4 digit nos. ?

package main

import (
	"fmt"
)

func isEvenEnded(s string) bool {
	l := len(s)
	if s[0] == s[l-1] {
		return true
	} else {
		return false
	}
}

func main() {
	count := 0
	var numS string
	for x := 1000; x <= 9999; x++ {
	for y := x; y <= 9999; y++ {
		numS = fmt.Sprintf("%d", x*y)
		if isEvenEnded(numS) {
			count++
		}
	  }
	}
	fmt.Printf("The number of even ended numbers = %d , type %T\n", count, count)
}
