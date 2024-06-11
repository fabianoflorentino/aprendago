package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { Array() }},
		{Options: "--slice-literal-composta", ExecFunc: func() { SliceLiteralComposta() }},
		{Options: "--slice-for-range", ExecFunc: func() { SliceForRange() }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { SliceFatiandoOuDeletandoDeUmaFatia() }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", ExecFunc: func() { ResolucaoFatiaOuDeletandoDeUmaFatia() }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { SliceAnexandoASlice() }},
	}
}
