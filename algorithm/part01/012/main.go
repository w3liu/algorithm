package main

import "fmt"

func main() {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	results := corpFlightBookings(bookings, 5)
	fmt.Println(results)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		booking := bookings[i]
		pStart := booking[0]
		pEnd := booking[1]
		if pStart > n {
			continue
		}
		if pEnd > n {
			pEnd = n
		}
		for j := pStart; j <= pEnd; j++ {
			ans[j-1] = ans[j-1] + booking[2]
		}
	}
	return ans
}
