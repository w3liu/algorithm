package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)

	for i := 0; i < 5; i++ {
		fmt.Println(p.Get())
	}

}

func TestLower(t *testing.T) {
	str := "0x0F508F143E77b39F8e20DD9d2C1e515f0f527D9F"
	fmt.Println(strings.ToLower(str))
}
