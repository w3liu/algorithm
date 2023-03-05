package main

import "fmt"

func main() {
	result := intToRoman(20)
	fmt.Println(result)
}

func intToRoman(num int) string {
	var arr = []struct {
		key int
		val string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var result string
	for _, v := range arr {
		for num >= v.key {
			num = num - v.key
			result = result + v.val
		}
		if num == 0 {
			break
		}
	}
	return result
}
