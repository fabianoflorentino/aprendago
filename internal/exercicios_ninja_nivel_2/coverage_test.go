package exercicios_ninja_nivel_2

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

func TestResolucaoNaPraticaExercicio1_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio1()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}

func TestResolucaoNaPraticaExercicio2_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio2()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}

func TestResolucaoNaPraticaExercicio3_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio3()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}

func TestResolucaoNaPraticaExercicio4_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio4()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}

func TestResolucaoNaPraticaExercicio5_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio5()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}

func TestResolucaoNaPraticaExercicio6_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio6()
	cleanup()

	if !strings.Contains(buf.String(), "Resolução:") {
		t.Error("expected output to contain 'Resolução:'")
	}
}
