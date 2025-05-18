package aplicacoes

import (
	"os"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

func TestTopics(t *testing.T) {
	os.Setenv("GOENV", "test")

	tt := topic.New()

	list := []string{
		documentacaoJSON,
		jsonMarshal,
		jsonUnmarshal,
		interfaceWriter,
		pacoteSort,
		customizandoSort,
		bcrypt,
	}

	t.Run("TestRootDirValidation", func(t *testing.T) {
		if rootDir != "internal/aplicacoes" {
			t.Errorf("Expected rootDir to be 'internal/aplicacoes', got '%s'", rootDir)
		}
	})

	t.Run("TestListOfTopicsWithValidInput", func(t *testing.T) {
		expectedLength := 7
		topics := tt.ListOfTopics(list, expectedLength)
		if len(topics) != expectedLength {
			t.Errorf("Expected %d topics, got %d", expectedLength, len(topics))
		}

		expectedTopics := []string{
			documentacaoJSON,
			jsonMarshal,
			jsonUnmarshal,
			interfaceWriter,
			pacoteSort,
			customizandoSort,
			bcrypt,
		}

		for i, topic := range expectedTopics {
			if topics[i] != topic {
				t.Errorf("Expected topic at index %d to be %s, got %s", i, topic, topics[i])
			}
		}
	})

	t.Run("TestListOfTopicsWithLimitGreaterThanListSize", func(t *testing.T) {
		limit := 10
		topics := tt.ListOfTopics(list, limit)
		if len(topics) != len(list) {
			t.Errorf("Expected %d topics, got %d", len(list), len(topics))
		}
	})

	t.Run("TestListOfTopicsWithZeroLimit", func(t *testing.T) {
		limit := 0
		topics := tt.ListOfTopics(list, limit)
		if len(topics) != 0 {
			t.Errorf("Expected 0 topics, got %d", len(topics))
		}
	})

	t.Run("TestListOfTopicsWithNegativeLimit", func(t *testing.T) {
		limit := -1
		topics := tt.ListOfTopics(list, limit)
		if len(topics) != 0 {
			t.Errorf("Expected 0 topics for negative limit, got %d", len(topics))
		}
	})

	t.Run("TestListOfTopicsWithEmptyList", func(t *testing.T) {
		emptyList := []string{}
		topics := tt.ListOfTopics(emptyList, 5)
		if len(topics) != 0 {
			t.Errorf("Expected 0 topics for empty list, got %d", len(topics))
		}
	})
}
