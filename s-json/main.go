package main

import (
	"encoding/json"
	"fmt"
)

type BaseMsg struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Message struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type RespMessage struct {
	MsgId   int    `json:"msgId"`
	Message string `json:"message"`
}

func main() {
	b1 := &BaseMsg{
		Data: Message{Age: 12, Name: "xuqiwen"},
		Code: 0,
		Msg:  "success",
	}
	b2 := &BaseMsg{
		Data: RespMessage{MsgId: 102, Message: "message"},
		Code: 0,
		Msg:  "success",
	}
	bs1, _ := json.Marshal(*b1)
	fmt.Println(string(bs1))

	bs2, _ := json.Marshal(*b2)
	fmt.Println(string(bs2))

	sd := BaseMsg{Data: Message{}}
	err := json.Unmarshal(bs2, &sd)
	getName(&sd)
	if err != nil {
	}
	fmt.Println(sd)
	fmt.Println(*b2)
	fmt.Println()
}

func getName(b *BaseMsg) {
	fmt.Println("getName")
}
