## Algorithm
### 1. 题目
```
84. 柱状图中最大的矩形
```
### 2. 题目描述
```
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
```
示例 1:
[!qr](./images/0906_a_1.jpg)
```
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
```
[!qr](./images/0906_a_2.jpg)
```
输入： heights = [2,4]
输出： 4
```

### 3. 解答：
```golang
func largestRectangleArea(heights []int) int {
	var n = len(heights)
	var res = 0
	var stack = make([]int, 0)
	var temp = make([]int, n+2, n+2)
	for i := 0; i < n; i++ {
		temp[i+1] = heights[i]
	}
	heights = temp
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 {
			tail := stack[len(stack)-1]
			if heights[i] < heights[tail] {
				stack = stack[:len(stack)-1]
				curI := tail
				curH := heights[curI]
				if len(stack) > 0 {
					tail = stack[len(stack)-1]
					leftI := tail
					rightI := i
					curW := rightI - leftI - 1
					curS := curW * curH
					if curS > res {
						res = curS
					}
				}
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}
```
### 4. 说明
1. 用数组模拟栈；
2. 栈内元素是单调递增的；
3. 出栈计算面积，并判断是否大于当前结果；