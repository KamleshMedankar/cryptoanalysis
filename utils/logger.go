package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogRequest(requestID, status string) {
	Logger.Printf("RequestID: %s, Status: %s\n", requestID, status)
}
