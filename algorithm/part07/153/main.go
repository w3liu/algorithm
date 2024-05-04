package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	result := threeSumClosest1(nums, 1)
	fmt.Println(result)
}

func threeSumClosest(nums []int, target int) int {
	var result = math.MaxInt64
	var delta = math.MaxInt64
	var size = len(nums)
	for x := 0; x < size-2; x++ {
		for y := x + 1; y < size-1; y++ {
			for z := y + 1; z < size; z++ {
				val := nums[x] + nums[y] + nums[z]
				sub := getAbs(val - target)
				if delta > sub {
					delta = sub
					result = val
				}
			}
		}
	}
	return result
}

func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	var result = math.MaxInt64
	var delta = math.MaxInt64
	var size = len(nums)
	for x := 0; x < size-2; x++ {
		y := x + 1
		z := size - 1
		for y < z {
			val := nums[x] + nums[y] + nums[z]
			if val == target {
				return val
			}
			sub := getAbs(val - target)
			if delta > sub {
				delta = sub
				result = val
			}
			if val > target {
				z--
			} else {
				y++
			}
		}
	}
	return result
}

func getAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func threeSumClosest2(nums []int, target int) (ans int) {
	n := len(nums)
	if n == 3 {
		ans = nums[0] + nums[1] + nums[2]
		return
	}
	sort.Ints(nums)
	minDiff := math.MaxInt

	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}

		s := x + nums[i+1] + nums[i+2]
		if s > target {
			if s-target < minDiff {
				ans = s
			}
			break
		}

		s = x + nums[n-2] + nums[n-1]
		if s < target {
			if target-s < minDiff {
				minDiff = target - s
				ans = s
			}
			continue
		}

		j, k := i+1, n-1
		for j < k {
			s = x + nums[j] + nums[k]

			if s == target {
				return target
			}

			if s > target {
				if s-target < minDiff {
					minDiff = s - target
					ans = s
				}
				k--
			} else {
				if target-s < minDiff {
					minDiff = target - s
					ans = s
				}
				j++
			}
		}
	}
	return
}
