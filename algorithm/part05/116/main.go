package main

func main() {

}

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
