package log

import (
	"log"
	"os"
)

var serverLog *log.Logger

func initLog() {
	serverLog = log.New(
		os.Stdout,
		"Server: ",
		log.Ldate|log.Ltime,
	)
}

func getLogger() *log.Logger {
	if serverLog == nil {
		initLog()
	}
	return serverLog
}

func Logf(format string, args ...interface{}) {
	getLogger().Printf(format, args...)
}

func Log(msg string) {
	getLogger().Println(msg)
}
