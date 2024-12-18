package chat

import (
	"bytes"
	"encoding/json"
	"flowy/llmflow"
	"io"
	"log"
)

func ChatWithAgent(chatRequest ChatRequest) (respMessage ChatResponse) {
	log.Println("开始对话")

	// 获取结构体地址 *llmflow.Flowy
	var flowyClient = llmflow.NewFlowyChat(&llmflow.FlowyInitConfig{
		ChatCallback: nil,
	})
	var message = chatRequest.Message

	var llmRequest = llmflow.ChatRequest{
		Message: message,
		Stream:  true,
	}

	//执行大模型对话 获取chatChannel结果
	chatChannel, err := flowyClient.DoChat(llmRequest)

	if err != nil {
		respMessage.Error = true
		respMessage.RawResponse.Error = err
		respMessage.Pending = false

		log.Fatalf("flowyClient.DoChat Error = %s", err.Error())
		return respMessage
	}

	// 因为 oai 开启并发处理数据 这里会立即响应
	if chatChannel.Stream() {
		if chatChannel.InnerResponse().Error != nil {

			respMessage.Error = true
			respMessage.RawResponse.Error = err

			return respMessage
		}

		log.Println("开始读取数据，进入时数据应该为空")
		respMessage.RawResponse = *chatChannel.InnerResponse()
		// 开启一个协程
		go func() {
			log.Println("开启协程并发读取数据")

			var fullResp = bytes.Buffer{}

			if respMessage.RawResponse.Error != nil {
				goto failed
			}

			for {
				// 读取数据 通道无缓存 需要接收双方同时握手
				var dmsg, err = chatChannel.Read()

				if err == nil {

					if chatRequest.StreamOutputFn != nil {
						chatRequest.StreamOutputFn(dmsg.Message)
					}
					respMessage.RawResponse.Usage = dmsg.Usage
					respMessage.FinishReason = dmsg.FinishReason
					fullResp.WriteString(dmsg.Message)
				} else {

					if err == io.EOF {
						//正常结束
						break
					} else {

						respMessage.RawResponse.Error = err
						break
					}
				}
			}

			respMessage.RawResponse = *chatChannel.InnerResponse()
		failed:
			respMessage.Usage.PromptTokens = respMessage.RawResponse.Usage.PromptTokens
			respMessage.Usage.CompletionTokens = respMessage.RawResponse.Usage.CompletionTokens
			respMessage.Usage.TotalTokens = respMessage.RawResponse.Usage.TotalTokens

			respMessage.Pending = false

			if respMessage.RawResponse.Error != nil {
				respMessage.Error = true
				respMessage.Message = respMessage.RawResponse.Error.Error()
			} else {
				respMessage.Error = false
				respMessage.RawResponse.RespMessage = fullResp.String()
				respMessage.Message = fullResp.String()

				fullResp.Reset()
			}

			// 结束返回完整消息
			if chatRequest.StreamFinishFn != nil {
				chatRequest.StreamFinishFn(respMessage)
			}
		}()
	} else {
	}

	respBytes, _ := json.Marshal(respMessage)
	// 过去响应的结果
	log.Println("数据响应", string(respBytes))
	return
}
