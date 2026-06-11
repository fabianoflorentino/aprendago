package funcoes

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

func TestResolucaoDesafioCallback_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoDesafioCallback()
	cleanup()

	if !strings.Contains(buf.String(), "Soma dos números ímpares") {
		t.Error("expected output to contain 'Soma dos números ímpares'")
	}
}

func TestSoma_Coverage(t *testing.T) {
	result := soma(1, 2, 3)
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}

func TestSomaImapres_Coverage(t *testing.T) {
	result := somaImapres(soma, 1, 2, 3, 4, 5)
	if result != 9 {
		t.Errorf("expected 9 (1+3+5), got %d", result)
	}
}

func TestSomaImapresNoOdds_Coverage(t *testing.T) {
	result := somaImapres(soma, 2, 4, 6)
	if result != 0 {
		t.Errorf("expected 0 (no odds), got %d", result)
	}
}
