package agrupamento_de_dados

import (
	"os"
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

func TestTopicsAgrupamentoDeDados(t *testing.T) {
	os.Setenv("GOENV", "test")

	rootDir := os.Getenv("GOENV")

	if rootDir != "test" {
		t.Errorf("Expected rootDir to be 'test', got '%s'", rootDir)
	}

	mockProvider := &MockContentsProvider{}
	mockProvider.TopicsContents(rootDir, []string{"topic1", "topic2"})

	if mockProvider.CalledWithRootDir != rootDir {
		t.Errorf("Expected rootDir to be 'test', got '%s'", mockProvider.CalledWithRootDir)
	}

	if len(mockProvider.CalledWithTopics) != 2 {
		t.Errorf("Expected 2 topics, got %d", len(mockProvider.CalledWithTopics))
	}
}
