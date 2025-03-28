// Package logger provides a simple logging utility that allows creating log messages
// and writing them to a file. It includes functionality to initialize a logger with
// a custom message and register the log output to a file.
//
// The log file is created in a "log/" directory under the current working directory,
// and its name is determined by the "GOENV" environment variable.
package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Logger is a struct that holds the string to be logged.
type Logger struct {
	stringToLog string
}

// New creates a new instance of Logger with the provided strings to log.
// It joins the input strings with a space and initializes the Logger.
// If no strings are provided, it returns an error indicating that the input cannot be empty.
//
// Parameters:
//
//	stringToLog - A variadic parameter representing the strings to be logged.
//
// Returns:
//
//	*Logger - A pointer to the newly created Logger instance.
//	error - An error if the input is invalid (e.g., no strings provided).
func New(stringToLog ...string) *Logger {
	if len(stringToLog) == 0 {
		log.Fatalf("string to log cannot be empty")
	}

	return &Logger{stringToLog: strings.Join(stringToLog, " ")}
}

// Register creates a log directory and a log file based on the current working directory
// and the environment variable "GOENV". It initializes a logger to write logs to the file
// and logs the string stored in the Logger instance.
//
// The log directory is created under the current working directory with the name "log/",
// and the log file is named using the value of the "GOENV" environment variable with a ".log" extension.
//
// Returns an error if the log directory cannot be created or if the log file cannot be opened.
func (l *Logger) Register() error {
	logDir := "log/"

	if err := os.MkdirAll(os.Getenv("PWD")+"/"+logDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	file, err := os.OpenFile(os.Getenv("PWD")+"/"+logDir+"/"+os.Getenv("GOENV")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(l.stringToLog)

	return nil
}
