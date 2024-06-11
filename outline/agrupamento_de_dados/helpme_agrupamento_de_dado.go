package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeAgrupamentoDeDados() {
	hlp := []format.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array.", Width: 0},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta.", Width: 0},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-anexando-a-uma-slice", Description: "Apresenta o tópico Slice: anexando a uma slice.", Width: 0},
		{Flag: "--slice-make", Description: "Apresenta o tópico Slice: Make.", Width: 0},
		{Flag: "--slice-multi-dimensional", Description: "Apresenta o tópico Slice: Multi Dimensional.", Width: 0},
		{Flag: "--slice-a-surpresa-do-array-subjacente", Description: "Apresenta o tópico Slice: a surpresa do array subjacente.", Width: 0},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(hlp)
}
