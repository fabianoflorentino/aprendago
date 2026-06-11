package cmd

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestPrintChapterTopicsError(t *testing.T) {
	c := &chapter.Chapter{
		Number:  99,
		Title:   "Invalid Chapter",
		RootDir: "/nonexistent/path",
	}
	err := printChapterTopics(c)
	if err == nil {
		t.Fatal("expected error for invalid chapter root dir, got nil")
	}
	if !strings.Contains(err.Error(), "erro ao listar tópicos") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestPrintChapterTopicsEmptyTopics(t *testing.T) {
	// Use a valid chapter and just verify it prints the title
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	c := chapter.Get(1)
	if c == nil {
		t.Fatal("chapter 1 not registered")
	}
	err := printChapterTopics(c)
	cleanup()

	if err != nil {
		t.Fatalf("printChapterTopics() returned error: %v", err)
	}
	if !strings.Contains(buf.String(), "Visão Geral do Curso") {
		t.Error("output missing chapter title")
	}
}
