package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		MakeFoo()
	}
	for i := 0; i < 100; i++ {
		stats := debug.GCStats{}
		debug.ReadGCStats(&stats)
		if len(stats.Pause) > 0 {
			fmt.Println("numGC", stats.NumGC, "lastGC", stats.LastGC, "Pause", stats.Pause[0])
		}
		time.Sleep(time.Second * 5)
	}

	time.Sleep(time.Hour)
}

type Foo struct {
	table map[int][]byte
}

func NewFoo(size int) *Foo {
	tb := make(map[int][]byte)
	for i := 0; i < size; i++ {
		k := 1024 * 1024
		v := make([]byte, k)
		for j := 0; j < k; j++ {
			v[j] = 'a'
		}
		tb[i] = v
	}
	return &Foo{table: tb}
}

func Len(foo *Foo) int {
	return len(foo.table)
}

func MakeFoo() {
	foo := NewFoo(1024)
	l := Len(foo)
	fmt.Println("l", l)
	for k, v := range foo.table {
		fmt.Println(k, len(v))
	}
	foo.table = nil
}
