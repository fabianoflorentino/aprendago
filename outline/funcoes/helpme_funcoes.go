package funcoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeFuncoes() {
	hlp := []format.HelpMe{
		{Flag: "--sintaxe", Description: "Sintaxe", Width: 0},
	}

	fmt.Println("\nCapítulo 12: Funções")
	format.PrintHelpMe(hlp)
}
