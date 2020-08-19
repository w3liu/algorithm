package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name  string
	Age   int
	Phone string
}

var pool *sync.Pool

func init() {
	pool = &sync.Pool{New: func() interface{} {
		return new(User)
	}}
}

func main() {
	for i := 1; i < 1000; i++ {
		_ = NewUser(pool)
	}
	fmt.Println("over")
}

func NewUser(pool *sync.Pool) *User {
	return pool.Get().(*User)
}
