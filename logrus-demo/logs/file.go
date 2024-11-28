package logs

import "os"

const LOGPATH = "runtime/logs/logrus.log"

type fileWriter struct {
	*os.File
}

func (s *fileWriter) Flush() {
	err := s.Sync()
	if err != nil {
		return
	}
}

func newFileWriter() LogWriter {
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		file = os.Stderr
	}

	return &fileWriter{file}
}

func init() {
	RegisterInitWriterFunc("file", newFileWriter)
}
