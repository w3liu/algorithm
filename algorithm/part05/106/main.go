package main

import "math"

func main() {
	var nums = []int{7, 1, 5, 3, 6, 4}
	var result = maxProfit(nums)
	println(result)
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

func maxProfit(prices []int) int {
	var minPrice, maxProfit = math.MaxInt64, 0
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
