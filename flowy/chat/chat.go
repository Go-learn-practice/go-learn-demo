package chat

import (
	"context"
	"encoding/json"
	"flowy/llmflow"
	"log"
)

type ResponseType string

const (
	SPLASH    ResponseType = "splash"
	INCREMENT ResponseType = "increment"
	FULL      ResponseType = "full"
	FINISH    ResponseType = "finish"
)

type Usage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
	Duration         int `json:"duration"` //second
}

type ChatRequest struct {
	Ctx            context.Context `json:"-"`
	Message        string          `json:"message"`
	StreamOutputFn func(out string)
	StreamFinishFn func(ChatResponse)
}

type ChatResponse struct {
	ResponseType ResponseType         `json:"-"`
	Message      string               `json:"message"`
	Pending      bool                 `json:"pending"`
	Error        bool                 `json:"error"`
	RawResponse  llmflow.ChatResponse `json:"-"`
	Usage        Usage                `json:"usage"`
	FinishReason string               `json:"finishReason"`
}

func ChatAsyncWithAgent(request ChatRequest) <-chan ChatResponse {
	// 创建通道 缓存为1
	var result = make(chan ChatResponse, 1)

	var closeChan = func() {
		close(result)
	}

	go func() {
		// 定义一个指针类型的值
		resp := &ChatResponse{
			ResponseType: SPLASH,
			Message:      "稍等我想想~",
			Pending:      true,
		}
		result <- *resp

		var fullOutput = ""
		request.StreamOutputFn = func(out string) {
			fullOutput += out
			resp.Message = out
			resp.ResponseType = INCREMENT
			resp.Pending = true
			// 向通道写消息
			result <- *resp
		}

		request.StreamFinishFn = func(resp1 ChatResponse) {
			resp.Message = resp1.Message
			resp.ResponseType = FINISH
			resp.Pending = false

			result <- *resp
			closeChan()
			return
		}

		// 取指针的值
		*resp = ChatWithAgent(request)

		respBytes, _ := json.Marshal(*resp)

		log.SetPrefix("[chat-agent]")
		log.Println("ChatWithAgent resp = ", string(respBytes))
	}()

	return result
}
