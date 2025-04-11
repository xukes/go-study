package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// SSE服务器URL
	url := "https://hermes.pyth.network/v2/updates/price/stream?ids[]=0x7d669ddcdd23d9ef1fa9a9cc022ba055ec900e91c4cb960f3c20429d4447a411"

	// 监听SSE事件
	err := listenSSE(url)
	if err != nil {
		log.Fatalf("Error listening to SSE: %v", err)
	}
}

// listenSSE 监听SSE事件流
func listenSSE(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	// 设置Accept头为text/event-stream
	req.Header.Set("Accept", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// 创建一个reader来读取事件流
	reader := bufio.NewReader(resp.Body)
	for {
		// 读取直到遇到空行(一个完整的事件)
		event, err := readEvent(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Connection closed by server")
				return nil
			}
			return fmt.Errorf("error reading event: %v", err)
		}
		// 处理事件
		if len(event) > 0 {
			fmt.Printf("Received event: %s\n", event)
		}
	}
}

func readEvent(reader *bufio.Reader) (string, error) {
	var buffer bytes.Buffer
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return "", err
		}
		// 检查是否是空行(事件结束)
		if len(line) <= 2 { // \r\n 或 \n
			break
		}
		// 写入缓冲区
		buffer.Write(line)
	}
	return buffer.String(), nil
}
