// Help provides a comprehensive guide to the available flags and their descriptions
// for Chapter 8: "Agrupamento de Dados". This chapter covers topics related to arrays,
// slices, and maps, including their usage, operations, and specific scenarios. The
// function organizes the flags and their corresponding descriptions using the
// format.HelpMe structure and outputs the information using the format.PrintHelpMe
// function.
package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// help provides a list of available flags and their descriptions for Chapter 8:
// "Agrupamento de Dados". It displays topics related to arrays, slices, and maps,
// including their usage, operations, and specific scenarios. The function utilizes
// the format.HelpMe structure to organize the flags and their corresponding descriptions,
// and outputs the information using the format.PrintHelpMe function.
func Help() {
	hlp := []format.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array."},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta."},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range."},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia."},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia."},
		{Flag: "--slice-anexando-a-uma-slice", Description: "Apresenta o tópico Slice: anexando a uma slice."},
		{Flag: "--slice-make", Description: "Apresenta o tópico Slice: Make."},
		{Flag: "--slice-multi-dimensional", Description: "Apresenta o tópico Slice: Multi Dimensional."},
		{Flag: "--slice-a-surpresa-do-array-subjacente", Description: "Apresenta o tópico Slice: a surpresa do array subjacente."},
		{Flag: "--maps-introducao", Description: "Apresenta o tópico Maps: introdução."},
		{Flag: "--maps-range-e-deletando", Description: "Apresenta o tópico Maps: Range e Deletando."},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(hlp)
}
