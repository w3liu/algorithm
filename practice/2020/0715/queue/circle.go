package main

type CircleQueue struct {
	items []string
	n     int
	head  int
	tail  int
}

func NewCircleQueue(n int) *CircleQueue {
	return &CircleQueue{
		items: make([]string, n),
		n:     n,
		head:  0,
		tail:  0,
	}
}

// 入队
func (q *CircleQueue) enQueue(item string) bool {
	if (q.tail+1)%q.n == q.head {
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.n
	return true
}

// 出队
func (q *CircleQueue) deQueue() string {
	if q.head == q.tail {
		return ""
	}
	ret := q.items[q.head]
	q.head = (q.head + 1) % q.n
	return ret
}
