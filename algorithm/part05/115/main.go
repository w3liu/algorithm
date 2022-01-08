package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	result := maxSlidingWindow(nums, 3)
	fmt.Println(result)
}

//func maxSlidingWindow(nums []int, k int) []int {
//	if len(nums) < k {
//		return []int{max(nums)}
//	}
//	var result = make([]int, 0)
//	for i := 0; i <= len(nums)-k; i++ {
//		result = append(result, max(nums[i:i+k]))
//	}
//	return result
//}
//
//func max(nums []int) int {
//	var max = nums[0]
//	for i := 0; i < len(nums); i++ {
//		if max < nums[i] {
//			max = nums[i]
//		}
//	}
//	return max
//}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
