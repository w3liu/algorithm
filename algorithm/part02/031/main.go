package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println("result", result)
}

func findMedianSortedArrays0(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0)
	var j = 0
	for i := 0; i < len(nums1); i++ {
		for j < len(nums2) {
			if nums1[i] <= nums2[j] {
				nums = append(nums, nums1[i])
				break
			} else {
				nums = append(nums, nums2[j])
				j++
			}
		}
		if j >= len(nums2) {
			nums = append(nums, nums1[i])
		}
	}
	if j < len(nums2) {
		nums = append(nums, nums2[j:]...)
	}
	return getMid(nums)
}

func getMid(nums []int) float64 {
	l := len(nums)
	m := float64(0)
	i := 0
	if l%2 == 1 {
		i = l / 2
		m = float64(nums[i])
	} else {
		if l == 0 {
			m = 0
		} else {
			i = l/2 - 1
			m = (float64(nums[i]) + float64(nums[i+1])) / float64(2)
		}
	}
	return m
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	iMin := 0
	iMax := m

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}
			return float64(maxLeft+minRight) / float64(2)
		}
	}
	return 0
}
