// WAP to implement sync concurrency between goroutines.

package main

import (
	"fmt"
	"time"
)

func putVal(c chan int,value int) {
	fmt.Println("Putting in integer")
	time.Sleep(1 * time.Second)
	c <- value
	fmt.Println("Integer kept in channel")
}

func main() {

	fmt.Println("Executing Program ...... ")

	channel := make(chan int)
	defer close(channel)

	for i := 1; i<3; i++{
	go putVal(channel,i)
	}
        for k := 1; k<3; k++{	
	 val := <-channel
	 fmt.Printf("%d\n", val)
	}
}
