package aplicacoes

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/section"
)

func TestMenuAplicacoes(t *testing.T) {
	// Create a new section with a root directory.
	_, err := section.New(rootDir)
	if err != nil {
		t.Errorf("Error creating section: %v", err)
	}

	// Create a slice of MenuOptions for MenuAplicacoes.
	menuOptions := Menu(
		[]string{
			flagDocumentacaoJSON,
			flagDocumentacaoJSONExampleJSONMarshal,
			flagDocumentacaoJSONExampleJSONUnmarshal,
			flagDocumentacaoJSONExampleJSONEncoder,
			flagJSONMarshal,
			flagJSONUnmarshal,
			flagInterfaceWriter,
			flagPacoteSort,
			flagPacoteSortExample,
			flagCustomizandoSort,
			flagCustomizandoSortExample,
			flagBcrypt,
		})

	// Check if the menuOptions slice has the expected length.
	if len(menuOptions) != 12 {
		t.Errorf("Expected menuOptions length of 12, got %v", len(menuOptions))
	}

	// Check if the menuOptions slice has the expected options.
	expectedOptions := []string{
		"--documentacao-json",
		"--documentacao-json --example --json-marshal",
		"--documentacao-json --example --json-unmarshal",
		"--documentacao-json --example --json-encoder",
		"--json-marshal",
		"--json-unmarshal",
		"--interface-writer",
		"--pacote-sort",
		"--pacote-sort --example",
		"--customizando-sort",
		"--customizando-sort --example",
		"--bcrypt",
	}

	for i, option := range menuOptions {
		if option.Options != expectedOptions[i] {
			t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
		}
	}

	// Check if the menuOptions slice has the expected execution functions.
	for i, option := range menuOptions {
		if option.ExecFunc == nil {
			t.Errorf("Expected option %v to have a non-nil execution function", i)
		}
	}
}
