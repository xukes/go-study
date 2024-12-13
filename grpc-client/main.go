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
	var req1 *pb.GetMessageResp
	var err2 error

	ch := make(chan int)
	defer func() {
		close(ch)
	}()

	go func() {
		startT := time.Now().UnixMilli()
		_ = <-ch
		fmt.Println(time.Now().UnixMilli() - startT)
	}()
	for i := 0; i < 1000; i++ {
		req1, err2 = accountClient.GetMessage(ctx, &pb.GetMessageRps{BaseMsg: &pb.BaseMsg{Msg: fmt.Sprintf(" index:%d", i)}})
		if err2 != nil {
		}
	}
	ch <- 2
	fmt.Println(req1)
}
