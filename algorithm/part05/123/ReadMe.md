## Algorithm
### 1. 题目
```
309. 最佳买卖股票时机含冷冻期
```
### 2. 题目描述
```
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 

示例 1:

输入: prices = [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
示例 2:

输入: prices = [1]
输出: 0
```

### 3. 解答：
```golang
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// f[i][0]: maximum yield from holding stocks
	// f[i][1]: do not own stock, and in the frozen period of the cumulative maximum yield
	// f[i][2]: do not own stock, and not in the freeze period cumulative maximum yield
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```
### 4. 说明
```
有三种状态：
我们目前持有一支股票，对应的「累计最大收益」记为 f[i][0]；
我们目前不持有任何股票，并且处于冷冻期中，对应的「累计最大收益」记为 f[i][1]；
我们目前不持有任何股票，并且不处于冷冻期中，对应的「累计最大收益」记为 f[i][2]。

状态转移方程：
如果当前持有股票，则有可能是之前持有的，这次没有买，也有可能是这次买的。如果是这次买的，则收益为f[i][2]-prices[i]。
f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])

如果当前不持有股票，且处于冷冻期，那么说明在头一天卖出了股票，那么收益应该是f[i-1][0]+prices[i]。
f[i][1] = f[i-1][0] + prices[i]

如果当前不持有股票，且不处于冷冻期，则可能头一天处于冷冻期或者不处于冷冻期。
f[i][2] = max(f[i-1][1], f[i-1][2])
```