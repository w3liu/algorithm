package main

import "errors"

func main() {
	if err := Foo(); err != nil {
		panic(err)
	}
}

func Foo() (err error) {
	if err := Bar(); err != nil {
		return
	}
	return
}

func Bar() error {
	return errors.New("error")
}
