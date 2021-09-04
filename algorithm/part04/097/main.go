package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}
	var res = subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	var result = make([][]int, 0)
	back(0, []int{}, nums, &result)
	return result
}

// 回溯
func back(i int, tmp, nums []int, result *[][]int) {
	// result需要传递指针
	*result = append(*result, tmp)
	for j := i; j < len(nums); j++ {
		tmp = append(tmp, nums[j])
		back(j+1, tmp, nums, result)
		dst := make([]int, len(tmp)-1)
		copy(dst, tmp) // 这里需要深拷贝
		tmp = dst
	}
}
