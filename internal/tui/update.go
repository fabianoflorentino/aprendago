package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowSize(msg)
	case tea.KeyMsg:
		return m.handleKey(msg)
	case contentLoadedMsg:
		return m.handleContentLoaded(msg)
	}

	return m.muxListUpdate(msg)
}

func (m Model) handleWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	if msg.Width <= 0 || msg.Height <= 0 {
		return m, nil
	}
	m.width = msg.Width
	m.height = msg.Height
	m.ready = true

	mainH := m.height - 4
	leftTotal := int(float64(m.width) * 0.35)
	rightTotal := m.width - leftTotal - 1

	leftContentW := leftTotal - leftPanelFH
	leftContentH := mainH - leftPanelFV - 3 // -2 for help lines (wrapped), -1 for header
	if leftContentH < 1 {
		leftContentH = 1
	}
	rightContentW := rightTotal - rightPanelFH - 1 // -1 for scrollbar
	rightContentH := mainH - rightPanelFV
	if rightContentH < 1 {
		rightContentH = 1
	}

	if m.scrollX != nil && *m.scrollX >= leftContentW {
		*m.scrollX = leftContentW - 1
	}
	if m.scrollX != nil && *m.scrollX < 0 {
		*m.scrollX = 0
	}

	m.chapterList.SetSize(leftContentW, leftContentH)
	m.topicList.SetSize(leftContentW, leftContentH)
	m.contentPane.Width = rightContentW
	m.contentPane.Height = rightContentH

	if m.content != "" {
		wrapped := wrapContent(m.content, rightContentW)
		m.contentPane.SetContent(wrapped)
	}

	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		m.quitting = true
		return m, tea.Quit

	case "esc":
		if m.state == stateTopics {
			m.state = stateChapters
			m.selectedChapter = nil
			m.selectedTopic = ""
			m.content = ""
		}
		return m, nil

	case "enter":
		if m.state == stateChapters {
			return m.selectChapter()
		}
		return m, nil // topic content auto-loads on navigation

	case "backspace":
		m.contentPane.LineUp(3)
		return m, nil

	case "pgdown":
		m.contentPane.LineDown(3)
		return m, nil

	case "pgup":
		m.contentPane.LineUp(3)
		return m, nil

	case "left":
		if m.scrollX != nil && *m.scrollX > 0 {
			*m.scrollX--
		}
		return m, nil

	case "right":
		if m.scrollX != nil {
			*m.scrollX++
		}
		return m, nil
	}

	return m.muxListUpdate(msg)
}

func (m Model) selectChapter() (tea.Model, tea.Cmd) {
	item := m.chapterList.SelectedItem()
	if item == nil {
		return m, nil
	}

	ci, ok := item.(chapterItem)
	if !ok {
		return m, nil
	}

	ch := ci.ch
	topics, err := ch.Topics()
	if err != nil {
		m.err = err
		return m, nil
	}

	topicItems := make([]list.Item, len(topics))
	for i, t := range topics {
		topicItems[i] = topicItem{ch: ch, topic: t}
	}

	m.topicList.SetItems(topicItems)
	m.selectedChapter = ch
	m.state = stateTopics
	m.err = nil

	// Auto-load content for the first topic
	if len(topicItems) > 0 {
		if first, ok := topicItems[0].(topicItem); ok {
			m.selectedTopic = first.topic
			return m, loadContentCmd(first.ch, first.topic)
		}
	}

	return m, nil
}

func (m Model) handleContentLoaded(msg contentLoadedMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		m.err = msg.err
		return m, nil
	}
	m.content = msg.content
	wrapped := wrapContent(msg.content, m.contentPane.Width)
	m.contentPane.SetContent(wrapped)
	m.contentPane.GotoTop()
	m.err = nil
	return m, nil
}

func (m Model) muxListUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case stateChapters:
		var cmd tea.Cmd
		m.chapterList, cmd = m.chapterList.Update(msg)
		return m, cmd

	case stateTopics:
		prevSelected := m.topicList.SelectedItem()

		var cmd tea.Cmd
		m.topicList, cmd = m.topicList.Update(msg)

		currSelected := m.topicList.SelectedItem()
		if currSelected != nil && currSelected != prevSelected {
			if ti, ok := currSelected.(topicItem); ok {
				m.selectedTopic = ti.topic
				return m, tea.Batch(cmd, loadContentCmd(ti.ch, ti.topic))
			}
		}
		return m, cmd
	}

	return m, nil
}
