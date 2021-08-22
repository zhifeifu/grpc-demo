package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"grpc-demo/proto"
	"log"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8090", nil)
	defer clientConn.Close()
	tagServiceClient := proto.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &proto.GetTagListRequest{
		Name: "Go",
	})
	marshalResp, _ := json.Marshal(resp)
	log.Printf("resp: %v", string(marshalResp))
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
