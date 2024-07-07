package funcoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeFuncoes() {
	hlp := []format.HelpMe{
		{Flag: "--sintaxe", Description: "Sintaxe de uma função", Width: 0},
		{Flag: "--desenroland-enumerando-uma-slice", Description: "Descreve como desenrolar (enumerar) uma slice", Width: 0},
	}

	fmt.Println("\nCapítulo 12: Funções")
	format.PrintHelpMe(hlp)
}
