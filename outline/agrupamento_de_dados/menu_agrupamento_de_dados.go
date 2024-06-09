package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { Array() }},
	}
}
