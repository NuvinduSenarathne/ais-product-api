package logger

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	Logger = log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("Logger initialized successfully")
}
