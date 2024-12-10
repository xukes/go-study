package main

import (
	"context"
	"fmt"
	pb "github.com/xukes/go-study/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	accountAddr := fmt.Sprintf("%s:%d", "127.0.0.1", 8082)
	accountConn, err := grpc.Dial(accountAddr, grpc.WithInsecure())
	if err != nil {
	}
	defer accountConn.Close()
	accountClient := pb.NewHandleServiceClient(accountConn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := pb.SendMessageRequest{
		ChatId: 34,
		Text:   "first grpc message",
	}
	res, err := accountClient.SendMessage(ctx, &req)
	if err != nil {
	}
	fmt.Println(res)
}
