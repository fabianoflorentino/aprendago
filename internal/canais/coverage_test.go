package canais

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&buf, r)
		close(done)
	}()

	return &buf, func() {
		w.Close()
		<-done
		os.Stdout = old
	}
}

func TestUsingConverge_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingConverge()
	cleanup()

	if !strings.Contains(buf.String(), "Even Value") {
		t.Error("expected output to contain 'Even Value'")
	}
	if !strings.Contains(buf.String(), "Odd Value") {
		t.Error("expected output to contain 'Odd Value'")
	}
}

func TestUsingDivergence_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingDivergence()
	cleanup()

	if !strings.Contains(buf.String(), "Received") {
		t.Error("expected output to contain 'Received'")
	}
}
