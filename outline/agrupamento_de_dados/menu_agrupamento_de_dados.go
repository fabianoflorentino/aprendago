package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { Array() }},
		{Options: "--slice-literal-composta", ExecFunc: func() { SliceLiteralComposta() }},
		{Options: "--slice-for-range", ExecFunc: func() { SliceForRange() }},
	}
}
