package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Fatal(r)
	//	}
	//}()
	//panic(errors.New("test"))

	err := returnsError()

	if err == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(err)
	}

}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	written, err = io.Copy(dst, src)
	return
}

type MyError struct {
}

func (e MyError) Error() string {
	return "my error"
}

var ErrBad = &MyError{}

func bad() bool {
	return false
}

// 错误示例
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // 永远返回的都是一个非空error
}

// 改进示例
func returnError() error {
	if bad() {
		return ErrBad
	}
	return nil
}

// 错误示例
// 代码里的两个recover语句都不会执行到
func RunPanic() {
	if r := recover(); r != nil {
		log.Fatal(r)
	}
	panic(123)

	if r := recover(); r != nil {
		log.Fatal(r)
	}
}

// 错误示例
func RunWrapPanic() {
	defer func() {
		// 无法捕获异常
		if r := MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(1)
}

func MyRecover() interface{} {
	log.Println("trace...")
	return recover()
}

func RunPanic1() {
	// 这里可以捕获异常
	defer MyRecover()
	panic(1)
}

func RunPanic2() {
	// 无法捕获异常
	defer recover()
	panic(1)
}
