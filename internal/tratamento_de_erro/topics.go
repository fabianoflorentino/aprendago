// Package tratamento_de_erro provides functions and utilities for error handling in Go.
// It includes examples and sections on understanding errors, verifying errors, printing and logging errors,
// recovering from errors, and adding additional information to errors.
package tratamento_de_erro

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where error handling topics are stored within the project.
const (
	rootDir = "internal/tratamento_de_erro"
)

// TratamentoDeErro demonstrates various error handling techniques in Go.
// It prints the chapter title and executes sections covering:
// - Understanding errors
// - Checking errors
// - Print & log
// - Recover
// - Errors with additional information
func TratamentoDeErro() {
	fmt.Print("\n\nCapítulo 23 - Tratamento de Erro\n")

	executeSection("Entendendo erros")
	executeSection("Verificando erros")
	executeSection("Print & log")
	executeSection("Recover")
	executeSection("Erros com informações adicionais")
}

// MenuTratamentoDeErro returns a slice of format.MenuOptions, each representing a menu option
// for error handling topics. Each menu option includes a description and a corresponding
// function to execute when the option is selected.
//
// The available menu options are:
// - "--entendendo-erro": Executes the "Entendendo erros" section.
// - "--verificando-erros": Executes the "Verificando erros" section.
// - "--print-log": Executes the "Print & log" section.
// - "--recover": Executes the "Recover" section.
// - "--erros-com-informacoes-adicionais": Executes the "Erros com informações adicionais" section.
func MenuTratamentoDeErro([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--entendendo-erro", ExecFunc: func() { executeSection("Entendendo erros") }},
		{Options: "--verificando-erros", ExecFunc: func() { executeSection("Verificando erros") }},
		{Options: "--print-log", ExecFunc: func() { executeSection("Print & log") }},
		{Options: "--recover", ExecFunc: func() { executeSection("Recover") }},
		{Options: "--erros-com-informacoes-adicionais", ExecFunc: func() { executeSection("Erros com informações adicionais") }},
	}
}

// HelpMeTratamentoDeErro provides a list of help topics related to error handling in Go.
// It prints a chapter title and then displays a formatted list of help topics with their descriptions.
// The topics include understanding errors, checking errors, printing and logging, recovering from panics,
// and errors with additional information.
func HelpMeTratamentoDeErro() {
	hlp := []format.HelpMe{
		{Flag: "--entendendo-erro", Description: "Entendendo erros"},
		{Flag: "--verificando-erros", Description: "Verificando erros"},
		{Flag: "--print-log", Description: "Print & log"},
		{Flag: "--recover", Description: "Recover"},
		{Flag: "--erros-com-informacoes-adicionais", Description: "Erros com informações adicionais"},
	}

	fmt.Printf("\nCapítulo 23: Tratamento de Erro\n")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and applies the FormatSection function
// from the format package to the specified section within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
