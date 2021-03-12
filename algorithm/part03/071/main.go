package main

import "fmt"

func main() {
	cal8queens(0)
}

var result = make([]int, 8)

var found bool

func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		found = true
		return
	}
	for j := 0; j < 8; j++ {
		if isOk(row, j) {
			result[row] = j
			cal8queens(row + 1)
			if found {
				return
			}
		}
	}
}

func isOk(row, column int) bool {
	left := column - 1
	right := column + 1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if left >= 0 && result[i] == left {
			return false
		}
		if right < 8 && result[i] == right {
			return false
		}
		left--
		right++
	}
	return true
}

func printQueens(result []int) {
	for i := 0; i < len(result); i++ {
		s := ""
		for j := 0; j < 8; j++ {
			if result[i] == j {
				s = fmt.Sprintf("%s %s", s, "*")
			} else {
				s = fmt.Sprintf("%s %s", s, "-")
			}
		}
		fmt.Println(s)
	}
	fmt.Println("-------------------")
}
