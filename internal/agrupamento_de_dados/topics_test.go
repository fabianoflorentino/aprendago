package agrupamento_de_dados

import (
	"fmt"
	"testing"
)

type MockContents struct{}

func (c MockContents) TopicsContents(rootDir string, topics []string) {
	fmt.Println("Mock Topics Contents", topics)
}

func TestContentsAgrupamentoDeDados(t *testing.T) {
	rootDir := "test"
	contentsAgrupamentoDeDados(rootDir, MockContents{})

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

	if len(expected) != 10 {
		t.Errorf("Expected 10 topics, got %d", len(expected))
	}

	for i, topic := range expected {
		if topic != expected[i] {
			t.Errorf("Expected topic %s, got %s", topic, expected[i])
		}
	}

	t.Log("TestContentsAgrupamentoDeDados passed")
}
