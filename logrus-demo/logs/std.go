package logs

import "os"

type stdWriter struct {
	*os.File
}

func (s *stdWriter) Flush() {
	err := s.Sync()
	if err != nil {
		return
	}
}

func newStdWriter() LogWriter {
	return &stdWriter{
		os.Stderr,
	}
}

func init() {
	RegisterInitWriterFunc("std", newStdWriter)
}
