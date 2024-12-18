package bufBytes

import (
	"bytes"
	"fmt"
)

// bytes.Buffer 是 Go 中高效、简洁的字节缓冲区，常用于需要频繁读写字节数据的场景，特别是需要拼接或分段处理数据时

/*
常见用途
1. 字符串拼接： 使用 bytes.Buffer 拼接字符串比使用 + 更高效
2. 处理文件内容： 用于对文件内容进行缓冲读取或写入
3. HTTP 响应/请求： 构建 HTTP 请求体或解析 HTTP 响应数据
*/

func Writer() {
	buf := bytes.Buffer{}

	buf.Write([]byte("Hello, "))
	buf.WriteString("World!")

	fmt.Println(buf.String())
}

func Reader() {
	buf := bytes.NewBuffer([]byte("Hello, Go!"))

	// 读取固定大小的数据
	data := make([]byte, 5)
	n, err := buf.Read(data)

	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Println(string(data[:n]))

	// 读取单个字节
	b, _ := buf.ReadByte()
	fmt.Println(string(b))

	// 读取剩余数据
	fmt.Println(buf.String())
}

func Reset() {
	buf := bytes.NewBuffer([]byte("Temporary Data"))
	buf.Reset()
	fmt.Println(buf.Len())
}

func Print() {
	var buf bytes.Buffer
	fmt.Println(buf)
}
