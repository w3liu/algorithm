package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:8880")
	if err != nil {
		panic(err)
	}
	//ticker := time.NewTicker(time.Millisecond * 5)
	for {
		_, err := c.Write([]byte("hello world"))
		if err != nil {
			log.Println(err)
		}
	}
}
