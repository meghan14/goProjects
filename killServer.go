// Read a process ID from file, convert to integer and print killing <PID>

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

func killServer(pidFile string) (int, error) {
	pidInt, err := strconv.Atoi(pidFile)
	if err != nil {
		return 0, errors.Wrap(err, "String conversion failed")
	}

	return pidInt, nil
}

func main() {
	pid := "52aa42"
	id, err := killServer(pid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", id)

}
