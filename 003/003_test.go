package main

import "testing"

const (
	Read = 1 << iota
	Write
	Execute
)

func TestConstTry(t *testing.T) {
	t.Log(Read, Write, Execute)
}
