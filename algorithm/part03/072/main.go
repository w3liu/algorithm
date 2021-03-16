package main

import (
	"fmt"
	"math"
)

func main() {
	//var nums = []int{5, 20, 13, 20, 21, 76}
	//put(0, 0, nums, 100)
	//print(maxW)
	var nums = []int{5, 20, 13, 20, 21, 76}
	result := knapsack1(nums, len(nums), 100)
	fmt.Println("result", result)
}

/**
i 代表放入的第几块石头
cw 代表当前背包里的总重量
nums 代表全部石头
w 表示背包能承受的重量
*/

var maxW = math.MinInt32
var memo = make(map[string]struct{})

func put(i, cw int, nums []int, w int) {
	if i == len(nums) || cw == w {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	key := fmt.Sprintf("%d_%d", i, cw)
	if _, ok := memo[key]; ok {
		return
	}
	memo[key] = struct{}{}
	put(i+1, cw, nums, w)
	if cw+nums[i] <= w {
		put(i+1, cw+nums[i], nums, w)
	}
}

// 动态规划
func knapsack(weight []int, n, w int) int {
	var states = make([][]bool, 0)
	for i := 0; i < n; i++ {
		row := make([]bool, w+1)
		states = append(states, row)
	}
	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j < w; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}
	for j := w; j >= 0; j-- {
		if states[n-1][j] {
			return j
		}
	}
	return 0
}

func knapsack1(weight []int, n, w int) int {
	var states = make([]bool, w+1)
	if weight[0] <= w {
		states[0] = true
	}
	for i := 1; i < n; i++ {
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] {
				states[j+weight[i]] = true
			}
		}
	}
	for j := w; j >= 0; j-- {
		if states[j] {
			return j
		}
	}
	return 0
}
