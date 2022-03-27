package main

import "math"

func main() {

}

func coinChange(coins []int, amount int) int {
	res := make([]int, amount)
	return calc(coins, amount, res)

}

func calc(coins []int, amount int, res []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if res[amount-1] != 0 {
		return res[amount-1]
	}
	var min = math.MaxInt64
	for _, coin := range coins {
		result := calc(coins, amount-coin, res)
		if result >= 0 && result < min {
			min = 1 + result
		}
	}
	if min == math.MaxInt64 {
		res[amount-1] = -1
	} else {
		res[amount-1] = min
	}
	return res[amount-1]
}
