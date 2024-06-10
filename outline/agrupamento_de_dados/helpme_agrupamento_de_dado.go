package agrupamento_de_dados

import (
	"fmt"

	helpme "github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeAgrupamentoDeDados() {
	hlp := []helpme.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array.", Width: 0},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta.", Width: 0},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	helpme.PrintHelpMe(hlp)
}
