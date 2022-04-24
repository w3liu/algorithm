package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := decodeString("2[abc]3[cd]ef")
	fmt.Println(str)
}

type strStack struct {
	arr []string
}

func newStrStack() *strStack {
	return &strStack{arr: make([]string, 0)}
}

func (s *strStack) push(c string) {
	s.arr = append(s.arr, c)
}

func (s *strStack) pop() string {
	if len(s.arr) > 0 {
		c := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return c
	}
	return ""
}

type intStack struct {
	arr []int
}

func newIntStack() *intStack {
	return &intStack{arr: make([]int, 0)}
}

func (s *intStack) push(c int) {
	s.arr = append(s.arr, c)
}

func (s *intStack) pop() int {
	if len(s.arr) > 0 {
		c := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return c
	}
	return -1
}

func decodeString(s string) string {
	var res string
	strs := newStrStack()
	ints := newIntStack()
	var multi int
	for _, c := range s {
		if c == '[' {
			strs.push(res)
			ints.push(multi)
			multi = 0
			res = ""
		} else if c == ']' {
			var tmp string
			currMulti := ints.pop()
			for i := 0; i < currMulti; i++ {
				tmp = fmt.Sprintf("%s%s", res, tmp)
			}
			res = strs.pop() + tmp
		} else if c >= '0' && c <= '9' {
			num, _ := strconv.ParseInt(string(c), 10, 64)
			multi = 10*multi + int(num)
		} else {
			res = res + string(c)
		}
	}
	return res
}
