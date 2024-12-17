package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/shopspring/decimal"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	sOrder, _ := decimal.NewFromString("203.332")
	ssd, _ := decimal.NewFromString("21.3")
	u := sOrder.DivRound(ssd, 12)
	d := sOrder.DivRound(ssd, 12).RoundCeil(3)
	s := sOrder.DivRound(ssd, 12).RoundFloor(3)
	fmt.Println(sOrder, u, d, s)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://47.239.219.55:2379"}, // etcd 服务器地址
	})
	if err != nil {
		fmt.Println("Failed to create etcd client:", err)
		return
	}
	defer cli.Close()

	////cli.Put(ctx, "/xxxxKey", "/xxxxxVal", clientv3.WithLease(54), clientv3.WithLeaseTTL(int64(122)))
	//cli.Put(ctx, fmt.Sprintf("/xxxx:%d", 3), fmt.Sprintf("val:%d", 3))

	resp, err := cli.Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("Failed to get key from etcd:", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*12)
	ctx := context.Background()
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second * 5)
			cli.Delete(ctx, fmt.Sprintf("/xxxxKey:%d", i))
		}
	}()

	//go func() {
	ran := cli.Watch(context.Background(), "", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for ss := range ran {
		for _, ev := range ss.Events {
			if ev.Type == clientv3.EventTypePut {
				fmt.Printf("Key: %s, Value: %s\n", ev.Kv.Key, ev.Kv.Value)
			} else if ev.Type == clientv3.EventTypeDelete {
				fmt.Printf("Key: %s, Value: %s\n", ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
	//}()

	h := md5.New()
	bs := []byte("1")
	ns := fmt.Sprintf("%x", h.Sum(bs))

	fmt.Println(ns)

	cli.Delete(ctx, "/xxxx:3")

	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("xxxx:%d", i)
	//	fmt.Println(key)
	//	cli.Delete(ctx, key)
	//}

	//ds := cli.Ctx().Value("0.0.0.0:50053")
	//
	//fmt.Println(ds)

}
