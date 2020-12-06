package main

import (
	"container/list"
)

func main() {

}

func maxDistance(grid [][]int) int {
	var dx = []int{0, 0, 1, -1}
	var dy = []int{1, -1, 0, 0}
	queue := list.New()
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue.PushBack([]int{i, j})
			}
		}
	}
	hasOcean := false
	var point []int
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		point = ele.Value.([]int)
		x := point[0]
		y := point[1]
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1
			hasOcean = true
			queue.PushBack([]int{newX, newY})
		}
	}
	if point == nil || !hasOcean {
		return -1
	}
	return grid[point[0]][point[1]] - 1
}
