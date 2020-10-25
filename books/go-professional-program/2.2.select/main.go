package main

import (
	"fmt"
	"time"
)

func main() {
	selectFunc1()
}

func selectFunc() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 2
		time.Sleep(5 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready")
	case <-chan2:
		fmt.Println("chan2 ready")
	default:
		fmt.Println("default")
	}
	fmt.Println("main exit.")
}

func selectFunc1() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready")
	case <-chan2:
		fmt.Println("chan2 ready")
	}
	fmt.Println("main exit.")
}
