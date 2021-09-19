## Algorithm
### 1. 题目
```
85. 最大矩形
```
### 2. 题目描述
```
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：
```
[!qr](./images/0913_a_1.jpg)
```

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
```

### 3. 解答：
```golang
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var result int
	var heights = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[j] = heights[j] + 1
			} else {
				heights[j] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > result {
			result = area
		}
	}
	return result
}

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
1. 上周做过求`柱状图最大面积`的一道题；
2. 本题可以将问题拆分为n个`柱状图最大面积`的子问题题；
3. 定义heights数组，然后遍历每一层，将heights传入`largestRectangleArea`方法，求得最大面积；
4. 与`result`结果比较，最后返回result即可；
