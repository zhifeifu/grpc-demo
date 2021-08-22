package server

import (
	"context"
	"fmt"
	"grpc-demo/proto"
)

type TagServer struct {
}

func (t TagServer) GetTagList(ctx context.Context, request *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	//api := bapi.NewAPI("http://127.0.0.1:8000")
	//list, err := api.GetTagList(ctx, request.GetName())
	//if err != nil {
	//	return nil, err
	//}
	//tagList := proto.GetTagListReply{}
	//err = json.Unmarshal(list, &tagList)
	//if err != nil {
	//	return nil, errcode.TogRPCError(errcode.Fail)
	//}
	fmt.Println(request.GetName())
	tagList := proto.GetTagListReply{
		List: []*proto.Tag{
			{
				Id:    1,
				Name:  "a",
				State: 1,
			},
			{
				Id:    2,
				Name:  "b",
				State: 1,
			},
			{
				Id:    3,
				Name:  "c",
				State: 1,
			},
		},
		Pager: &proto.Pager{
			Page:      1,
			PageSize:  10,
			TotalRows: 100,
		},
	}
	return &tagList, nil
}

func NewTagServer() *TagServer {
	return &TagServer{}
}
