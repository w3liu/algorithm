package main

import "fmt"

func main() {
	var a [3]int
	var b = [...]int{1, 2, 3}
	var c = [...]int{2: 3, 1: 2}
	var d = [...]int{1, 2, 4: 5, 6}
	fmt.Println(a, b, c, d)
}
