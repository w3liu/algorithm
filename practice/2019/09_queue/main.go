package main

import (
	"errors"
	"fmt"
)

func main() {
	queue := New(10)
	for i := 0; i < 20; i++ {
		err := queue.push(i)
		if err != nil {
			fmt.Println(err)
		}
		num, err := queue.pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("num", num)
	}
	fmt.Println(queue.arr)
}

type ArrayQueue struct {
	head int
	tail int
	len  int
	arr  []int
}

func New(len int) *ArrayQueue {
	return &ArrayQueue{
		head: 0,
		tail: 0,
		len:  len,
		arr:  make([]int, len),
	}
}

func (q *ArrayQueue) push(num int) error {
	if q.len == q.tail {
		if q.head == 0 {
			return errors.New("queue is full")
		}
		q.arr = append(q.arr[q.head-1:q.len], q.arr[0:q.head]...)
		q.tail = q.len - q.head
		q.head = 0
	}
	q.arr[q.tail] = num
	q.tail++
	return nil
}

func (q *ArrayQueue) pop() (int, error) {
	if q.head == q.tail {
		return 0, errors.New("queue is nil")
	}
	result := q.arr[q.head]
	q.head++
	return result, nil
}
