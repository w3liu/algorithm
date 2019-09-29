package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	result := reverse(-123400000)
	fmt.Println("result", result)
}

func reverse(x int) int {
	var result int
	s := fmt.Sprintf("%d", x)
	prefix := ""
	if strings.Index(s, "-") >= 0 {
		prefix = "-"
		s = strings.TrimLeft(s, "-")
	}
	arr := []rune(s)
	temp := []rune("")
	zero := 0
	for i := len(arr) - 1; i >= 0; i-- {
		item := arr[i]
		if zero == 0 && string(item) == "0" {
			continue
		}
		temp = append(temp, item)
		zero++
	}
	y := string(temp)
	if prefix == "-" {
		y = fmt.Sprintf("-%s", y)
	}
	result, _ = strconv.Atoi(y)
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}

func reverse1(x int) int {
	var res int64
	for x != 0 {
		res = res*10 + int64(x)%10
		x = x / 10
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return int(res)
}
