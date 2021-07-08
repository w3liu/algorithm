## Algorithm
### 1. 题目
```
48. 旋转图像
```
### 2. 题目描述
```
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
```

### 3. 解答：
```golang
func rotate(matrix [][]int) {
	var n = len(matrix)
	//上下交换
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	//对角线交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

```
### 4. 说明
* 先上下交换，再对角交换即可
* 上下交换的终止条件是 i < n/2