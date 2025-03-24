package logger

import (
	"testing"
)

func TestRegister(t *testing.T) {
	var stringNewLog string = "New test log"

	newLog, err := New(stringNewLog)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	newLog.Register()

	if newLog.stringToLog != "New test log" {
		t.Fatalf("Expected stringToLog to be 'New test log', got %s", newLog.stringToLog)
	}
}
