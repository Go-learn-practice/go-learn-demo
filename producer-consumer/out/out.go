package out

import "fmt"

type Out struct {
	data chan interface{}
}

var out *Out

func NewOut() *Out {
	if out == nil {
		out = &Out{make(chan interface{}, 65535)}
	}
	return out
}

func Close() {
	if out != nil {
		close(out.data)
	}
}

func Println(v interface{}) {
	//写数据
	out.data <- v
}

func (o *Out) OutPut() {
	//exit := false
	//for !exit {
	//	select {
	//	case v, ok := <-o.data:
	//		if !ok {
	//			exit = true
	//			break
	//		}
	//		fmt.Println(v)
	//	}
	//}

	for v := range o.data { // 使用 `range` 正确退出循环
		fmt.Println(v)
	}
}
