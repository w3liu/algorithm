package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Title   string
	Content string
}

func main() {
	ch := make(chan Message)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			ch <- Message{
				Title:   fmt.Sprintf("title_%d", i),
				Content: fmt.Sprintf("content_%d", i),
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c.Title, c.Content)
	}
}
