package main

import (
	"fmt"
	"time"
)

func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go addNumToChan(chan1)
	go addNumToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		}
	}

}

func addNumToChan(chanName chan<- int) {
	for {
		chanName <- 1
		time.Sleep(time.Second)
	}
}

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}
