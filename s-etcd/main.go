package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://47.239.219.55:2379"}, // etcd 服务器地址
	})
	if err != nil {
		fmt.Println("Failed to create etcd client:", err)
		return
	}
	defer cli.Close()
	resp, err := cli.Get(context.Background(), "/", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("Failed to get key from etcd:", err)
		return
	}
	h := md5.New()

	bs := []byte("1")

	ns := fmt.Sprintf("%x", h.Sum(bs))

	fmt.Println(ns)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	//cli.Put(ctx, "/xxxxKey", "/xxxxxVal", clientv3.WithLease(54), clientv3.WithLeaseTTL(int64(122)))

	cli.Delete(ctx, "/xxxxKey")

	ds := cli.Ctx().Value("0.0.0.0:50053")

	fmt.Println(ds)
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
}
