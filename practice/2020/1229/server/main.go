package main

import (
	"context"
	"fmt"
	foo "github.com/w3liu/algorithm/practice/2020/1229/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *foo.HelloRequest) (*foo.HelloResponse, error) {
	log.Printf("Recievied:%v", in.GetName())
	out := &foo.HelloResponse{
		Message: fmt.Sprintf("Hello %s", in.GetName()),
	}
	return out, nil
}

func main() {
	// 监听随机端口
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listen.Addr().(*net.TCPAddr).Port
	log.Println("Listen On Port:", port)
	s := grpc.NewServer()
	foo.RegisterPersonServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
