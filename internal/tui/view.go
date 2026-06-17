package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	leftPanelFH  = 6 // border(2) + padding(4)
	leftPanelFV  = 4 // border(2) + padding(2)
	rightPanelFH = 6 // border(2) + padding(4)
	rightPanelFV = 4 // border(2) + padding(2)
)

func (m Model) View() string {
	if !m.ready {
		return "\n  Carregando..."
	}
	if m.quitting {
		return ""
	}

	rendered := lipgloss.JoinVertical(
		lipgloss.Left,
		m.titleView(),
		m.mainView(),
		m.footerView(),
	)
	return strings.TrimRight(rendered, "\n")
}

func (m Model) titleView() string {
	title := " aprendago "
	if m.selectedChapter != nil {
		title = fmt.Sprintf(" aprendago — %s", m.selectedChapter.Title)
	}
	if m.selectedTopic != "" {
		title = fmt.Sprintf("%s: %s", title, m.selectedTopic)
	}

	padL := titleStyle.GetPaddingLeft()
	padR := titleStyle.GetPaddingRight()
	maxContent := m.width - padL - padR
	if maxContent < 1 {
		maxContent = 1
	}

	visibleW := 0
	var truncated strings.Builder
	for _, r := range title {
		rw := lipgloss.Width(string(r))
		if visibleW+rw > maxContent {
			break
		}
		truncated.WriteRune(r)
		visibleW += rw
	}

	rendered := titleStyle.Width(m.width).Render(truncated.String())
	return strings.TrimRight(rendered, " \n")
}

func (m Model) mainView() string {
	leftPct := 0.35
	gapW := 0
	leftTotal := int(float64(m.width) * leftPct)
	rightTotal := m.width - leftTotal - gapW

	if leftTotal < 1 || rightTotal < 1 {
		return lipgloss.NewStyle().Width(m.width).Render("Janela muito pequena")
	}

	mainH := m.height - 4
	if mainH < 4 {
		return lipgloss.NewStyle().Width(m.width).Render("Janela muito pequena")
	}

	leftContentW := leftTotal - leftPanelFH
	rightContentW := rightTotal - rightPanelFH - 1 // -1 for scrollbar
	rightContentH := mainH - rightPanelFV

	leftContent := m.renderListContent()
	if leftContent == "" {
		leftContent = lipgloss.NewStyle().Width(leftContentW).Render("Nenhum item")
	}
	leftHelp := helpStyle.Width(leftContentW).Render(m.helpText())
	leftScrollbar := m.renderScrollbar(leftContentW)

	leftBox := lipgloss.JoinVertical(lipgloss.Left, m.renderLeftHeader(), leftContent, leftScrollbar, leftHelp)
	clippedLeft := lipgloss.NewStyle().MaxHeight(mainH).Render(leftBox)
	leftPanel := leftPanelStyle.Width(leftTotal).Height(mainH).Render(clippedLeft)

	// Build a full-height background block for the right content area
	var rawContent string
	if m.err != nil {
		rawContent = fmt.Sprintf("Erro: %v", m.err)
	} else if m.content == "" && m.selectedTopic == "" {
		rawContent = "Selecione um tópico"
	} else if m.content == "" {
		rawContent = "Carregando..."
	} else {
		rawContent = m.contentPane.View()
	}

	// Split into lines and render each line into a background-styled row
	contentLines := strings.Split(rawContent, "\n")
	bgStyle := lipgloss.NewStyle().Width(rightContentW).Background(lipgloss.Color("#0b0f12"))
	var rows []string
	for i := 0; i < rightContentH; i++ {
		var line string
		if i < len(contentLines) {
			line = contentLines[i]
		} else {
			line = ""
		}
		// Render the line with width so it fills and gets the background
		rows = append(rows, bgStyle.Render(line))
	}

	rightBox := strings.Join(rows, "\n")

	rightScrollbar := m.renderRightScrollbar(rightContentH)
	if rightScrollbar != "" {
		rightBox = lipgloss.JoinHorizontal(lipgloss.Top, rightBox, rightScrollbar)
	}

	clippedRight := lipgloss.NewStyle().MaxHeight(mainH).Render(rightBox)
	rightPanel := rightPanelStyle.Width(rightTotal).Height(mainH).Render(clippedRight)
	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
}

func (m Model) renderLeftHeader() string {
	title := "Capítulos"
	if m.state == stateTopics {
		title = "Tópicos"
	}
	return lipgloss.NewStyle().
		Bold(true).
		Render(title)
}

func (m Model) renderListContent() string {
	var content string
	switch m.state {
	case stateChapters:
		content = m.chapterList.View()
	case stateTopics:
		content = m.topicList.View()
	}
	if content == "" {
		return ""
	}

	leftContentW := int(float64(m.width)*0.35) - leftPanelFH
	scrollX := 0
	if m.scrollX != nil {
		scrollX = *m.scrollX
	}
	return scrollLines(content, scrollX, leftContentW)
}

func (m Model) renderScrollbar(width int) string {
	if m.scrollX == nil || *m.scrollX == 0 {
		return ""
	}

	barWidth := width
	scrollX := *m.scrollX

	maxWidth := scrollX + width
	if maxWidth <= 0 {
		return ""
	}

	visibleRatio := float64(width) / float64(maxWidth)
	visibleBlocks := int(visibleRatio * float64(barWidth))
	if visibleBlocks < 1 {
		visibleBlocks = 1
	}
	if visibleBlocks > barWidth {
		visibleBlocks = barWidth
	}

	totalBlocks := barWidth
	startPos := int(float64(scrollX) / float64(maxWidth) * float64(totalBlocks))
	if startPos < 0 {
		startPos = 0
	}
	if startPos >= totalBlocks {
		startPos = totalBlocks - visibleBlocks
	}

	var bar string
	for i := 0; i < totalBlocks; i++ {
		if i >= startPos && i < startPos+visibleBlocks {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Render(bar)
}

func (m Model) renderRightScrollbar(height int) string {
	if height <= 0 || m.content == "" {
		return ""
	}

	contentLines := m.contentPane.TotalLineCount()
	if contentLines <= height {
		return ""
	}

	scrollPercent := m.contentPane.ScrollPercent()
	visibleRatio := float64(height) / float64(contentLines)
	visibleBlocks := int(visibleRatio * float64(height))
	if visibleBlocks < 1 {
		visibleBlocks = 1
	}

	totalBlocks := height
	startPos := int(scrollPercent * float64(totalBlocks-visibleBlocks))
	if startPos < 0 {
		startPos = 0
	}
	if startPos >= totalBlocks {
		startPos = totalBlocks - visibleBlocks
	}

	var bar strings.Builder
	for i := 0; i < totalBlocks; i++ {
		if i >= startPos && i < startPos+visibleBlocks {
			bar.WriteString("█")
		} else {
			bar.WriteString("░")
		}
		if i < totalBlocks-1 {
			bar.WriteString("\n")
		}
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Render(bar.String())
}

func (m Model) helpText() string {
	switch m.state {
	case stateChapters:
		return " ↑↓ navegar  ↵ abrir  / buscar  q sair "
	case stateTopics:
		return " ↑↓ navegar  esc voltar  / buscar  q sair "
	}
	return ""
}

func (m Model) footerView() string {
	mode := "Capítulos"
	if m.state == stateTopics {
		mode = "Tópicos"
	}

	info := fmt.Sprintf(" %s | q/ctrl+c sair", mode)
	if m.selectedTopic != "" {
		info = fmt.Sprintf(" %s | %s | q/ctrl+c sair", mode, m.selectedTopic)
	}

	rendered := statusBarStyle.Width(m.width).Render(info)
	return strings.TrimRight(rendered, " \n")
}
