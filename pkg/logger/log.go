/*
Package logger provides a method to log messages to a file and print them to the console
  - Log: writes a log message to a file and prints it to the console
*/
package logger

import (
	"fmt"
	"log"
	"os"
)

// Log writes a log message to a file and prints it to the console
//   - format: format of the log message
//   - args: arguments to be formatted
//
// The method opens a file in the log directory and writes the log message
// If an error occurs while opening the file, the method logs an error message
// and exits the program
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
