package aplicacoes

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

func TestUsingJsonMarshal_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingJsonMarshal()
	cleanup()

	if !strings.Contains(buf.String(), "Fabiano") {
		t.Error("expected output to contain 'Fabiano'")
	}
}

func TestUsingJsonUnmarshal_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingJsonUnmarshal()
	cleanup()

	if !strings.Contains(buf.String(), "Fabiano") {
		t.Error("expected output to contain 'Fabiano'")
	}
}

func TestUsingJsonEncoder_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingJsonEncoder()
	cleanup()

	if !strings.Contains(buf.String(), "Fabiano") {
		t.Error("expected output to contain 'Fabiano'")
	}
}

func TestUsingPackageSort_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingPackageSort()
	cleanup()

	if !strings.Contains(buf.String(), "Sorted Strings") {
		t.Error("expected output to contain 'Sorted Strings'")
	}
}

func TestUsingCustomSort_Coverage(t *testing.T) {
	buf, cleanup := captureStdout(t)
	UsingCustomSort()
	cleanup()

	if !strings.Contains(buf.String(), "Sorted by Model") {
		t.Error("expected output to contain 'Sorted by Model'")
	}
}

func TestSortMethods_Coverage(t *testing.T) {
	carros := []carro{
		{"Fusca", 1978, 50},
		{"Brasilia", 1975, 60},
		{"Chevette", 1980, 70},
		{"Corcel", 1979, 80},
	}

	byModel := ByModel(carros)
	if byModel.Len() != 4 {
		t.Errorf("expected 4, got %d", byModel.Len())
	}

	byYear := ByYear(carros)
	if byYear.Len() != 4 {
		t.Errorf("expected 4, got %d", byYear.Len())
	}

	byPower := ByPower(carros)
	if byPower.Len() != 4 {
		t.Errorf("expected 4, got %d", byPower.Len())
	}

	if byModel.Less(0, 1) {
		t.Error("expected 'Fusca' < 'Brasilia' to be false (F > B)")
	}

	if !byModel.Less(1, 0) {
		t.Error("expected 'Brasilia' < 'Fusca' to be true (B < F)")
	}

	if !byYear.Less(0, 2) {
		t.Error("expected 1978 < 1980")
	}

	if !byPower.Less(0, 1) {
		t.Error("expected 50 < 60")
	}

	swapByModel := ByModel([]carro{
		{"Fusca", 1978, 50},
		{"Brasilia", 1975, 60},
	})
	swapByModel.Swap(0, 1)
	if swapByModel[0].Modelo != "Brasilia" {
		t.Error("expected first element after swap to be 'Brasilia'")
	}

	swapByYear := ByYear([]carro{
		{"Fusca", 1978, 50},
		{"Brasilia", 1975, 60},
	})
	swapByYear.Swap(0, 1)
	if swapByYear[0].Ano != 1975 {
		t.Error("expected first element after swap to have Ano 1975")
	}

	swapByPower := ByPower([]carro{
		{"Fusca", 1978, 50},
		{"Brasilia", 1975, 60},
	})
	swapByPower.Swap(0, 1)
	if swapByPower[0].Potencia != 60 {
		t.Error("expected first element after swap to have Potencia 60")
	}
}

func TestUsingSortStrings_Coverage(t *testing.T) {
	s := []string{"c", "a", "b"}
	usingSortSrtrings(s)
	if s[0] != "a" || s[1] != "b" || s[2] != "c" {
		t.Errorf("expected sorted [a b c], got %v", s)
	}
}

func TestUsingSortInts_Coverage(t *testing.T) {
	s := []int{3, 1, 2}
	usingSortInts(s)
	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Errorf("expected sorted [1 2 3], got %v", s)
	}
}
