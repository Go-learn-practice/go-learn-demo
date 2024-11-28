package main

import (
	"fmt"
	"os"
)

func main() {
	//_chan.Run()

	//_json.RunStruct2Json()
	//_json.RunJson2Struct()
	//_json.Nested()

	//_ctx.Run()

	ptr()
}

type A struct {
	*os.File
}

func ptr() {
	a := &A{os.Stderr}
	fmt.Println(a)
}
