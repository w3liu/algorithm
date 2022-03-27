## Algorithm
### 1. 题目
```
322. 零钱兑换
```
### 2. 题目描述
```
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0

```

### 3. 解答：
```golang
func coinChange(coins []int, amount int) int {
	res := make([]int, amount)
	return calc(coins, amount, res)

}

func calc(coins []int, amount int, res []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if res[amount-1] != 0 {
		return res[amount-1]
	}
	var min = math.MaxInt64
	for _, coin := range coins {
		result := calc(coins, amount-coin, res)
		if result >= 0 && result < min {
			min = 1 + result
		}
	}
	if min == math.MaxInt64 {
		res[amount-1] = -1
	} else {
		res[amount-1] = min
	}
	return res[amount-1]
}

```
### 4. 说明
这道题采用的是动态规划+递归，定义res存储阶段性结果。
通过遍历coins，计算result。
临界条件是：amount < 0，amount == 0。