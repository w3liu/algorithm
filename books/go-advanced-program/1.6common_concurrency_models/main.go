package main

import (
	"fmt"
	"sync"
)

func main() {
	//var mu sync.Mutex
	//mu.Lock()
	//go func() {
	//	fmt.Println("你好，世界")
	//	mu.Unlock()
	//}()
	//mu.Lock()

	//done := make(chan int)
	//go func() {
	//	fmt.Println("你好，世界")
	//	<-done
	//}()
	//
	//done <- 1

	//done := make(chan int, 1)
	//go func() {
	//	fmt.Println("你好，世界")
	//	done <- 1
	//}()
	//
	//<-done

	//done := make(chan int, 10)
	//
	//for i := 0; i < cap(done); i++ {
	//	go func() {
	//		fmt.Println("你好，世界")
	//		done <- 1
	//	}()
	//}
	//
	//for i := 0; i < 10; i++ {
	//	<-done
	//}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("你好，世界")
		}()
	}

	wg.Wait()

}
