package _log

import (
	"flag"
	"log"
	"os"
)

func RunNormal() {
	log.Print("这是一个普通日志")
	log.Println("这是一个带换行的日志")
	log.Printf("这是一个带格式的日志: %d", 100)
}

// go run main.go --code=1
func RunError() {
	var code int

	flag.IntVar(&code, "code", 0, "错误码")

	//解析命令行参数
	flag.Parse()

	switch code {
	case 0:
		log.Fatal("这是一个致命错误日志")
	case 1:
		log.Fatalf("致命错误代码: %d", 123)
	case 2:
		log.Panic("发生异常，程序将崩溃")
	case 3:
		log.Panicf("崩溃原因: %s", "某个问题")
	}
}

/*
log.New 函数用于创建一个新的日志记录器
func New(out io.Writer, prefix string, flag int) *Logger

第一个参数是日志输出的目标，可以是文件、标准输出等
第二个参数是日志前缀，用于标识日志的来源
第三个参数是日志记录的格式，包括时间、文件名、行号等信息
log.Ldate | log.Ltime | log.Lshortfile
log.Ldate 表示日志中包含日期信息 格式为 YYYY/MM/DD
log.Ltime：在日志中输出时间，格式为 HH:MM:SS
log.Lshortfile：在日志中输出文件名和行号（例如：main.go:12）
*/

/*
os.OpenFile 函数用于打开或创建文件
func OpenFile(name string, flag int, perm FileMode) (*File, error)
1. name: 文件的路径名称，例如 log/custom.log。
2. flag: 指定文件的打开模式（如读、写、追加等）。
3. perm: 用于设置新建文件的权限（使用类 Unix 权限表示）。

os.O_CREATE|os.O_WRONLY|os.O_APPEND： 这是一个组合的 flag 参数，多个标志通过 按位或运算符 | 连接起来
os.O_CREATE：如果文件不存在，则创建文件。

os.O_WRONLY：以写入模式打开文件，不允许读取操作。

os.O_APPEND：写入时，文件的内容会追加到文件末尾，而不是覆盖文件内容
*/

func RunCustom() {
	// 创建日志文件
	file, _ := os.OpenFile("log/custom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	defer file.Close()

	// 创建自定义日志器
	infoLogger := log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	// 使用自定义日志器
	infoLogger.Println("这是一个信息日志")
	errorLogger.Println("这是一个错误日志")
}
