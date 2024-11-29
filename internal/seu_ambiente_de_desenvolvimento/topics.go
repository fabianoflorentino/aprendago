package seu_ambiente_de_desenvolvimento

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

var rootDir = "internal/seu_ambiente_de_desenvolvimento"

func SeuAmbienteDeDesenvolvimento() {
	fmt.Print("\n\nCapítulo 19: Seu Ambiente de Desenvolvimento\n")

	executeSection("O terminal")
}

func MenuSeuAmbienteDeDesenvolvimento([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--o-terminal", ExecFunc: func() { executeSection("O terminal") }},
	}
}

func HelpMeSeuAmbienteDeDesenvolvimento() {
	hlp := []format.HelpMe{
		{Flag: "--o-terminal", Description: "Exibe informações sobre o terminal."},
	}

	fmt.Println("\nCapítulo 19: Seu Ambiente de Desenvolvimento")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
