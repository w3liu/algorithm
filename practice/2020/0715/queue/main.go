package main

import "fmt"

func main() {
	q := NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		q.enQueue(fmt.Sprintf("%d", i))
		fmt.Println("enQueue ", i)
	}
	for true {
		ret := q.deQueue()
		if ret == "" {
			break
		}
		fmt.Println("deQueue ", ret)
	}
}

type ArrayQueue struct {
	items []string
	n     int
	head  int
	tail  int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]string, n),
		n:     n,
		head:  0,
		tail:  0,
	}
}

// 入队
func (q *ArrayQueue) enQueue(item string) bool {
	if q.tail == q.n {
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.n; i++ {
			q.items[i-q.head] = q.items[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
	}
	q.items[q.tail] = item
	q.tail++
	return true
}

// 出队
func (q *ArrayQueue) deQueue() string {
	if q.head == q.tail {
		return ""
	}
	ret := q.items[q.head]
	q.head++
	return ret
}
