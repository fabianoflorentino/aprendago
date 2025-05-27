package agrupamento_de_dados

import (
	"os"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func TestMenu(t *testing.T) {
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
		if len(mo) != 10 {
			t.Errorf("Expected menuOptions length of 10, got %v", len(mo))
		}
	})

	t.Run("TestMenuOptionsContent", func(t *testing.T) {
		mo := menuOptions(topics)

		expectedOptions := []string{
			"--array",
			"--slice-literal-composta",
			"--slice-for-range",
			"--slice-fatiando-ou-deletando-de-uma-fatia",
			"--slice-anexando-a-uma-slice",
			"--slice-make",
			"--slice-multi-dimensional",
			"--slice-a-surpresa-do-array-subjacente",
			"--maps-introducao",
			"--maps-range-e-deletando",
		}

		for i, option := range mo {
			if option.Options != expectedOptions[i] {
				t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
			}
		}
	})

	t.Run("TestMenuOptionsExecFunc", func(t *testing.T) {
		mo := menuOptions(topics)

		if len(mo) != 10 {
			t.Fatalf("Expected menuOptions length of 10, got %v", len(mo))
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
