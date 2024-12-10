package main

import (
	"context"
	"fmt"
	pb "github.com/xukes/go-study/proto"
	"google.golang.org/grpc"
	"net"
)

type HandleServiceServer struct {
	pb.UnimplementedHandleServiceServer
}

var u = HandleServiceServer{}

func main() {
	addr := "127.0.0.1:8082"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	pb.RegisterHandleServiceServer(s, &u)
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (p HandleServiceServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (resp *pb.SendMessageResponse, err error) {
	resp = &pb.SendMessageResponse{
		Success:   true,
		MessageId: "msg success",
	}
	fmt.Println(req.Text)
	return resp, nil
}
