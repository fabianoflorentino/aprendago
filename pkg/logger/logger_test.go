package logger

import (
	"os"
	"testing"
)

func TestRegister(t *testing.T) {
	os.Setenv("GOENV", "development")
	os.Setenv("PWD", "../../")

	newLog := New("New test log")
	newLog.register()

	if newLog.stringToLog != "New test log" {
		t.Fatalf("Expected stringToLog to be 'New test log', got %s", newLog.stringToLog)
	}
}
