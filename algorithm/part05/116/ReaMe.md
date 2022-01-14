## Algorithm
### 1. 题目
```
240. 搜索二维矩阵 II
```
### 2. 题目描述
```
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。 

示例 1：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false
```

### 3. 解答：
```golang
func searchMatrix(matrix [][]int, target int) bool {
	var xn = len(matrix)
	var yn = len(matrix[0])
	var x = 0
	var y = yn - 1

	for x < xn && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
```
### 4. 说明
Z 字形查找，从右上角开始查找，如果相等返回true，如果比target大则y--，如果比target小则x++，