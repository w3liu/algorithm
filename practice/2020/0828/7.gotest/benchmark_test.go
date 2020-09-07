package main

import "testing"

func makeSliceWithoutAlloc() []int {
	var newSlice []int
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

func makeSliceWithAlloc() []int {
	var newSlice []int
	newSlice = make([]int, 9, 100000)
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSliceWithAlloc()
	}
}
