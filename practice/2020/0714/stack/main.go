package main

import "fmt"

// 栈

func main() {
	stack := NewArrayStack(10)
	for i := 0; i < 10; i++ {
		fmt.Println("push ", i)
		stack.push(fmt.Sprintf("%d", i))
	}
	for true {
		val := stack.pop()
		fmt.Println("pop ", val)
		if val == "" {
			break
		}
	}
}

type ArrayStack struct {
	items []string
	count int
	n     int
}

func NewArrayStack(n int) *ArrayStack {
	items := make([]string, n)
	return &ArrayStack{
		items: items,
		count: 0,
		n:     n,
	}
}

// 入栈操作
func (s *ArrayStack) push(item string) bool {
	if s.count == s.n {
		return false
	}
	s.items[s.count] = item
	s.count++
	return true
}

// 出栈操作
func (s *ArrayStack) pop() string {
	if s.count == 0 {
		return ""
	}
	temp := s.items[s.count-1]
	s.count--
	return temp
}
