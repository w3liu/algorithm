package main

import "fmt"

func main() {
	b := judgeCircle("LL")
	fmt.Println("b", b)
}

func judgeCircle(moves string) bool {
	var x, y int
	for _, ele := range moves {
		switch ele {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y--
		case 'D':
			y++
		}
	}
	return x == 0 && y == 0
}
