package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuFuncoes([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--sintaxe", ExecFunc: func() { Sintaxe() }},
	}
}
