package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayGoodBye() {
	fmt.Println("Good")
	fmt.Println("Bye")
}

func printNames() {
	students := make(map[int]string, 4)
	students[1] = "Jim"
	students[2] = "Bob"
	students[3] = "Tom"
	students[4] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}

func ExampleSayHello() {
	sayHello()
	// OutPut: Hello World
}

func ExampleSayGoodBye() {
	sayGoodBye()
	// OutPut:
	// Good
	// Bye
}

func ExamplePrintNames() {
	printNames()
	// Unordered OutPut:
	// Jim
	// Bob
	// Tom
	// Sue
}
