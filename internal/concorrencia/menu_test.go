package concorrencia

import (
	"testing"
)

func TestMenuConcorrencia(t *testing.T) {
	t.Run("TestMenuLength", func(t *testing.T) {
		menuOptions := Menu([]string{
			flagConcorrenciaVsParalelismo,
			flagGoroutinesWaitgroups,
			flagDiscussaoCondicaoDeCorrida,
			flagCondicaoDeCorrida,
			flagMutex,
			flagAtomic,
		})

		if len(menuOptions) != 6 {
			t.Errorf("Expected menuOptions length of 6, got %v", len(menuOptions))
		}
	})

	t.Run("TestMenuOptions", func(t *testing.T) {
		menuOptions := Menu([]string{
			flagConcorrenciaVsParalelismo,
			flagGoroutinesWaitgroups,
			flagDiscussaoCondicaoDeCorrida,
			flagCondicaoDeCorrida,
			flagMutex,
			flagAtomic,
		})

		expectedOptions := []string{
			"--concorrencia-vs-paralelismo",
			"--goroutines-waitgroups",
			"--discussao-condicao-de-corrida",
			"--condicao-de-corrida",
			"--mutex",
			"--atomic",
		}

		for i, option := range menuOptions {
			if option.Options != expectedOptions[i] {
				t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
			}
		}
	})

	t.Run("TestMenuExecutionFunctions", func(t *testing.T) {
		menuOptions := Menu([]string{
			flagConcorrenciaVsParalelismo,
			flagGoroutinesWaitgroups,
			flagDiscussaoCondicaoDeCorrida,
			flagCondicaoDeCorrida,
			flagMutex,
			flagAtomic,
		})

		for i, option := range menuOptions {
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			} else {
				option.ExecFunc()
			}
		}
	})
}
