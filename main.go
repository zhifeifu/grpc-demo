package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbHeartbeat "github.com/zhifeifu/grpc-demo/proto/heartbeat"
	"github.com/zhifeifu/grpc-demo/service"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	pbHeartbeat.RegisterHeartbeatServiceServer(s, service.NewHeartbeat())
	go func() {
		s.Serve(listen)
	}()
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return
	}
	mux := runtime.NewServeMux()
	err = pbHeartbeat.RegisterHeartbeatServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		return
	}
}
