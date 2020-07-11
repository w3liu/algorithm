package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	//signal.Notify(c, os.Interrupt)
	signal.Notify(c)
	s := <-c
	fmt.Println(s)
}
