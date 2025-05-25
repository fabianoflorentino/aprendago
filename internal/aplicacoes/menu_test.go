package aplicacoes

import (
	"os"
	"testing"

	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
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

	t.Run("TestMenuOptionsExecutionFunctions", func(t *testing.T) {
		mo := menuOptions(topics)

		for _, option := range mo {
			if option.ExecFunc == nil {
				t.Errorf("Expected execution function for %s to be not nil", option.Options)
			}
		}
	})

	t.Run("TestMenuOptionsExecution", func(t *testing.T) {
		mo := menuOptions(topics)

		m := base.New()
		newExecFunc := []func(){
			func() { m.SectionFormat(documentacaoJSON, m.SectionFactory(rootDir)) },
			func() { UsingJsonMarshal() },
			func() { UsingJsonUnmarshal() },
			func() { UsingJsonEncoder() },
			func() { m.SectionFormat(jsonMarshal, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(jsonUnmarshal, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(interfaceWriter, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(pacoteSort, m.SectionFactory(rootDir)) },
			func() { UsingPackageSort() },
			func() { m.SectionFormat(customizandoSort, m.SectionFactory(rootDir)) },
			func() { UsingCustomSort() },
			func() { m.SectionFormat(bcrypt, m.SectionFactory(rootDir)) },
		}

		for i, execFunc := range newExecFunc {
			if mo[i].ExecFunc == nil || execFunc == nil {
				t.Errorf("Execution function for option %v is nil", i)
			} else {
				mo[i].ExecFunc()
				execFunc()
			}
		}
	})
}

func menuOptions(listMenu []string) []format.MenuOptions {
	return Menu(listMenu)
}
