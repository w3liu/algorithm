package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	result := convert(s, 2)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var letters = make([][]byte, numRows)
	var x int
	var direct int // 0 表示向下，1 表示向上
	for _, v := range s {
		letters[x] = append(letters[x], byte(v))
		if x == numRows-1 {
			direct = 1
		}
		if x == 0 {
			direct = 0
		}
		if direct == 0 {
			x++
		} else {
			x--
		}
	}
	var result []byte
	for i := 0; i < len(letters); i++ {
		result = append(result, letters[i]...)
	}
	return string(result)
}
