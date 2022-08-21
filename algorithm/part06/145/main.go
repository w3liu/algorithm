package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	predictions := dailyTemperatures1(temperatures)
	fmt.Println(predictions)
}

func dailyTemperatures(temperatures []int) []int {
	predictions := make([]int, len(temperatures))
	for x := 0; x < len(temperatures); x++ {
		v := 0
		for y := x + 1; y < len(temperatures); y++ {
			if temperatures[y] > temperatures[x] {
				v = y - x
				break
			}
		}
		predictions[x] = v
	}
	return predictions
}

func dailyTemperatures1(temperatures []int) []int {
	predictions := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		for {
			if len(stack) == 0 {
				break
			} else {
				preIndex := stack[len(stack)-1]
				if temperatures[preIndex] < temperatures[i] {
					predictions[preIndex] = i - preIndex
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, i)
	}
	return predictions
}
