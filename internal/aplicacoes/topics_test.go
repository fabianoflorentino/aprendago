// Package aplicacoes provides functionalities and tests related to various application topics.
// It includes utilities for handling JSON operations, sorting custom data, working with interfaces,
// and encrypting data using bcrypt. The package also contains tests to ensure the correctness
// of these functionalities.
package aplicacoes

import (
	"testing"
)

type MockContentsProvider struct {
	CalledWithRootDir string
	CalledWithTopics  []string
}

func (m *MockContentsProvider) TopicsContents(rootDir string, topics []string) {
	m.CalledWithRootDir = rootDir
	m.CalledWithTopics = topics
}

func TestTopics(t *testing.T) {
	// Mock the root directory and contents provider
	rootDir := "mockRootDir"
	mockContents := &MockContentsProvider{}

	// Call the function under test
	contentsAplicacoes(rootDir, mockContents)

	// Verify the root directory passed to TopicsContents
	if mockContents.CalledWithRootDir != rootDir {
		t.Errorf("Expected rootDir to be %s, got %s", rootDir, mockContents.CalledWithRootDir)
	}

	// Verify the topics passed to TopicsContents
	expectedTopics := listOfTopicsAplicacoes()
	if len(mockContents.CalledWithTopics) != len(expectedTopics) {
		t.Errorf("Expected %d topics, got %d", len(expectedTopics), len(mockContents.CalledWithTopics))
	}

	for i, topic := range expectedTopics {
		if mockContents.CalledWithTopics[i] != topic {
			t.Errorf("Expected topic %s at index %d, got %s", topic, i, mockContents.CalledWithTopics[i])
		}
	}
}

func TestListOfTopicsAplicacoes(t *testing.T) {
	expected := []string{
		DocumentacaoJSON,
		JSONMarshal,
		JSONUnmarshal,
		InterfaceWriter,
		PacoteSort,
		CustomizandoSort,
		Bcrypt,
	}

	result := listOfTopicsAplicacoes()

	if len(result) != len(expected) {
		t.Errorf("Expected %d topics, got %d", len(expected), len(result))
	}

	for i, topic := range expected {
		if result[i] != topic {
			t.Errorf("Expected topic %s at index %d, got %s", topic, i, result[i])
		}
	}
}

func TestListOfTopics(t *testing.T) {
	input := []string{"Topic1", "Topic2", "Topic3"}
	expected := []string{"Topic1", "Topic2", "Topic3"}

	result := listOfTopics(input)

	if len(result) != len(expected) {
		t.Errorf("Expected %d topics, got %d", len(expected), len(result))
	}

	for i, topic := range expected {
		if result[i] != topic {
			t.Errorf("Expected topic %s at index %d, got %s", topic, i, result[i])
		}
	}
}
