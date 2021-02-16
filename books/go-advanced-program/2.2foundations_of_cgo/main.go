package main

/*
#include <stdio.h>

void printint(int v) {
	printf("printint: %d", v);
}
*/
import "C"

func main() {
	v := 42
	C.printint(C.int(v))
}
