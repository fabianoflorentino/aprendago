package tui

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func countLines(s string) int {
	n := strings.Count(s, "\n")
	if len(s) > 0 && s[len(s)-1] != '\n' {
		n++
	}
	return n
}

func TestView_FillsTerminalHeight(t *testing.T) {
	chapters := chapter.All()
	if len(chapters) == 0 {
		t.Skip("no chapters registered")
	}

	m := NewModel(chapters)
	m.width = 100
	m.height = 40
	m.ready = true

	got := m.View()
	lines := countLines(got)

	if lines != m.height {
		t.Errorf("View() produced %d lines, want %d (terminal height=%d)\n%q",
			lines, m.height, m.height, got)
	}
}

func TestView_FillsTerminalHeight_Small(t *testing.T) {
	chapters := chapter.All()
	if len(chapters) == 0 {
		t.Skip("no chapters registered")
	}

	m := NewModel(chapters)
	m.width = 60
	m.height = 15
	m.ready = true

	got := m.View()
	lines := countLines(got)

	if lines != m.height {
		t.Errorf("View() produced %d lines, want %d (terminal height=%d)\n%q",
			lines, m.height, m.height, got)
	}
}

func TestView_FillsTerminalHeight_Wide(t *testing.T) {
	chapters := chapter.All()
	if len(chapters) == 0 {
		t.Skip("no chapters registered")
	}

	m := NewModel(chapters)
	m.width = 200
	m.height = 50
	m.ready = true

	got := m.View()
	lines := countLines(got)

	if lines != m.height {
		t.Errorf("View() produced %d lines, want %d (terminal height=%d)\n%q",
			lines, m.height, m.height, got)
	}
}

func TestView_NotReady(t *testing.T) {
	m := Model{ready: false}
	got := m.View()
	if !strings.Contains(got, "Carregando") {
		t.Errorf("unready View() should show loading, got: %q", got)
	}
}

func TestView_Quitting(t *testing.T) {
	m := Model{ready: true, quitting: true}
	got := m.View()
	if got != "" {
		t.Errorf("quitting View() should be empty, got: %q", got)
	}
}
