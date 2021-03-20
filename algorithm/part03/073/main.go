package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 1, 2, 5, 10, 10, 1, 5}
	filter(coins, 0, 0, 1, 3)
	fmt.Println(minVal)
}

var minVal = math.MaxInt64

func filter(coins []int, i, val, num, w int) {
	if i == len(coins)-1 || val == w {
		if val == w {
			if minVal > num {
				minVal = num
			}
		}
		return
	}
	filter(coins, i+1, val, num, w) // 不放入
	if num+coins[i+1] <= w {
		filter(coins, i+1, val+coins[i+1], num+1, w) // 放入
	}
}
