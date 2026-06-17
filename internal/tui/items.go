package tui

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

type chapterItem struct {
	ch *chapter.Chapter
}

func (i chapterItem) Title() string       { return fmt.Sprintf("%d. %s", i.ch.Number, i.ch.Title) }
func (i chapterItem) Description() string  { return "" }
func (i chapterItem) FilterValue() string  { return i.Title() }

type topicItem struct {
	ch    *chapter.Chapter
	topic string
}

func (i topicItem) Title() string       { return i.topic }
func (i topicItem) Description() string  { return "" }
func (i topicItem) FilterValue() string  { return i.topic }
