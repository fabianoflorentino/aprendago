package agrupamento_de_dados

import (
	"os"
	"testing"
)

func init() {
	rootDir = "./"
}

func TestTopics(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("failed to open null device: %v", err)
	}
	defer nullFile.Close()

	os.Stdout = nullFile

	Topics()

	if err != nil {
		t.Errorf("Topics() failed: %v", err)
	}
}
