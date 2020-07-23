package main

import "fmt"

func main() {
	b := backspaceCompare("a#c", "a##c")
	fmt.Println(b)
}

func backspaceCompare(s string, t string) bool {
	sArr := genArr(s)
	tArr := genArr(t)
	return string(sArr) == string(tArr)
}

func genArr(s string) []int32 {
	arr := make([]int32, 0, len(s))
	for _, item := range s {
		if item == '#' {
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}
		} else {
			arr = append(arr, item)
		}
	}
	return arr
}
