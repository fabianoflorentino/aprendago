package aplicacoes

import (
	"os"
	"testing"
)

func TestHelp(t *testing.T) {
	var err error

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("failed to open null device: %v", err)
	}
	defer nullFile.Close()

	os.Stdout = nullFile

	Help()

	if err != nil {
		t.Errorf("Help() failed: %v", err)
	}
}
