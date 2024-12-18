package main

import (
	"flowy/chat"
	"fmt"
	"log"
	"sync"
)

// 并发启动多个 ChatAsyncWithAgent 请求并返回独立通道
func LoopChatAgent(reqs []chat.ChatRequest) []<-chan chat.ChatResponse {
	numReqs := len(reqs)

	// 用于存储每个请求的响应通道
	responseChans := make([]<-chan chat.ChatResponse, numReqs)

	// 使用 WaitGroup 控制请求的启动和响应
	var wg sync.WaitGroup

	for i, req := range reqs {
		wg.Add(1)
		// 闭包
		go func(idx int, req chat.ChatRequest) {
			defer wg.Done()
			// 启动请求并存储响应通道
			responseChans[idx] = chat.ChatAsyncWithAgent(req)
		}(i, req)
	}

	// 确保所有请求已启动
	wg.Wait()

	return responseChans
}

// 遍历切片通道获取结果
func LoopGetResponse(requests []chat.ChatRequest) {
	// 模拟并发
	respChans := LoopChatAgent(requests)

	// 并发处理每个响应
	var wg sync.WaitGroup
	for i, respChan := range respChans {
		wg.Add(1)
		go func(idx int, respChan <-chan chat.ChatResponse) {
			defer wg.Done()

			fmt.Printf("开始处理用户请求 #%d\n", idx+1)
			for resp := range respChan {
				if !resp.Pending {
					log.Printf("用户请求 #%d 执行结束：%v\n", idx+1, resp.Pending)
					fmt.Printf("用户请求 #%d 结束消息：%s\n", idx+1, resp.Message)
				} else {
					switch resp.ResponseType {
					case chat.SPLASH:
						fmt.Printf("用户请求 #%d 启动消息：%s\n", idx+1, resp.Message)
					case chat.INCREMENT:
						fmt.Printf("用户请求 #%d 增量消息：%s\n", idx+1, resp.Message)
					case chat.FULL:
						fmt.Printf("用户请求 #%d 完整消息：%s\n", idx+1, resp.Message)
					case chat.FINISH:
						fmt.Printf("用户请求 #%d 结束消息：%s\n", idx+1, resp.Message)
					}
				}
			}
		}(i, respChan)
	}
	wg.Wait()
	log.SetPrefix("[main]")
	log.Println("所有请求已完成")
}

func ReqParams() {
	requests := []chat.ChatRequest{
		{Message: "你好"},
		{Message: "你是谁"},
		{Message: "你好，你能帮我做什么"},
		{Message: "如何博取心上人一笑"},
	}
	LoopGetResponse(requests)
}

// single 请求
func SingleChat() {
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

func main() {
	SingleChat()

	// 并发
	//ReqParams()
}
