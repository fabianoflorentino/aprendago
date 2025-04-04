package agrupamento_de_dados

import (
	"os"
	"testing"

	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
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
		menuOptions := Menu(
			[]string{
				flagArray,
				flagSliceLiteralComposta,
				flagSliceForRange,
				flagSliceFatiandoOuDeletando,
				flagSliceAnexando,
				flagSliceMake,
				flagSliceMultiDimensional,
				flagSliceSurpresaArraySubjacente,
				flagMapsIntroducao,
			})

		// Check if the menuOptions slice has the expected length.
		if len(menuOptions) != 10 {
			t.Errorf("Expected menuOptions length of 10, got %v", len(menuOptions))
		}

		// Check if the menuOptions slice has the expected options.
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

		for i, option := range menuOptions {
			if option.Options != expectedOptions[i] {
				t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
			}
		}
	})

	t.Run("TestMenuOptionsExecutionFunctions", func(t *testing.T) {
		// Create a slice of MenuOptions.
		menuOptions := Menu(
			[]string{
				flagArray,
				flagSliceLiteralComposta,
				flagSliceForRange,
				flagSliceFatiandoOuDeletando,
				flagSliceAnexando,
				flagSliceMake,
				flagSliceMultiDimensional,
				flagSliceSurpresaArraySubjacente,
				flagMapsIntroducao,
			})

		// Check if the menuOptions slice has the expected execution functions.
		for i, option := range menuOptions {
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		}

		m := base.New()
		newExecFunc := []func(){
			func() { m.SectionFormat(array, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceLiteralComposta, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceForRange, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceFatiandoOuDeletando, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceAnexando, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceMake, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceMultiDimensional, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(sliceSurpresaArraySubjacente, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(mapsIntroducao, m.SectionFactory(rootDir)) },
			func() { m.SectionFormat(mapsRangeEDeletando, m.SectionFactory(rootDir)) },
		}

		for i, execFunc := range newExecFunc {
			// Test if the execution functions behave as expected.
			if menuOptions[i].ExecFunc == nil || execFunc == nil {
				t.Errorf("Execution function for option %v is nil", i)
			} else {
				// Example: Compare the behavior of the functions by invoking them.
				// This assumes the functions have observable side effects or outputs.
				// You may need to adapt this part based on your actual function behavior.
				menuOptions[i].ExecFunc()
				execFunc()
			}
		}
	})
}
