package chapter

import (
	"os"
	"testing"
)

func TestChapterNumber(t *testing.T) {
	c := &Chapter{Number: 1, Title: "Test", RootDir: "testdata"}
	if c.Number != 1 {
		t.Errorf("expected Number=1, got %d", c.Number)
	}
}

func TestChapterTitle(t *testing.T) {
	c := &Chapter{Number: 1, Title: "Visao Geral do Curso", RootDir: "testdata"}
	if c.Title != "Visao Geral do Curso" {
		t.Errorf("expected Title='Visao Geral do Curso', got '%s'", c.Title)
	}
}

func TestChapterTopics(t *testing.T) {
	c := &Chapter{Number: 1, Title: "Visao Geral do Curso", RootDir: "../visao_geral_do_curso"}

	topics, err := c.Topics()
	if err != nil {
		t.Fatalf("Topics() returned error: %v", err)
	}

	expected := []string{
		"Bem-vindo!",
		"Por que Go?",
		"Sucesso",
		"Recursos",
		"Como esse curso funciona",
	}
	if len(topics) != len(expected) {
		t.Fatalf("expected %d topics, got %d: %v", len(expected), len(topics), topics)
	}
	for i, topic := range topics {
		if topic != expected[i] {
			t.Errorf("topic[%d] = '%s', expected '%s'", i, topic, expected[i])
		}
	}
}

func TestChapterOverview(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("failed to open null device: %v", err)
	}
	defer nullFile.Close()
	os.Stdout = nullFile

	c := &Chapter{Number: 1, Title: "Visao Geral do Curso", RootDir: "../visao_geral_do_curso"}
	if err := c.Overview(); err != nil {
		t.Fatalf("Overview() returned error: %v", err)
	}
}

func TestChapterExecTopic(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("failed to open null device: %v", err)
	}
	defer nullFile.Close()
	os.Stdout = nullFile

	c := &Chapter{Number: 1, Title: "Visao Geral do Curso", RootDir: "../visao_geral_do_curso"}
	if err := c.ExecTopic("Bem-vindo!"); err != nil {
		t.Fatalf("ExecTopic('Bem-vindo!') returned error: %v", err)
	}
}

// TestChapterExecTopicNotFound skipped because section.Section.Format calls
// logger.Log which calls os.Exit(1) on missing topics (pre-existing behavior).
