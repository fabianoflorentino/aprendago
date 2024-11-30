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
	executeSection("Comandos go")
	executeSection("Repositórios no github")
	executeSection("Explorando o github")
	executeSection("Compilação cruzada")
	executeSection("Pacotes")
}

func MenuSeuAmbienteDeDesenvolvimento([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--o-terminal", ExecFunc: func() { executeSection("O terminal") }},
		{Options: "--go-workspace-environment-variables", ExecFunc: func() { executeSection("Go workspace & environment variables") }},
		{Options: "--ides", ExecFunc: func() { executeSection("IDEs") }},
		{Options: "--comandos-go", ExecFunc: func() { executeSection("Comandos go") }},
		{Options: "--repositorios-no-github", ExecFunc: func() { executeSection("Repositórios no github") }},
		{Options: "--explorando-o-github", ExecFunc: func() { executeSection("Explorando o github") }},
		{Options: "--compilacao-cruzada", ExecFunc: func() { executeSection("Compilação cruzada") }},
		{Options: "--pacotes", ExecFunc: func() { executeSection("Pacotes") }},
	}
}

func HelpMeSeuAmbienteDeDesenvolvimento() {
	hlp := []format.HelpMe{
		{Flag: "--o-terminal", Description: "Exibe informações sobre o terminal."},
		{Flag: "--go-workspace-environment-variables", Description: "Exibe informações sobre o workspace e variáveis de ambiente do Go."},
		{Flag: "--ides", Description: "Exibe informações sobre IDEs."},
		{Flag: "--comandos-go", Description: "Exibe informações sobre comandos go."},
		{Flag: "--repositorios-no-github", Description: "Exibe informações sobre repositórios no github."},
		{Flag: "--explorando-o-github", Description: "Exibe informações sobre explorando o github."},
		{Flag: "--compilacao-cruzada", Description: "Exibe informações sobre compilação cruzada."},
		{Flag: "--pacotes", Description: "Exibe informações sobre pacotes."},
	}

	fmt.Println("\nCapítulo 19: Seu Ambiente de Desenvolvimento")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
