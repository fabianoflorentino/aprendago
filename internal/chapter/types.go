package chapter

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// Chapter represents a chapter of the Aprenda Go course.
// It holds the chapter metadata and provides methods for navigating
// and displaying course content loaded from overview.yml files.
//
// For chapters that need custom behavior (e.g. interactive exercises,
// code examples), embed Chapter in a struct and override specific methods.
type Chapter struct {
	Number  int
	Title   string
	RootDir string
}

// Overview prints all topics in the chapter sequentially.
// It reads the topic list from the YAML and renders each one.
func (c *Chapter) Overview() error {
	topics, err := c.Topics()
	if err != nil {
		return fmt.Errorf("capitulo %d: erro ao listar topicos: %w", c.Number, err)
	}

	fmt.Printf("\n%s\n", c.Title)

	for _, topic := range topics {
		if err := c.ExecTopic(topic); err != nil {
			return err
		}
	}
	return nil
}

// ExecTopic renders a single topic/section by title.
// It uses the section package to read and format content from overview.yml.
func (c *Chapter) ExecTopic(topic string) error {
	s, err := section.New(c.RootDir)
	if err != nil {
		return fmt.Errorf("erro ao criar secao para '%s': %w", topic, err)
	}
	return s.Format(topic)
}

// Topics returns the list of section titles from the chapter's overview.yml.
func (c *Chapter) Topics() ([]string, error) {
	docs, err := reader.ReadOverview(c.RootDir)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler overview de '%s': %w", c.RootDir, err)
	}
	if len(docs) == 0 {
		return nil, nil
	}

	sections := docs[0].Description.Sections
	titles := make([]string, len(sections))
	for i, s := range sections {
		titles[i] = s.Title
	}
	return titles, nil
}


