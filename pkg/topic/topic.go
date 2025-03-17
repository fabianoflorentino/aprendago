// Package format provides utilities for formatting and processing sections of topics.
// It includes functions to iterate over a list of topics and execute specific sections
// within a given root directory.
package topic

import "github.com/fabianoflorentino/aprendago/pkg/format"

// Contents represents a collection of items or information related to a specific topic.
// It is currently an empty struct, but can be expanded in the future to include fields
// and methods that provide more functionality.
type Contents struct{}

// New creates and returns a new instance of Contents.
func New() *Contents {
	return &Contents{}
}

// TopicsContents processes a list of topics by formatting each topic's section
// within the specified root directory.
//
// Parameters:
//   - rootDir: The root directory where the topics' sections are located.
//   - topics: A slice of topic names to be processed.
func (c Contents) TopicsContents(rootDir string, topics []string) {
	for _, topic := range topics {
		format.FormatSection(rootDir, topic)
	}
}
