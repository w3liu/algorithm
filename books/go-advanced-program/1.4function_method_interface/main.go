package main

import "fmt"

func main() {
	//v := Inc()
	//fmt.Println("v", v)
	//ForDefer()
	ForMap()
}

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

func ForDefer() {
	for i := 0; i < 3; i++ {
		//i := i
		//defer func() { println(i) }()
		defer func(i int) { println(i) }(i)
	}
}

func ForMap() {

	strMap := make(map[int]string, 10)

	for i := 0; i < 100; i++ {
		strMap[i] = "ok"
	}

	for k, v := range strMap {
		fmt.Println(k, v)
	}
}
