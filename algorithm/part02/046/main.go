package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 0.001
	// 0.1
	result := compareVersion("1.10.0", "1.1.0")
	fmt.Println(result)
}

func compareVersion(version1 string, version2 string) int {
	var i int
	v1 := split(version1)
	v2 := split(version2)
	for i < len(v1) && i < len(v2) {
		r := judge(v1[i], v2[i])
		if r != 0 {
			return r
		}
		i++
	}
	if i < len(v1) {
		if checkNum(v1[i:]) {
			return 1
		}
	}
	if i < len(v2) {
		if checkNum(v2[i:]) {
			return -1
		}
	}
	return 0
}

func split(s string) []string {
	return strings.Split(s, ".")
}

func judge(s1 string, s2 string) int {
	s1 = strings.TrimLeft(s1, "0")
	s2 = strings.TrimLeft(s2, "0")
	if len(s1) == 0 && len(s2) == 0 {
		return 0
	}
	if len(s1) == 0 {
		return -1
	}
	if len(s2) == 0 {
		return 1
	}
	a1, _ := strconv.Atoi(s1)
	a2, _ := strconv.Atoi(s2)
	if a1 > a2 {
		return 1
	}
	if a1 < a2 {
		return -1
	}
	return 0
}

func checkNum(nums []string) bool {
	for i := 0; i < len(nums); i++ {
		r := judge(nums[i], "")
		if r == 1 {
			return true
		}
	}
	return false
}
