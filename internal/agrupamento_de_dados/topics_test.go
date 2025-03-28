package agrupamento_de_dados

import (
	"testing"
)

// MockContentsProvider is a mock implementation of the topic.ContentsProvider interface.
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
	contentsAgrupamentoDeDados(rootDir, mockContents)

	// Verify the root directory passed to TopicsContents
	if mockContents.CalledWithRootDir != rootDir {
		t.Errorf("Expected rootDir to be %s, got %s", rootDir, mockContents.CalledWithRootDir)
	}

	// Verify the topics passed to TopicsContents
	expectedTopics := listOfTopicsAgrupamentoDeDados()
	if len(mockContents.CalledWithTopics) != len(expectedTopics) {
		t.Errorf("Expected %d topics, got %d", len(expectedTopics), len(mockContents.CalledWithTopics))
	}

	for i, topic := range expectedTopics {
		if mockContents.CalledWithTopics[i] != topic {
			t.Errorf("Expected topic %s at index %d, got %s", topic, i, mockContents.CalledWithTopics[i])
		}
	}
}

func TestListOfTopicsAgrupamentoDeDados(t *testing.T) {
	expected := []string{
		Array,
		SliceLiteralComposta,
		SliceForRange,
		SliceFatiandoOuDeletando,
		SliceAnexando,
		SliceMake,
		SliceMultiDimensional,
		SliceSurpresaArraySubjacente,
		MapsIntroducao,
		MapsRangeEDeletando,
	}

	result := listOfTopicsAgrupamentoDeDados()

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
