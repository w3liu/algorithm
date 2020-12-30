package main

import (
	"context"
	"fmt"
	foo "github.com/w3liu/algorithm/practice/2020/1229/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial(":56230", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	c := foo.NewPersonServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &foo.HelloRequest{
		Name: "w3liu",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Greeting: %s \n", r.GetMessage())
}
