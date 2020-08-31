package main

import "fmt"

func main() {
	fmt.Println(deferFuncReturn())
}

func deferFuncParameter() {
	var aInt = 1
	defer fmt.Println(aInt)
	aInt = 2
	return
}

func printArray(arr *[3]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func deferFuncParameter1() {
	var aArr = [3]int{1, 2, 3}
	defer printArray(&aArr)
	aArr[0] = 10
	return
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}
