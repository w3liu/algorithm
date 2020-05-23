package main

import "fmt"

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals := [][]int{{1, 4}, {0, 0}}
	//intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	//intervals := [][]int{{1, 4}, {0, 1}}
	//intervals := [][]int{{1, 4}, {2, 3}}
	result := merge(intervals)
	fmt.Println("result", result)
}

func merge(intervals [][]int) [][]int {
	var result = make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		a := intervals[i]
		x := a[0]
		y := a[1]
		isSelected := true
		for j := i + 1; j < len(intervals); j++ {
			isSelected = false
			b := intervals[j]
			if x <= b[0] && y >= b[0] && y <= b[1] {
				intervals[j][0] = x
				break
			} else if x >= b[0] && y >= b[1] && b[1] >= x {
				intervals[j][1] = y
				break
			} else if x >= b[0] && y <= b[1] {
				break
			} else if x <= b[0] && y >= b[1] {
				intervals[j][0] = x
				intervals[j][1] = y
				break
			} else {
				isSelected = true
			}
		}
		if isSelected {
			result = append(result, []int{x, y})
		}
	}
	return sortArr(result)
}

func sortArr(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		itemX := intervals[i]
		for j := i + 1; j < len(intervals); j++ {
			itemY := intervals[j]
			if itemX[0] > itemY[0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	return intervals
}
