package main

import "fmt"

func main() {
	Copy()
}

func Question1() {
	var array [10]int
	var slice = array[5:6]
	fmt.Println("slice len: ", len(slice))
	fmt.Println("slice cap: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}

func Question2() {
	var slice []int
	slice = append(slice, 1, 2, 3)
	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0])
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func Question3() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(pollorder) = ", cap(lockorder))
}

func Copy() {
	arr1 := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		arr1[i] = i + 1 + 10
	}

	arr2 := make([]int, 5, 5)

	for i := 0; i < 5; i++ {
		arr2[i] = i + 1
	}
	copy(arr1, arr2)

	for i, _ := range arr2 {
		fmt.Println(arr2[i])
	}
}
