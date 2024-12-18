package llmflow

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type ChatRequest struct {
	Message string `json:"message"`
	Stream  bool   `json:"stream"`
}

type Usage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
	Duration         int `json:"duration"`
}

type DeltaMessage struct {
	Message      string `json:"message"`
	Usage        Usage  `json:"usage"`
	FinishReason string `json:"finishReason"`
}

type DeltaMessageWriter interface {
	Write(msg DeltaMessage) error
	Read() (DeltaMessage, error)
	Close()
	CloseWithError(err error)
}

type DeltaMessageReader interface {
	Read() (DeltaMessage, error)
}

// DeltaMessageStream 实现 DeltaMessageWriter 接口
type DeltaMessageStream struct {
	messageChan chan DeltaMessage
	locker      sync.Locker
	err         error
	closed      bool
}

type ChatResponse struct {
	ReqMessage   string `json:"reqMessage"`
	RespMessage  string `json:"respMessage"`
	Usage        Usage  `json:"usage"`
	Error        error  `json:"-"`
	FinishReason string `json:"finishReason"`
	Stream       bool   `json:"-"`
}

type ChatChannel struct {
	stream        bool
	chatResponse  *ChatResponse
	messageStream *DeltaMessageStream
	readEOFCall   func()
}

// NewStreamChatChannel 创建新的流式对话通道,并返回异步消息writer
func NewStreamChatChannel() (ChatChannel, DeltaMessageWriter) {

	// messageStream实现了DeltaMessageWriter接口
	var messageStream = &DeltaMessageStream{
		messageChan: make(chan DeltaMessage, 65535),
		// 加锁处理并发
		locker: &sync.Mutex{},
	}

	// chatResponse是一个指针
	var c = ChatChannel{
		chatResponse:  &ChatResponse{},
		messageStream: messageStream,
		stream:        true,
	}

	c.chatResponse.Stream = true

	return c, messageStream
}

// Read 获取下一条消息
func (acr *ChatChannel) Stream() bool {

	return acr.stream
}

// 读数据
func (ms *DeltaMessageStream) Read() (DeltaMessage, error) {

	if ms.messageChan == nil || ms.closed {
		return DeltaMessage{}, io.EOF
	}

	// 数据读完关闭通道
	select {

	case m, ok := <-ms.messageChan:

		if ok {
			return m, nil
		} else {
			if ms.err != nil {
				return m, ms.err
			}
			return m, io.EOF
		}

	}
}

// 写数据 (通道同步 ≠ 线程安全)
func (m *DeltaMessageStream) Write(msg DeltaMessage) error {
	m.locker.Lock()
	defer m.locker.Unlock()
	// 加互斥锁保证并发安全
	if m.messageChan != nil && !m.closed {
		m.messageChan <- msg

		return nil
	} else {
		fmt.Printf("closed,buf write:%s\n", msg.Message)
		return io.EOF
	}
}

func (m *DeltaMessageStream) Close() {

	m.locker.Lock()
	defer m.locker.Unlock()

	if m.messageChan != nil && !m.closed {
		close(m.messageChan)
		m.closed = true
		//	m.messageChan = nil
	}
}

func (m *DeltaMessageStream) CloseWithError(err error) {
	//TODO

	m.locker.Lock()
	defer m.locker.Unlock()

	if m.messageChan != nil && !m.closed {
		m.err = err
		// 流结束
		if err == io.EOF {
			log.SetPrefix("[oai-stream-end]")
			log.Println("流式输出结束")
		}
		// 关闭通道
		close(m.messageChan)
		m.closed = true
		//	m.messageChan = nil
	}
}

// Read 获取下一条消息
func (acr *ChatChannel) Read() (DeltaMessage, error) {

	if acr.stream && acr.messageStream != nil {

		var dm, err = acr.messageStream.Read()

		if err != nil {

			if err == io.EOF {

				acr.chatResponse.Usage = dm.Usage
				acr.chatResponse.RespMessage += dm.Message

				acr.messageStream.Close()
			} else {
				acr.messageStream.CloseWithError(err)
			}

			acr.readEOFCall()

		} else {

			acr.chatResponse.RespMessage += dm.Message

		}

		return dm, err

	} else {
		return DeltaMessage{}, fmt.Errorf("not a streamable channel")
	}

}

// InnerResponse 返回完整相应内容
func (acr *ChatChannel) InnerResponse() *ChatResponse {

	return acr.chatResponse
}

// setReadEOFCallback 设置读取结束回调
func (acr *ChatChannel) setReadEOFCallback(c func()) {

	acr.readEOFCall = c
}
