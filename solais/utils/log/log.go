package log

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open the logging file", err)
	}

	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}
