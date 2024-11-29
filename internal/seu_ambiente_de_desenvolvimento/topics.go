package seu_ambiente_de_desenvolvimento

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

var rootDir = "internal/seu_ambiente_de_desenvolvimento"

func SeuAmbienteDeDesenvolvimento() {
	fmt.Print("\n\nCapítulo 19: Seu Ambiente de Desenvolvimento\n")

	executeSection("O terminal")
	executeSection("Go workspace & environment variables")
	executeSection("IDEs")
}

func MenuSeuAmbienteDeDesenvolvimento([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--o-terminal", ExecFunc: func() { executeSection("O terminal") }},
		{Options: "--go-workspace-environment-variables", ExecFunc: func() { executeSection("Go workspace & environment variables") }},
		{Options: "--ides", ExecFunc: func() { executeSection("IDEs") }},
	}
}

func HelpMeSeuAmbienteDeDesenvolvimento() {
	hlp := []format.HelpMe{
		{Flag: "--o-terminal", Description: "Exibe informações sobre o terminal."},
		{Flag: "--go-workspace-environment-variables", Description: "Exibe informações sobre o workspace e variáveis de ambiente do Go."},
		{Flag: "--ides", Description: "Exibe informações sobre IDEs."},
	}

	fmt.Println("\nCapítulo 19: Seu Ambiente de Desenvolvimento")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
