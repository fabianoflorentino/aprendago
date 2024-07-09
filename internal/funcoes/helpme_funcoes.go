package funcoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func HelpMeFuncoes() {
	// hlp é uma slice de struct format.HelpMe que contém informações sobre comandos de ajuda.
	hlp := []format.HelpMe{
		{Flag: "--sintaxe", Description: "Sintaxe de declaração de função", Width: 0},
		{Flag: "--desenroland-enumerando-uma-slice", Description: "Descreve como iterar (enumerar) uma slice", Width: 0},
		{Flag: "--defer", Description: "Descreve o uso da declaração defer", Width: 0},
		{Flag: "--metodos", Description: "Descreve o que são métodos em Go", Width: 0},
		{Flag: "--interfaces-e-polimorfismo", Description: "Descreve o que são interfaces e polimorfismo em Go", Width: 0},
		{Flag: "--funcoes-anonimas", Description: "Descreve o que são funções anônimas em Go", Width: 0},
	}

	fmt.Println("\nCapítulo 12: Funções")
	format.PrintHelpMe(hlp)
}