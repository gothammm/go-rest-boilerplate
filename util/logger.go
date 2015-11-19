package util

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func createLogger(
	infoHandle io.Writer,
	debugHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Lshortfile|log.Ltime)

	Debug = log.New(debugHandle, "DEBUG: ", log.Llongfile|log.Ldate|log.Ltime)

	Error = log.New(errorHandle, "ERROR: ", log.Llongfile|log.Ldate|log.Ltime)
}

func InitLogger() {
	createLogger(os.Stdout, os.Stdout, os.Stderr)
}