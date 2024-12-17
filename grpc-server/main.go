package main

import (
	"context"
	"fmt"
	pb "github.com/xukes/go-study/proto"
	"google.golang.org/grpc"
	"net"
)

type MyPerson interface {
	Less()
}
type MyPerson1 struct {
}
type MyPerson2 struct {
	Age  int
	Name string
}

func (m MyPerson2) Less() bool {
	return m.Age < 10
}

func (m *MyPerson1) Less() {
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type HandleServiceServer struct {
	pb.UnimplementedHandleServiceServer
}

func main() {
	addr := ":8082"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHandleServiceServer(s, &HandleServiceServer{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
func (HandleServiceServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (resp *pb.SendMessageResponse, err error) {
	resp = &pb.SendMessageResponse{
		Success:   true,
		MessageId: "msg success",
	}
	fmt.Println(req.Text)
	return resp, nil
}

func (HandleServiceServer) GetMessage(ctx context.Context, req *pb.GetMessageRps) (resp *pb.GetMessageResp, err error) {
	fmt.Println(req.BaseMsg.Msg)
	return &pb.GetMessageResp{ChatId: 23, Text: "msg"}, nil
}
