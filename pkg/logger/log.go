package logger

import (
	"fmt"
	"log"
	"os"
)

func Log(format string, args ...any) {
	file, err := os.OpenFile("log/"+os.Getenv("GOENV")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Printf(format, args...)

	fmt.Printf(format+"\n", args...)
	os.Exit(1)
}
