package main

func main() {

}

func exist(board [][]byte, word string) bool {
	l := len(board)
	w := len(board[0])

	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = byte(0)
	index := k + 1
	if dfs(board, word, i-1, j, index) || // 向左
		dfs(board, word, i+1, j, index) || // 向右
		dfs(board, word, i, j-1, index) || // 向下
		dfs(board, word, i, j+1, index) { // 向上
		return true
	}
	board[i][j] = tmp
	return false
}
