package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	queue := reconstructQueue(people)
	fmt.Println(queue)
}

func reconstructQueue(people [][]int) [][]int {
	var queue = make([][]int, 0, len(people))
	sort.Slice(people, func(i, j int) bool {
		x := people[i]
		y := people[j]
		return x[0] > y[0] || x[0] == y[0] && x[1] < y[1]
	})
	for _, person := range people {
		num := person[1]
		queue = append(queue[:num], append([][]int{person}, queue[num:]...)...)
	}
	return queue
}
