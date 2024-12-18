package llmflow

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"strings"
)

func RawQueryStreaming(ctx context.Context, baseURL, apiKey string, msgs []openai.ChatCompletionMessage, model string, responseFormat *openai.ChatCompletionResponseFormat) (*openai.ChatCompletionStream, error) {
	var dconfig = openai.DefaultConfig(apiKey)
	dconfig.BaseURL = baseURL
	client := openai.NewClientWithConfig(dconfig)

	req := openai.ChatCompletionRequest{
		Stream:         true,
		Model:          model,
		Messages:       msgs,
		ResponseFormat: responseFormat,
	}

	resp, err := client.CreateChatCompletionStream(
		ctx,
		req,
	)
	if err != nil {
		print(err.Error())
		return nil, err
	}
	return resp, nil
}

func openAIChatStream(request ChatRequest, response *ChatResponse, msgWriter DeltaMessageWriter) {
	response.ReqMessage = request.Message
	response.Stream = true

	msgsWithPlugin := make([]openai.ChatCompletionMessage, 0)

	// 请求的userMessage
	if request.Message != "" {

		//user message
		msgsWithPlugin = append(msgsWithPlugin, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: request.Message,
		})
	}

	var apiKey = "sk-8f229cb93e78416e96430020d260f2b7"
	var model = "deepseek-chat"
	var baseURL = "https://api.deepseek.com"

	response.Stream = true
	var openaiResponseFormat = &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeText}

	var streamingResponse, err = RawQueryStreaming(context.Background(), baseURL, apiKey, msgsWithPlugin, model, openaiResponseFormat)
	response.Error = err
	if err != nil {
		log.Fatal("error")
		return
	}

	// 开启协程读取数据
	go func() {
		log.SetPrefix("[oai-stream]")
		log.Println("start stream")

		var emptyCount = 0
		for {
			// 从streamingResponse中接收数据
			rawResponse, err := streamingResponse.Recv()

			if err != nil {
				streamingResponse.Close()

				msgWriter.CloseWithError(err)

				// io.EOF 表示流已经结束
				if err != io.EOF {
					log.Fatalf("Streaming Interrupt:%s", err.Error())
				}
				return
			}

			if len(rawResponse.Choices) > 0 {
				var co = rawResponse.Choices[0]

				//coBytes, _ := json.Marshal(co)
				//fmt.Printf("响应的结果 = %s\n", string(coBytes))

				respDeltaContent := co.Delta.Content
				response.RespMessage += respDeltaContent

				if strings.TrimSpace(respDeltaContent) == "" {
					emptyCount++

					// 连续空的字符串超过10个 直接关闭通道
					if emptyCount > 10 {
						streamingResponse.Close()

						msgWriter.CloseWithError(err)

						log.Fatalf("Streaming Interrupt:%s", err.Error())
						return
					}
				} else {
					emptyCount = 0
				}

				// 写入通道
				// go-openai 没有提供消费 token 的计算
				msgWriter.Write(DeltaMessage{
					Message:      respDeltaContent,
					Usage:        Usage{},
					FinishReason: string(co.FinishReason),
				})
			}
		}
	}()
}

func OpenAIChat(request ChatRequest) ChatChannel {
	if request.Stream {
		var chatChannel, msgWriter = NewStreamChatChannel()

		// 参数2: *ChatResponse 参数3: *DeltaMessageStream的指针结构体并且实现了DeltaMessageWriter接口
		openAIChatStream(request, chatChannel.InnerResponse(), msgWriter)
		// 写入的数据会保存在 chatChannel 的 messageStream 字段中 而 messageStream 的结果是 *DeltaMessageStream
		// 因为是并发所以会直接返回 无缓存的通道的需要收发双方同时握手
		return chatChannel
	} else {
		var c = ChatChannel{
			stream: false,
		}
		return c
	}
}
