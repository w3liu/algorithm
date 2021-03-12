package main

import "math"

func main() {
	var nums = []int{5, 20, 13, 20, 21, 76}
	put(0, 0, nums, 100)
	print(maxW)
}

/**
i 代表放入的第几块石头
cw 代表当前背包里的总重量
nums 代表全部石头
w 表示背包能承受的重量
*/

var maxW = math.MinInt32

func put(i, cw int, nums []int, w int) {
	if i == len(nums) || cw == w {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	put(i+1, cw, nums, w)
	if cw+nums[i] <= w {
		put(i+1, cw+nums[i], nums, w)
	}
}
