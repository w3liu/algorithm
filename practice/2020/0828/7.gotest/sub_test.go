package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func main() {
	fmt.Println("hello test main")
}

func add(a, b int) int {
	return a + b
}

func sub1(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
	time.Sleep(time.Second)
}

func sub2(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

func TestSub(t *testing.T) {
	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
}

func parallelSub1(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func parallelSub2(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestSubParallel(t *testing.T) {
	t.Logf("Setup")

	t.Run("group", func(t *testing.T) {
		t.Run("Test1", parallelSub1)
		t.Run("Test2", parallelSub2)
	})

	t.Logf("Tear Down")
}

func TestMain(m *testing.M) {
	println("TestMain setup.")
	retCode := m.Run()
	println("TestMain tear-down.")
	os.Exit(retCode)
}
