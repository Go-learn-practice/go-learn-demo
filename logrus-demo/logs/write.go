package logs

import "io"

var writerAdapter = make(map[string]InitLogWriterFunc)

type InitLogWriterFunc func() LogWriter

type LogWriter interface {
	Flush()
	io.Writer
}

func RegisterInitWriterFunc(adapterName string, writerFunc InitLogWriterFunc) {
	writerAdapter[adapterName] = writerFunc
}
