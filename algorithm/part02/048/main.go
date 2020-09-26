package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	result := numIslands(grid)
	fmt.Println("result", result)
}

func numIslands(grid [][]byte) int {
	var cnt int
	lx := len(grid)
	if lx == 0 {
		return 0
	}
	ly := len(grid[0])
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, x, y int) {
	lx := len(grid)
	ly := len(grid[0])

	grid[x][y] = '0'
	if x-1 >= 0 && grid[x-1][y] == '1' {
		dfs(grid, x-1, y)
	}
	if x+1 < lx && grid[x+1][y] == '1' {
		dfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		dfs(grid, x, y-1)
	}
	if y+1 < ly && grid[x][y+1] == '1' {
		dfs(grid, x, y+1)
	}
}
