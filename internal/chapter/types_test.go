package chapter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTopicsWithRealChapter(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	projRoot := filepath.Join(root, "..", "..")
	ch := &Chapter{
		Number:  1,
		Title:   "Visão Geral do Curso",
		RootDir: filepath.Join(projRoot, "internal", "visao_geral_do_curso"),
	}

	topics, err := ch.Topics()
	if err != nil {
		t.Fatalf("Topics() error: %v", err)
	}
	if len(topics) == 0 {
		t.Error("expected at least one topic")
	}
}

func TestExecTopicWithRealChapter(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	projRoot := filepath.Join(root, "..", "..")
	ch := &Chapter{
		Number:  1,
		Title:   "Visão Geral do Curso",
		RootDir: filepath.Join(projRoot, "internal", "visao_geral_do_curso"),
	}

	err = ch.ExecTopic("Bem-vindo!")
	if err != nil {
		t.Fatalf("ExecTopic() error: %v", err)
	}
}

func TestOverviewWithRealChapter(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	projRoot := filepath.Join(root, "..", "..")
	ch := &Chapter{
		Number:  1,
		Title:   "Visão Geral do Curso",
		RootDir: filepath.Join(projRoot, "internal", "visao_geral_do_curso"),
	}

	err = ch.Overview()
	if err != nil {
		t.Fatalf("Overview() error: %v", err)
	}
}
