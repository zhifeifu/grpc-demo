package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-demo/proto"
	"grpc-demo/server"
	"log"
	"net"
)

var (
	port = "8090"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}

}
