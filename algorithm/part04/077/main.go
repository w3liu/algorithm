package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, -1, -1, 0}
	var results = threeSum(nums)
	fmt.Println(results)
}

func threeSum1(nums []int) [][]int {
	var l = len(nums)
	var e1, e2, e3 int
	var results = make([][]int, 0)
	var cache = make(map[string]struct{})
	for i := 0; i < l-2; i++ {
		e1 = nums[i]
		for j := i + 1; j < l-1; j++ {
			e2 = nums[j]
			for k := j + 1; k < l; k++ {
				e3 = nums[k]
				if e1+e2+e3 == 0 {
					result := []int{e1, e2, e3}
					sort.Ints(result)
					key := fmt.Sprintf("%d_%d_%d", result[0], result[1], result[2])
					if _, ok := cache[key]; ok {
						continue
					}
					results = append(results, result)
					cache[key] = struct{}{}
				}
			}
		}
	}
	return results
}

func threeSum(nums []int) [][]int {
	var n = len(nums)
	var results = make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			res := nums[i] + nums[l] + nums[r]
			if res == 0 {
				result := []int{nums[i], nums[l], nums[r]}
				results = append(results, result)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if res > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return results
}
