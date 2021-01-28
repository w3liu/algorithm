package main

func main() {
	routine2()
}

func routine1() {
	var a string
	var done bool
	var setup = func() {
		a = "hello world"
		done = true
	}
	go setup()
	for !done {
	}
	print(a)
}

func routine2() {
	var done = make(chan bool)
	var msg string
	a := func() {
		msg = "你好，世界"
		done <- true
	}

	go a()

	<-done

	print(msg)
}
