package logs

import "github.com/sirupsen/logrus"

type Log struct {
	*logrus.Entry
	LogWriter
}

func (l *Log) Flush() {
	l.LogWriter.Flush()
}

type LogConf struct {
	Level       logrus.Level
	AdapterName string
}

func InitLog(conf LogConf) *Log {
	adapterName := "std"
	if conf.AdapterName != "" {
		adapterName = conf.AdapterName
	}
	writer, ok := writerAdapter[adapterName]
	if !ok {
		adapterName = "std"
		writer, _ = writerAdapter[adapterName]
	}

	log := &Log{
		Entry:     logrus.NewEntry(logrus.New()),
		LogWriter: writer(),
	}

	log.Logger.SetOutput(log.LogWriter)
	if conf.Level != 0 {
		log.Logger.SetLevel(conf.Level)
	}

	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	log.Logger.SetReportCaller(true)
	return log
}
