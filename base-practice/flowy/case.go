package flowy

import (
	"io"
	"sync"
)

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
	if m.messageChan != nil && !m.closed {
		m.messageChan <- msg

		return nil
	} else {

		//	println("closed,buf write:" + msg.Message)
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
		close(m.messageChan)
		m.closed = true
		//	m.messageChan = nil
	}
}
