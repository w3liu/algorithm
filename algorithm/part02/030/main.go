package main

import "fmt"

func main() {
	weight := []int{2, 6, 4, 8, 2, 4, 10, 12}
	val := knapsack(weight, len(weight), 15)
	fmt.Println(val)
}

//对于一组不同重量、不可分割的物品，我们需要选择一些装入背包，在满足背包最大重量限制的前提下，背包中物品总重量的最大值是多少呢？

// weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack(weight []int, n, w int) int {
	var states = make([][]bool, n)
	for i := 0; i < len(states); i++ {
		states[i] = make([]bool, w+1)
	}
	states[0][0] = true      // 第一行的数据要特殊处理，可以利用哨兵优化
	for i := 1; i < n; i++ { // 动态规划状态转移
		for j := 0; j <= w; j++ { // 不把第i个物品放入背包
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //把第i个物品放入背包
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if states[n-1][i] == true {
			return i
		}
	}
	return 0
}
