// Package seu_ambiente_de_desenvolvimento provides functionalities to display and manage
// information related to the development environment setup. It includes sections on
// terminal usage, Go workspace and environment variables, IDEs, Go commands, GitHub
// repositories, cross-compilation, and packages. The package offers functions to display
// these sections, generate menu options, and provide help descriptions for each section.
package seu_ambiente_de_desenvolvimento

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the root directory path for the development environment configuration files.
const (
	rootDir = "internal/seu_ambiente_de_desenvolvimento"
)

// SeuAmbienteDeDesenvolvimento prints the title of Chapter 19 and executes a series of sections
// related to setting up and understanding your development environment in Go. The sections include
// topics such as using the terminal, Go workspace and environment variables, IDEs, Go commands,
// GitHub repositories, exploring GitHub, cross-compilation, and packages.
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

// MenuSeuAmbienteDeDesenvolvimento returns a slice of format.MenuOptions,
// each representing a different development environment topic with an associated
// execution function. The options include topics such as terminal usage, Go workspace
// and environment variables, IDEs, Go commands, GitHub repositories, exploring GitHub,
// cross-compilation, and packages. Each option is linked to a function that executes
// the corresponding section.
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

// HelpMeSeuAmbienteDeDesenvolvimento provides a list of help topics related to the development environment.
// It includes information about the terminal, Go workspace environment variables, IDEs, Go commands,
// GitHub repositories, exploring GitHub, cross-compilation, and packages.
// The function prints the chapter title and then displays the help topics using the format.PrintHelpMe function.
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

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to format the specified section within
// the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
