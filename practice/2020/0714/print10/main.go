package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 开启10个线程打印1到10
func main() {
	var num int32
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt32(&num, 1)
			fmt.Println(num)
		}()
	}
	time.Sleep(10)
}
