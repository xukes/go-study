package main

import (
	"fmt"
)

type Inter struct {
	sum int
}

/*
协程中使用通道输出结果
*/
func main() {
	ch := make(chan Inter)
	go celSum(2, 4, ch)
	it := <-ch
	fmt.Println(it.sum)

	// 通道的方向
	var send_only chan<- int
	var recv_only <-chan int
	go func() {
		send_only <- 3
	}()
	go func() {
		fmt.Println(<-recv_only)
	}()
}

func celSum(x int, y int, c chan<- Inter) {
	c <- Inter{x + y}
}
