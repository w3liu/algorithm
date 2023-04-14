package main

import "fmt"

func main() {
	num := romanToInt("III")
	fmt.Println(num)
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	var preNum = getNum(s[0])
	var total int
	for i := 1; i < len(s); i++ {
		currNum := getNum(s[i])
		if currNum > preNum {
			total -= preNum
		} else {
			total += preNum
		}
		preNum = currNum
	}
	total += preNum
	return total
}

func getNum(ch uint8) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
