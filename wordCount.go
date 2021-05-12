// Count each time a word appears in a string

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Needles and Pins 
		needles and pins 
		there you go 
		catch the world but first 
		catch me and then pins`

	fmt.Printf("TEXT IS ==%s\n", text)
	dmap := map[string]int{}
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		dmap[word] = 0
	}

	for keyword, _ := range dmap {
		count := 0
		for _, cword := range words {
			if keyword == cword {
				count++
			}
		}
		dmap[keyword] = count
	}
	for key, value := range dmap {
		fmt.Println(key, value)
	}
}
