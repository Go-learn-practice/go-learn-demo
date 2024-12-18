package llmflow

import (
	"log"
	"log/slog"
)

type FlowyInitConfig struct {

	//logger
	//gpt key
	//memory implement
	Logger  *slog.Logger
	Session string

	ChatCallback func(ChatRequest, ChatResponse)
}

type Flowy struct {
	logger       *slog.Logger
	session      string
	chatCallback func(ChatRequest, ChatResponse)
}

func NewFlowyChat(initConfig *FlowyInitConfig) *Flowy {

	var logger *slog.Logger

	var chatCallback func(ChatRequest, ChatResponse)

	if initConfig != nil {
		logger = initConfig.Logger
		chatCallback = initConfig.ChatCallback
	}

	if logger == nil {
		logger = slog.Default()
	}

	var f = &Flowy{
		logger:       logger,
		chatCallback: chatCallback,
	}

	return f
}

func (flow *Flowy) DoChat(request ChatRequest) (response ChatChannel, err error) {
	// 日志加前缀
	log.SetPrefix("[DoChat-flowy]")
	// 请求的消息
	log.Println("message: ", request.Message)

	// 发送请求 结果：ChatChannel
	response = OpenAIChat(request)

	if response.Stream() {
		response.setReadEOFCallback(func() {
			//var frPtr = response.InnerResponse()
			//处理响应结果
		})
	} else {
	}

	return
}
