package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	ret := stack.GetMin()
	fmt.Println(ret)
	stack.Pop()
	ret = stack.Top()
	fmt.Println(ret)
	ret = stack.GetMin()
	fmt.Println(ret)
}

type MinStack struct {
	min int
	arr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: 0,
		arr: make([]int, 0),
	}
}

func (s *MinStack) Push(x int) {
	if len(s.arr) == 0 {
		s.arr = append(s.arr, 0)
		s.min = x
	} else {
		diff := int64(x) - int64(s.min)
		s.arr = append(s.arr, int(diff))
		if s.min > x {
			s.min = x
		}
	}

}

func (s *MinStack) Pop() {
	if len(s.arr) == 0 {
		return
	}
	diff := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	if diff < 0 {
		s.min = s.min - diff
	}
}

func (s *MinStack) Top() int {
	if len(s.arr) == 0 {
		return 0
	}
	diff := s.arr[len(s.arr)-1]
	if diff >= 0 {
		return s.min + diff
	} else {
		return s.min
	}
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
