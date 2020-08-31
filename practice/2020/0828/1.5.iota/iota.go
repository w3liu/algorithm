package main

import "fmt"

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift     = iota
	starvationThresholds = 1e6
)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1
	_, _
	bit3, mask3
)

func main() {
	fmt.Printf("mutexLocked:%v\n", mutexLocked)
	fmt.Printf("mutexWoken:%v\n", mutexWoken)
	fmt.Printf("mutexStarving:%v\n", mutexStarving)
	fmt.Printf("mutexWaiterShift:%v\n", mutexWaiterShift)
	fmt.Printf("starvationThresholds:%v\n", starvationThresholds)

	fmt.Printf("bit0:%v mask0:%v\n", bit0, mask0)
	fmt.Printf("bit1:%v mask1:%v\n", bit1, mask1)
	fmt.Printf("bit3:%v mask3:%v\n", bit3, mask3)
}
