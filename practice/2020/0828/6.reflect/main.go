package main

import (
	"fmt"
	"reflect"
)

func main() {
	ref4()
}

func ref1() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
}

func ref2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	var y float64 = v.Interface().(float64)

	fmt.Println("value:", y)
}

func ref3() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)
}

func ref4() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)

	fmt.Println(x)
}
