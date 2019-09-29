package main

import (
	"fmt"
	"math/bits"
)

func main() {
	res := pathInZigZagTree(100)
	fmt.Println(res)
}

func pathInZigZagTree(label int) []int {
	// 偶数行从大到小 每一行的数字 2的行-1次方-2的行次方-1，
	// 奇数行从小到大
	// 1000
	res := make([]int, 0)

	h := bits.Len(uint(label))

	for i := h; i > 0; i-- {
		res = append(res, label)
		max := 1<<uint(i-1) - 1
		// min := max + 1>>1
		// 偶数行
		m := label % (1 << uint(i-1)) / 2
		if label%2 == 1 {
			m = (label - 1) % (1 << uint(i-1)) / 2
		}
		label = max - m
	}
	arr := make([]int, h)
	for i := 0; i < h/2; i++ {
		arr[i] = res[h-i-1]
		arr[h-i-1] = res[i]
	}
	if h%2 == 1 {
		arr[h/2] = res[h/2]
	}

	return arr
}
