package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// findModuleRoot walks up from cwd to locate the directory containing go.mod.
func findModuleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

// chdirToModuleRoot changes the working directory to the module root so that
// registered Chapter RootDir paths (relative to the module root) resolve correctly.
// Returns a restore function.
func chdirToModuleRoot(t *testing.T) func() {
	t.Helper()
	root := findModuleRoot()
	if root == "" {
		t.Fatal("could not find module root (go.mod)")
	}
	cwd, _ := os.Getwd()
	if root != cwd {
		if err := os.Chdir(root); err != nil {
			t.Fatalf("failed to chdir to %s: %v", root, err)
		}
	}
	return func() { os.Chdir(cwd) }
}

// captureStdout replaces os.Stdout with a pipe and returns a cleanup function
// that drains the pipe into the returned buffer. Call cleanup() BEFORE reading
// the buffer to ensure all output has been captured.
func captureStdout(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(&buf, r)
		close(done)
	}()

	return &buf, func() {
		_ = w.Close()
		<-done
		os.Stdout = old
	}
}

// wantOutput is a helper for string containment checks.
type wantOutput struct {
	substring string
	msg       string
}

// assertContainsAll checks that output contains all expected substrings.
func assertContainsAll(t *testing.T, output string, wants []wantOutput) {
	t.Helper()
	for _, w := range wants {
		if !strings.Contains(output, w.substring) {
			t.Errorf("%s: output missing %q", w.msg, w.substring)
		}
	}
}

// ---------------------------------------------------------------------------
// caps — list chapters
// ---------------------------------------------------------------------------

func TestCapsListsAllChapters(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := runCaps(); err != nil {
		t.Fatal(err)
	}
	cleanup()

	output := buf.String()

	if !strings.Contains(output, "Capítulos do Curso") {
		t.Error("caps output missing header")
	}

	lines := strings.Split(output, "\n")
	var count int
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "aprendago cap") {
			count++
		}
	}
	if count != 27 {
		t.Errorf("expected 27 chapters, got %d", count)
	}

	assertContainsAll(t, output, []wantOutput{
		{"aprendago cap 1 topics    Visão Geral do Curso", "cap 1"},
		{"aprendago cap 8 topics    Agrupamento de Dados", "cap 8"},
		{"aprendago cap 14 topics    Ponteiros", "cap 14"},
		{"aprendago cap 27 topics    Teste e Benchmarking", "cap 27"},
	})
}

func TestCapsNoLegacyFlagReferences(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	_ = runCaps()
	cleanup()

	if strings.Contains(buf.String(), "--cap") {
		t.Error("caps output should not reference legacy --cap flags")
	}
}

// ---------------------------------------------------------------------------
// cap — access chapters
// ---------------------------------------------------------------------------

func TestCapTopics_Chapter8(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := capCmd.RunE(capCmd, []string{"8", "topics"}); err != nil {
		t.Fatal(err)
	}
	cleanup()

	assertContainsAll(t, buf.String(), []wantOutput{
		{"Agrupamento de Dados", "chapter title"},
		{"Array", "topic"},
		{"Maps: introdução", "topic"},
		{"Maps: range e deletando", "topic"},
	})
}

func TestCapOverviewDefault_Chapter1(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := capCmd.RunE(capCmd, []string{"1"}); err != nil {
		t.Fatal(err)
	}
	cleanup()

	assertContainsAll(t, buf.String(), []wantOutput{
		{"Visão Geral do Curso", "chapter title"},
		{"Bem-vindo!", "first topic"},
	})
}

func TestCapOverviewExplicit_Chapter8(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := capCmd.RunE(capCmd, []string{"8", "overview"}); err != nil {
		t.Fatal(err)
	}
	cleanup()

	if !strings.Contains(buf.String(), "Agrupamento de Dados") {
		t.Error("cap 8 overview missing chapter title")
	}
}

func TestCapSpecificTopic(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := capCmd.RunE(capCmd, []string{"8", "Array"}); err != nil {
		t.Fatal(err)
	}
	cleanup()

	if !strings.Contains(buf.String(), "Array") {
		t.Error("cap 8 'Array' should contain the topic content")
	}
}

func TestCapInvalidChapter(t *testing.T) {
	err := capCmd.RunE(capCmd, []string{"999"})
	if err == nil {
		t.Fatal("expected error for non-existent chapter, got nil")
	}
	if !strings.Contains(err.Error(), "não encontrado") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCapNonNumericArg(t *testing.T) {
	err := capCmd.RunE(capCmd, []string{"abc"})
	if err == nil {
		t.Fatal("expected error for non-numeric chapter, got nil")
	}
}

// ---------------------------------------------------------------------------
// outline — full course outline
// ---------------------------------------------------------------------------

func TestOutlineShowsChapters(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	if err := runOutline(); err != nil {
		t.Fatal(err)
	}
	cleanup()

	output := buf.String()

	assertContainsAll(t, output, []wantOutput{
		{"Visão Geral do Curso", "first chapter"},
		{"Agrupamento de Dados", "middle chapter"},
		{"Ponteiros", "middle chapter"},
		{"Teste e Benchmarking", "last chapter"},
	})

	if len(output) < 200 {
		t.Errorf("outline too short (%d chars)", len(output))
	}
}

// ---------------------------------------------------------------------------
// root — help and error handling
// ---------------------------------------------------------------------------

func TestRootHelp(t *testing.T) {
	buf, cleanup := captureStdout(t)
	if err := rootRun(rootCmd, []string{}); err != nil {
		t.Fatal(err)
	}
	cleanup()

	assertContainsAll(t, buf.String(), []wantOutput{
		{"aprendago", "command name"},
		{"Available Commands", "commands section"},
		{"caps", "caps subcommand"},
		{"cap", "cap subcommand"},
		{"outline", "outline subcommand"},
	})
}

func TestRootUnknownCommand(t *testing.T) {
	err := rootRun(rootCmd, []string{"unknown"})
	if err == nil {
		t.Fatal("expected error for unknown command, got nil")
	}
	if !strings.Contains(err.Error(), "comando desconhecido") {
		t.Errorf("unexpected error message: %v", err)
	}
}

// ---------------------------------------------------------------------------
// integration — every chapter is accessible and has content
// ---------------------------------------------------------------------------

func TestAllChapters_OverviewAccessible(t *testing.T) {
	defer chdirToModuleRoot(t)()

	for i := 1; i <= 27; i++ {
		num := strconv.Itoa(i)
		err := capCmd.RunE(capCmd, []string{num})
		if err != nil {
			t.Errorf("chapter %d overview failed: %v", i, err)
		}
	}
}

func TestAllChapters_TopicsAccessible(t *testing.T) {
	defer chdirToModuleRoot(t)()

	for i := 1; i <= 27; i++ {
		num := strconv.Itoa(i)
		buf, cleanup := captureStdout(t)
		err := capCmd.RunE(capCmd, []string{num, "topics"})
		if err != nil {
			t.Errorf("chapter %d topics failed: %v", i, err)
			cleanup()
			continue
		}
		cleanup()

		if strings.TrimSpace(buf.String()) == "" {
			t.Errorf("chapter %d topics returned empty output", i)
		}
	}
}

func TestAllChapters_HaveAtLeastOneTopic(t *testing.T) {
	defer chdirToModuleRoot(t)()

	for i := 1; i <= 27; i++ {
		num := strconv.Itoa(i)
		buf, cleanup := captureStdout(t)
		err := capCmd.RunE(capCmd, []string{num, "topics"})
		if err != nil {
			t.Errorf("chapter %d topics failed: %v", i, err)
			cleanup()
			continue
		}
		cleanup()

		lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
		var nonEmpty int
		for _, l := range lines {
			if strings.TrimSpace(l) != "" {
				nonEmpty++
			}
		}
		if nonEmpty < 2 {
			t.Errorf("chapter %d has <2 non-empty lines (title + at least one topic): %q", i, buf.String())
		}
	}
}

// ---------------------------------------------------------------------------
// printChapterTopics helper
// ---------------------------------------------------------------------------

func TestPrintChapterTopicsForChapter1(t *testing.T) {
	defer chdirToModuleRoot(t)()

	buf, cleanup := captureStdout(t)
	c := chapter.Get(1)
	if c == nil {
		t.Fatal("chapter 1 not registered")
	}
	if err := printChapterTopics(c); err != nil {
		t.Fatal(err)
	}
	cleanup()

	assertContainsAll(t, buf.String(), []wantOutput{
		{"Visão Geral do Curso", "chapter title"},
		{"Bem-vindo!", "topic"},
	})
}
