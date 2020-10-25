package main

import (
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIp   string `key2:"value2"`
}

func main() {
	s := Server{}
	st := reflect.TypeOf(s)
	feild1 := st.Field(0)
	fmt.Printf("key1:%v\n", feild1.Tag.Get("key1"))
	fmt.Printf("key11:%v\n", feild1.Tag.Get("key11"))

	feild2 := st.Field(1)
	fmt.Printf("key2:%v\n", feild2.Tag.Get("key2"))
}
