package main

import "fmt"

func main() {
	nums := []int{-3, 0, 1, -2}
	v := maxProduct(nums)
	fmt.Println(v)
}

func maxProduct1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max int
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			continue
		}
		if max < v {
			max = v
		}
		for k := i + 1; k < len(nums); k++ {
			v = v * nums[k]
			if max < v {
				max = v
			}
		}
	}
	return max
}

func maxProduct(nums []int) int {
	var maxv, imax, imin int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(nums[i]*imax, nums[i])
		imin = min(nums[i]*imin, nums[i])
		maxv = max(maxv, imax)
	}
	return maxv
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if y > x {
		return x
	}
	return y
}
