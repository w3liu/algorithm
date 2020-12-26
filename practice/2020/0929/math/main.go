package main

import (
	"fmt"
	"math"
)

func main() {
	result := math.Pow(2, 100)

	nan := 0x7FF8000000000001
	b := math.IsNaN(math.Float64frombits(uint64(nan)))

	fmt.Println(fmt.Sprintf("%#v", math.Float64frombits(uint64(nan))))

	fmt.Println("result", result)
	fmt.Println("b", b)
}
