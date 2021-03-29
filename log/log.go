package log

import (
	"log"
	"os"
)

type Log struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func init() {
	logFilePath := "info_log.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalln("open log file failed")
	}
	Info := log.New(logFile, "[Info]", log.Llongfile)
	Warn := log.New(logFile, "[Warn]", log.Llongfile)
	Error := log.New(logFile, "[Error]", log.Llongfile)
	Info.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	Warn.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	Error.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	Info.Printf("test info log")
}
