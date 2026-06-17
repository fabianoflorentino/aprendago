package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/list"
)

type testItem struct {
	title string
}

func (i testItem) Title() string       { return i.title }
func (i testItem) Description() string  { return "" }
func (i testItem) FilterValue() string  { return i.title }

func TestNewListDelegate_LineCount(t *testing.T) {
	tests := []struct {
		name       string
		itemCount  int
		height     int
		width      int
		wantHeight int
	}{
		{name: "27 items at height 33", itemCount: 27, height: 33, width: 34, wantHeight: 33},
		{name: "27 items at height 17", itemCount: 27, height: 17, width: 22, wantHeight: 17},
		{name: "3 items at height 6", itemCount: 3, height: 6, width: 34, wantHeight: 6},
		{name: "50 items at height 10", itemCount: 50, height: 10, width: 30, wantHeight: 10},
		{name: "0 items at height 15", itemCount: 0, height: 15, width: 30, wantHeight: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items := make([]list.Item, tt.itemCount)
			for i := range items {
				items[i] = testItem{title: "Test"}
			}

			d := newListDelegate()
			l := list.New(items, d, tt.width, tt.height)
			l.SetShowTitle(false)
			l.SetFilteringEnabled(true)
			l.SetShowStatusBar(false)
			l.SetShowHelp(false)
			l.SetShowPagination(false)
			l.DisableQuitKeybindings()

			got := l.View()
			lineCount := strings.Count(got, "\n")
			if len(got) > 0 && got[len(got)-1] != '\n' {
				lineCount++
			}

			if lineCount != tt.wantHeight {
				t.Errorf("View() produced %d lines, want %d\nFull output:\n%q",
					lineCount, tt.wantHeight, got)
			}
		})
	}
}
