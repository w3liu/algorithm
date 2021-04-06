package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8880")
	if err != nil {
		panic(err)
	}
	var conns []net.Conn
	defer func() {
		for _, c := range conns {
			_ = c.Close()
		}
	}()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(cn net.Conn) {
			handle(cn)
		}(c)
		conns = append(conns, c)
	}

}

func handle(c net.Conn) {
	defer func() {
		_ = c.Close()
	}()
	data := make([]byte, 256)
	for {
		n, err := c.Read(data)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		fmt.Println(string(data[:n]))
	}

}
