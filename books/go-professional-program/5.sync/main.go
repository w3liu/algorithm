package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	DoWaitGroupProcess()
}

func ChannelProcess(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
}

func DoChannelProcess() {
	channels := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go ChannelProcess(channels[i])
	}
	for i, ch := range channels {
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}

func DoWaitGroupProcess() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine 1 finished!")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine 2 finished!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutine finished!")
	// ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
}
