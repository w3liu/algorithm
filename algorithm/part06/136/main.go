package main

import (
	"fmt"
	"math/bits"
)

func main() {
	cnt := hammingDistance(1, 4)
	fmt.Println(cnt)
}

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
