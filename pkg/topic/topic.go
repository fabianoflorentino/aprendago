// Package format provides utilities for formatting and processing sections of topics.
// It includes functions to iterate over a list of topics and execute specific sections
// within a given root directory.
package topic

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// ContentsProvider defines an interface for retrieving the contents of topics.
// Implementations of this interface should provide a method to fetch the contents
// of specified topics from a given root directory.
//
// TopicsContents takes the following parameters:
// - rootDir: The root directory where the topics are located.
// - topics: A slice of topic names for which the contents need to be retrieved.
type ContentsProvider interface {
	TopicsContents(rootDir string, topics []string)
	ListOfTopics(inputList []string, lenOfArray int) []string
	TopicConstructor(rootDir string, topicTitle string, topicList []string, topicInstance ContentsProvider)
}

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

// ListOfTopics creates and returns a new list of topics based on the provided input list.
// It initializes an output list with a specified capacity and appends all elements
// from the input list to the output list.
//
// Parameters:
//   - inputList: A slice of strings representing the input list of topics.
//   - lenOfArray: An integer specifying the initial capacity of the output list.
//
// Returns:
//   - A slice of strings containing the topics from the input list.
func (c *Contents) ListOfTopics(inputList []string, lenOfArray int) []string {
	validateList := validateLenOfArray(inputList, lenOfArray)

	outputList := make([]string, 0, len(validateList))
	outputList = append(outputList, validateList...)

	return outputList
}

// TopicConstructor initializes and processes the contents of a topic.
// It takes the following parameters:
// - rootDir: The root directory where the topic contents are located.
// - topicTitle: The title of the topic to be displayed.
// - topicList: A slice of strings representing the list of topics.
// - topicInstance: An instance of ContentsProvider used to handle topic-specific operations.
//
// The method prints the topic title, generates a list of topics using the provided topicList,
// and processes the topic contents by invoking the TopicsContents method of the topicInstance.
func (c *Contents) TopicConstructor(rootDir string, topicTitle string, topicList []string, topicInstance ContentsProvider) {
	fmt.Printf("\n%s\n", topicTitle)
	topicInstance.TopicsContents(rootDir, topicInstance.ListOfTopics(topicList, len(topicList)))
}

// validateLenOfArray trims or adjusts the length of the input slice to the specified length.
// If the specified length is less than or equal to zero, it returns an empty slice.
// If the specified length is greater than the length of the input slice, it returns the entire slice.
//
// Parameters:
//   - inputList: A slice of strings to be trimmed or adjusted.
//   - lenOfArray: The desired length of the resulting slice.
//
// Returns:
//
//	A slice of strings with a length equal to lenOfArray, or the entire input slice
//	if lenOfArray exceeds its length. If lenOfArray is less than or equal to zero, an empty slice is returned.
func validateLenOfArray(inputList []string, lenOfArray int) []string {
	if lenOfArray <= 0 {
		return []string{}
	}

	if lenOfArray > len(inputList) {
		lenOfArray = len(inputList)
	}

	return inputList[:lenOfArray]
}
