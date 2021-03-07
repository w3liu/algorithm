package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	"fmt"
	"github.com/w3liu/algorithm/books/go-advanced-program/2.6qsort_practice/qsort"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}
	qsort.Sort(unsafe.Pointer(&values[0]), len(values),
		int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)
	fmt.Println(values)
}
