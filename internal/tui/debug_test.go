package tui

import (
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fabianoflorentino/aprendago/internal/chapter"
	_ "github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
)

func TestDebugViewOutput(t *testing.T) {
	chapters := chapter.All()
	if len(chapters) == 0 {
		t.Skip("no chapters registered")
	}

	m := NewModel(chapters)
	// simulate a wide terminal to inspect the right panel metrics
	model, _ := m.handleWindowSize(tea.WindowSizeMsg{Width: 140, Height: 40})
	m = model.(Model)
	m.ready = true

	// Diagnostics: compute and print layout metrics used by the view
	leftTotal := int(float64(m.width) * 0.35)
	rightTotal := m.width - leftTotal - 1
	leftContentW := leftTotal - leftPanelFH
	rightContentW := rightTotal - rightPanelFH - 1
	mainH := m.height - 4
	rightContentH := mainH - rightPanelFV

	fmt.Printf("metrics: width=%d height=%d leftTotal=%d rightTotal=%d leftContentW=%d rightContentW=%d mainH=%d rightContentH=%d\n",
		m.width, m.height, leftTotal, rightTotal, leftContentW, rightContentW, mainH, rightContentH)

	view := m.View()
	fmt.Printf("title: %q\n", m.titleView())
	fmt.Printf("main view (preview):\n%s\n", m.mainView())
	fmt.Printf("footer: %q\n", m.footerView())
	lines := strings.Split(view, "\n")
	fmt.Printf("total lines: %d\n", len(lines))
	for i, line := range lines {
		fmt.Printf("%02d|%q\n", i+1, line)
	}
}
