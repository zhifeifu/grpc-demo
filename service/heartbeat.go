package service

import (
	"context"
	pbHeartbeat "github.com/zhifeifu/grpc-demo/proto/heartbeat"
)

type heartbeat struct {
	pbHeartbeat.UnimplementedHeartbeatServiceServer
}

func NewHeartbeat() *heartbeat {
	return &heartbeat{}
}

func (h *heartbeat) Heartbeat(ctx context.Context, request *pbHeartbeat.Request) (*pbHeartbeat.Response, error) {
	return &pbHeartbeat.Response{Value: "pong"}, nil
}
