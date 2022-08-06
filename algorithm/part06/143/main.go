package main

import "math"

func main() {

}

func leastInterval(tasks []byte, n int) (minTime int) {
	stat := map[byte]int{}
	for _, t := range tasks {
		stat[t]++
	}

	nextValid := make([]int, 0, len(stat))
	rest := make([]int, 0, len(stat))
	for _, v := range stat {
		nextValid = append(nextValid, 1)
		rest = append(rest, v)
	}

	for range tasks {
		minTime++
		minNextValid := math.MaxInt64
		for i, v := range rest {
			if v > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return
}
