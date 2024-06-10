package agrupamento_de_dados

import (
	"fmt"

	helpme "github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeAgrupamentoDeDados() {
	hlp := []helpme.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array.", Width: 0},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	helpme.PrintHelpMe(hlp)
}
