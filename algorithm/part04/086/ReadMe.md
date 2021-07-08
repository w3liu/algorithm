## Algorithm
### 1. 题目
```
42. 接雨水
```
### 2. 题目描述
```
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

 

示例 1：
```

[!qr](./images/0620_a_1.png)

```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9
```

### 3. 解答：
```golang
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	var ans int
	var lMax = make([]int, len(height))
	var rMax = make([]int, len(height))
	lMax[0], rMax[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height)-1; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	for i := len(height) - 2; i > 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		ans = ans + min(lMax[i], rMax[i]) - height[i]
	}
	return ans
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
```
### 4. 说明
* 定义两个数组：lMax、rMax
* lMax表示从左往右，记录当前位置的最大高度
* rMax表示从右往左，记录当前位置的最大高度
* 最后求结果很巧妙，用当前位置lMax与rMax的最小值减去当前位置的高度并累加到最后