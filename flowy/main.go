package main

import (
	"flowy/chat"
	"fmt"
	"log"
)

func main() {
	var params chat.ChatRequest
	params.Message = "你好，你能帮我做什么"

	// 定义只读通道 接收回复的消息
	var respChan <-chan chat.ChatResponse

	// 处理请求
	respChan = chat.ChatAsyncWithAgent(params)

	// 从通道中接收消息 直到通道关闭
	for resp := range respChan {

		if !resp.Pending {
			log.Printf("执行结束：%v\n", resp.Pending)
			fmt.Printf("结束消息：%s\n", resp.Message)
		} else {
			switch resp.ResponseType {
			case chat.SPLASH:
				fmt.Printf("启动消息：%s\n", resp.Message)
			case chat.INCREMENT:
				fmt.Printf("增量消息：%s\n", resp.Message)
			case chat.FULL:
				fmt.Printf("完整消息：%s\n", resp.Message)
			case chat.FINISH:
				fmt.Printf("结束消息：%s\n", resp.Message)

			}
		}
	}
}
