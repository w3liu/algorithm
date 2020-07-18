package main

import "fmt"

func main() {
	printChar("({[]})")
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
	l := len(s)
	arr := make([]int32, 0, l)
	for _, item := range s {
		if item == 40 || item == 123 || item == 91 {
			arr = append(arr, item)
		} else {
			if len(arr) > 0 && ((arr[len(arr)-1] == 91 && item == 93) ||
				(arr[len(arr)-1] == 123 && item == 125) ||
				(arr[len(arr)-1] == 40 && item == 41)) {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}

func printChar(s string) {
	for _, item := range s {
		fmt.Println(item)
	}
}
