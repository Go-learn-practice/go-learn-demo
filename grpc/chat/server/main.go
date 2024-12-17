package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/chat/chat"
	"io"
	"log"
	"net"
	"time"
)

/*
	flag 包是 Go 标准库提供的命令行参数解析工具，用于解析命令行参数
*/

var (
	port = flag.Int("port", 50051, "The server port")
)

// ChatServiceServer 定义 ChatService 服务的实现
type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServiceServer) ChatStream(stream pb.ChatService_ChatStreamServer) error {
	log.Println("ChatStream started")

	// 循环读取客户端发送的消息并回复
	for {
		// 从流中接收消息
		req, err := stream.Recv()
		if err == io.EOF {
			// 流关闭
			log.Println("Client closed the stream")
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
			return err
		}
		// 打印客户端的消息
		log.Printf("Received message from %s: %s", req.User, req.Message)

		// 服务端回复消息
		resp := &pb.ChatMessage{
			User:      "Server",
			Message:   fmt.Sprintf("Hello %s, I received your message: '%s'", req.User, req.Message),
			Timestamp: time.Now().Unix(),
		}

		// 发送消息给客户端
		if err := stream.Send(resp); err != nil {
			log.Fatalf("Error sending message: %v", err)
			return err
		}
	}
}

func main() {
	// 解析命令行参数
	flag.Parsed()

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
