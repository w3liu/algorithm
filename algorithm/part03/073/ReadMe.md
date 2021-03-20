## Algorithm
### 1. 题目
```
硬币找零问题
```
### 2. 题目描述
```
假设我们有几种不同币值的硬币 v1，v2，……，vn（单位是元）。如果我们要支付 w 元，求最少需要多少个硬币。比如，我们有 3 种不同的硬币，1 元、3 元、5 元，我们要支付 9 元，最少需要 3 个硬币（3 个 3 元的硬币）。
```

### 3. 解答：
```golang
var minVal = math.MaxInt64

func filter(coins []int,i, val, num, w int) {
	if i == len(coins) -1 || val == w {
		if val == w {
			if minVal > num {
				minVal = num
			}
		}
		return
	}
	filter(coins, i+1, val, num, w) // 不放入
	if num + coins[i+1] <= w {
		filter(coins, i+1, val + coins[i+1], num+1, w) // 放入
	}
}
```
### 4. 说明
采用回溯算法，逻辑简单，但是时间复杂度比较高