package main

func main() {

}

func rotate1(matrix [][]int) {
	var n = len(matrix)
	var temp = make([][]int, n)
	for i, _ := range temp {
		temp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			temp[j][n-1-i] = v
		}
	}
	copy(matrix, temp)
}

func rotate2(matrix [][]int) {
	var n = len(matrix)
	// 对称图形，只需要旋转前半行即可
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			x := n - 1 - j
			y := n - 1 - i
			matrix[i][j] = matrix[x][i]
			matrix[x][i] = matrix[y][x]
			matrix[y][x] = matrix[j][y]
			matrix[j][y] = temp
		}
	}
}

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
