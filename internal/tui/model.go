package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// contentLoadedMsg is sent when a topic's content has been loaded asynchronously.
type contentLoadedMsg struct {
	content string
	err     error
}

// state represents the current navigation state.
type state int

const (
	stateChapters state = iota
	stateTopics
)

// Model is the main TUI model.
type Model struct {
	chapters []*chapter.Chapter

	state           state
	selectedChapter *chapter.Chapter
	selectedTopic   string

	chapterList list.Model
	topicList   list.Model
	contentPane viewport.Model

	width, height int
	ready         bool

	content string
	err     error

	quitting bool

	scrollX *int
}

func loadContentCmd(ch *chapter.Chapter, topic string) tea.Cmd {
	return func() tea.Msg {
		s, err := section.New(ch.RootDir)
		if err != nil {
			return contentLoadedMsg{err: err}
		}
		content, err := s.FormatToString(topic)
		if err != nil {
			return contentLoadedMsg{err: err}
		}
		return contentLoadedMsg{content: content}
	}
}

// NewModel creates and initialises the TUI model.
func NewModel(chapters []*chapter.Chapter) Model {
	delegate := newListDelegate()
	scrollX := 0

	chapterItems := make([]list.Item, len(chapters))
	for i, ch := range chapters {
		chapterItems[i] = chapterItem{ch: ch}
	}

	cl := list.New(chapterItems, delegate, 0, 0)
	cl.SetShowTitle(false)
	cl.SetShowStatusBar(false)
	cl.SetFilteringEnabled(true)
	cl.SetShowHelp(false)
	cl.SetShowPagination(true)
	cl.DisableQuitKeybindings()

	tl := list.New([]list.Item{}, delegate, 0, 0)
	tl.SetShowTitle(false)
	tl.SetShowStatusBar(false)
	tl.SetFilteringEnabled(true)
	tl.SetShowHelp(false)
	tl.SetShowPagination(true)
	tl.DisableQuitKeybindings()

	cp := viewport.New(0, 0)
	cp.MouseWheelEnabled = true

	return Model{
		chapters:    chapters,
		chapterList: cl,
		topicList:   tl,
		contentPane: cp,
		state:       stateChapters,
		scrollX:     &scrollX,
	}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Run starts the TUI with alternate screen and mouse support.
func Run(chapters []*chapter.Chapter) error {
	m := NewModel(chapters)
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err := p.Run()
	return err
}
