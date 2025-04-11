package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Op   string `json:"op"`
	Args []Args `json:"args"`
}
type Args struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

func main() {
	// WebSocket 服务器地址
	serverURL := "wss://wsaws.okx.com:8443/ws/v5/public"

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	// 设置中断信号监听器（用于优雅退出）
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 启动一个 goroutine 来接收消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
		}
	}()

	msg := &Message{Op: "subscribe", Args: []Args{
		//{Channel: "index-tickers", InstId: "BTC-USDT"},
		//{Channel: "index-tickers", InstId: "ETH-USDT"},
		//{Channel: "index-tickers", InstId: "DOGE-USDT"},
		//{Channel: "index-tickers", InstId: "PI-USDT"},
		//{Channel: "index-tickers", InstId: "SHIB-USDT"},
		//{Channel: "index-tickers", InstId: "SOL-USDT"},
		{Channel: "price-limit", InstId: "BTC-USDT-SWAP"},
	}}

	sendMsg, _ := json.Marshal(msg)
	err = conn.WriteMessage(websocket.TextMessage, sendMsg)
	if err != nil {
		log.Println("Error writing message:", err)
		return
	}
	// 主循环：发送消息或处理退出
	for {
		select {
		case <-interrupt:
			// 发送关闭消息并退出
			log.Println("Interrupt received, closing connection...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error sending close message:", err)
			}
			return
		default:
			// 示例：每隔 5 秒发送一条消息
			//message := fmt.Sprintf("Hello, WebSocket! Time: %v", time.Now())
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}
			time.Sleep(10 * time.Second)
		}
	}
}
