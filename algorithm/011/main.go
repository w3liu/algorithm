package main

import "fmt"

func main() {
	books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	h := minHeightShelves(books, 4)
	fmt.Println(h)
}

func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = 1000000
		tw, j, h := 0, i, 0
		for j > 0 {
			tw = tw + books[j-1][0]
			if tw > shelf_width {
				break
			}
			if h < books[j-1][1] {
				h = books[j-1][1]
			}
			if dp[i] > dp[j-1]+h {
				dp[i] = dp[j-1] + h
			}
			j--
		}
	}
	return dp[n]
}
