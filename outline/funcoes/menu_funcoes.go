package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuFuncoes([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--sintaxe", ExecFunc: func() { Sintaxe() }},
		{Options: "--desenrolando-enumerando-uma-slice", ExecFunc: func() { DesenrolandoEnumerandoUmaSlice() }},
		{Options: "--defer", ExecFunc: func() { Defer() }},
	}
}
