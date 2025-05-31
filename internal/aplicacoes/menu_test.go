package aplicacoes

import (
	"os"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func TestMenuAplicacoes(t *testing.T) {
	var err error

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("failed to open null device: %v", err)
	}

	defer nullFile.Close()

	os.Stdout = nullFile

	t.Run("TestMenuOptionsLengthAndContent", func(t *testing.T) {
		// Create a slice of MenuOptions.
		mo := menuOptions(topics)

		// Check if the menuOptions slice has the expected length.
		if len(mo) != 12 {
			t.Errorf("Expected menuOptions length of 12, got %v", len(mo))
		}
	})

	t.Run("TestMenuOptionsContent", func(t *testing.T) {
		mo := menuOptions(topics)

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

		for i, option := range mo {
			if option.Options != expectedOptions[i] {
				t.Errorf("Expected %s, got %s", expectedOptions[i], option.Options)
			}
		}
	})

	t.Run("TestMenuOptionsExecFunc", func(t *testing.T) {
		mo := menuOptions(topics)

		if len(mo) != 12 {
			t.Fatalf("Expected menuOptions length of 12, got %v", len(mo))
		}

		for i, option := range mo {
			if option.ExecFunc == nil {
				t.Errorf("Expected ExecFunc for option %v to be non-nil", i)
			}
		}
	})
}

func menuOptions(listMenu []string) []format.MenuOptions {
	return Menu(listMenu)
}
