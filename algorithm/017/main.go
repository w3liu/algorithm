package main

import "fmt"

func main() {
	// [['O','X','O','O','X','X'],['O','X','X','X','O','X'],['X','O','O','X','O','O'],['X','O','X','X','X','X'],['O','O','X','O','X','X'],['X','X','O','O','O','O']]
	// [['X','X','X'],['X','O','X'],['X','X','X']]
	board := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X', 'X'},
	}
	//board := [][]byte{
	//	{'X','X','X'},
	//	{'X','O','X'},
	//	{'X','X','X'},
	//}
	solve(board)
	fmt.Println(board)
}

var BD [][]byte

func solve(board [][]byte) {
	BD = board
	r := len(board)
	if r == 0 {
		return
	}
	c := len(board[0])
	if r < 3 || c < 3 {
		return
	}

	for i := 0; i < r; i++ {
		dfs(i, 0)
		dfs(i, c-1)
	}

	for i := 0; i < c; i++ {
		dfs(0, i)
		dfs(r-1, i)
	}

	for ri, row := range board {
		for ci, _ := range row {
			if board[ri][ci] == 'O' {
				board[ri][ci] = 'X'
			}
			if board[ri][ci] == '#' {
				board[ri][ci] = 'O'
			}
		}
	}
}

func dfs(i, j int) {
	r, c := len(BD), len(BD[0])
	if i < 0 || j < 0 || i >= r || j >= c || BD[i][j] != 'O' {
		return
	}
	BD[i][j] = '#'
	dfs(i-1, j)
	dfs(i+1, j)
	dfs(i, j-1)
	dfs(i, j+1)
}
