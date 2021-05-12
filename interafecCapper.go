// WAP for Capper interfact to convert text to uppercase

package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer
}

func (cap *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	output := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		output[i] = c
	}
	return cap.wtr.Write(output)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello World!")

}
