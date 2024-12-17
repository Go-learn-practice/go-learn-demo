package main

import (
	"bufio"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/chat/chat"
	"io"
	"log"
	"os"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parsed()
	// 连接到 gRPC 服务器
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// 开启流式会话
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("Error opening stream: %v", err)
	}
	// 启动一个 goroutine 用于接收服务器的消息
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			log.Printf("Server: %s", resp.Message)
		}
	}()

	// 读取用户输入并发送给服务器
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		req := &pb.ChatMessage{
			User:      "Client",
			Message:   text,
			Timestamp: time.Now().Unix(),
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	// 关闭流
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}
}
