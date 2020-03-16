package main

import "fmt"

func main() {
	nums1 := []int{}
	nums2 := []int{2, 3}
	result := findMedianSortedArrays(nums2, nums1)
	fmt.Println("result", result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	arr := make([]int, 0)
	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		arr = nums2
	} else if l2 == 0 {
		arr = nums1
	} else {
		var k int
		for _, i1 := range nums1 {
			for k < len(nums2) {
				i2 := nums2[k]
				if i1 <= i2 {
					arr = append(arr, i1)
					break
				} else {
					arr = append(arr, i2)
					k++
				}
			}
			if k == len(nums2) {
				arr = append(arr, i1)
			}
		}
		if k < len(nums2) {
			for k < len(nums2) {
				arr = append(arr, nums2[k])
				k++
			}
		}
	}
	l := len(arr)
	if l == 1 {
		return float64(arr[0])
	} else {
		if l%2 == 0 {
			return float64(arr[l/2-1]+arr[l/2]) / float64(2)
		} else {
			return float64(arr[l/2])
		}
	}
}
