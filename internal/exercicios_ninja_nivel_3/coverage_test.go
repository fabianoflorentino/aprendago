package exercicios_ninja_nivel_3

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

	if !strings.Contains(buf.String(), "1") {
		t.Error("expected output to contain '1'")
	}
}

func TestResolucaoNaPraticaExercicio2_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio2()
	cleanup()

	if !strings.Contains(buf.String(), "65") {
		t.Error("expected output to contain '65'")
	}
}

func TestResolucaoNaPraticaExercicio3_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio3()
	cleanup()

	if !strings.Contains(buf.String(), "1985") {
		t.Error("expected output to contain '1985'")
	}
}

func TestResolucaoNaPraticaExercicio4_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio4()
	cleanup()

	if !strings.Contains(buf.String(), "1985") {
		t.Error("expected output to contain '1985'")
	}
}

func TestResolucaoNaPraticaExercicio5_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio5()
	cleanup()

	if !strings.Contains(buf.String(), "10: 2") {
		t.Error("expected output to contain '10: 2'")
	}
}

func TestResolucaoNaPraticaExercicio6_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio6()
	cleanup()

	if !strings.Contains(buf.String(), "Sim") {
		t.Error("expected output to contain 'Sim'")
	}
}

func TestResolucaoNaPraticaExercicio7_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio7()
	cleanup()

	if !strings.Contains(buf.String(), "Sim") {
		t.Error("expected output to contain 'Sim'")
	}
}

func TestResolucaoNaPraticaExercicio8_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio8()
	cleanup()

	if !strings.Contains(buf.String(), "Não deve ser impresso") {
		t.Error("expected output to contain 'Não deve ser impresso'")
	}
}

func TestResolucaoNaPraticaExercicio9_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio9()
	cleanup()

	if !strings.Contains(buf.String(), "futebol americano") {
		t.Error("expected output to contain 'futebol americano'")
	}
}

func TestResolucaoNaPraticaExercicio10_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	ResolucaoNaPraticaExercicio10()
	cleanup()

	if !strings.Contains(buf.String(), "true") {
		t.Error("expected output to contain 'true'")
	}
}
