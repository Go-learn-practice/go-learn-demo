package _flag

import (
	"flag"
	"fmt"
	"time"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	host    = flag.String("host", "localhost", "The server host")
	debug   = flag.Bool("debug", false, "Enable debug mode")
	timeout = flag.Duration("timeout", 5*time.Second, "Timeout duration")
	rate    = flag.Float64("rate", 0.8, "Rate limit")
)

// go run main.go --host=127.0.0.1 --port=9090 --debug=true --timeout=10s --rate=1.2
func RunFlags() {
	flag.Parse()

	// 访问解析后的参数值
	fmt.Println("Server will start with the following parameters:")
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
	fmt.Println("Debug:", *debug)
	fmt.Println("Timeout:", *timeout)
	fmt.Println("Rate:", *rate)
}

// go run main.go --host=127.0.0.1 --port=9090
func RunFlagsVar() {
	var port int
	var host string

	// 将标志绑定到现有变量上
	flag.IntVar(&port, "port", 8080, "The server port")
	flag.StringVar(&host, "host", "localhost", "The server hostname")

	// 解析命令行参数
	flag.Parse()

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %d\n", port)
}
