// find maximal value in a slice

package main

import (
	"fmt"
)


func main() {
	slice := []int{23, 45, 54, 67, 89, 23, 12, 45, 30}
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Maximum Value is %d\n Values of slice = %v\n", max,slice)
}
