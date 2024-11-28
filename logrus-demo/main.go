package main

import (
	"github.com/sirupsen/logrus"
	"logrus-demo/logs"
)

var log *logs.Log

func init() {
	conf := logs.LogConf{
		Level:       logrus.InfoLevel,
		AdapterName: "std",
	}
	log = logs.InitLog(conf)
}

func main() {
	log.Info()
}
