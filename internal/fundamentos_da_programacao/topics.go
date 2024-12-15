// Package fundamentos_da_programacao provides fundamental concepts of programming in Go.
// It includes various sections that cover topics such as boolean types, how computers work,
// numeric types, overflow, string types, numeric systems, constants, iota, and bit shifting.
// The package offers functions to display these topics, generate menu options, and provide help descriptions.
package fundamentos_da_programacao

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// foundational programming topics within the project structure.
const (
	rootDir = "internal/fundamentos_da_programacao"
)

// FundamentosDaProgramacao prints the title "Fundamentos da Programação" and
// executes a series of sections related to fundamental programming concepts.
// Each section is executed by calling the executeSection function with a
// specific topic name as an argument.
func FundamentosDaProgramacao() {
	fmt.Print("\n\n04 - Fundamentos da Programação\n")

	executeSection("Tipo booleano")
	executeSection("Como os computadores funcionam")
	executeSection("Tipos numéricos")
	executeSection("Overflow")
	executeSection("Tipo string")
	executeSection("Sistemas numéricos")
	executeSection("Constantes")
	executeSection("Iota")
	executeSection("Deslocamento de bits")
}

// MenuFundamentosDaProgramcao returns a slice of format.MenuOptions, each representing
// a different topic in the "Fundamentos da Programação" (Programming Fundamentals) section.
// Each menu option has an associated execution function that calls executeSection with
// the corresponding topic name.
func MenuFundamentosDaProgramcao([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--tipo-booleano", ExecFunc: func() { executeSection("Tipo booleano") }},
		{Options: "--como-os-computadores-funcionam", ExecFunc: func() { executeSection("Como os computadores funcionam") }},
		{Options: "--tipos-numericos", ExecFunc: func() { executeSection("Tipos numéricos") }},
		{Options: "--overflow", ExecFunc: func() { executeSection("Overflow") }},
		{Options: "--tipo-string", ExecFunc: func() { executeSection("Tipo string") }},
		{Options: "--sistemas-numericos", ExecFunc: func() { executeSection("Sistemas numéricos") }},
		{Options: "--constantes", ExecFunc: func() { executeSection("Constantes") }},
		{Options: "--iota", ExecFunc: func() { executeSection("Iota") }},
		{Options: "--deslocamento-de-bits", ExecFunc: func() { executeSection("Deslocamento de bits") }},
	}
}

// HelpMeFundamentosDaProgramacao provides a list of topics related to the fundamentals of programming in Go.
// Each topic is represented by a flag and a description, which explains the concept in detail.
func HelpMeFundamentosDaProgramacao() {
	hlp := []format.HelpMe{
		{Flag: "--tipo-booleano", Description: "Explora o tipo de dados booleano em Go.", Width: 0},
		{Flag: "--como-os-computadores-funcionam", Description: "Descreve o funcionamento dos computadores e sua importância para a programação.", Width: 0},
		{Flag: "--tipos-numericos", Description: "Explora os tipos numéricos em Go.", Width: 0},
		{Flag: "--overflow", Description: "Discute o conceito de overflow e como ele pode afetar seu código.", Width: 0},
		{Flag: "--tipo-string", Description: "Explora o tipo de dados string em Go.", Width: 0},
		{Flag: "--sistemas-numericos", Description: "Apresenta os sistemas numéricos e sua importância para a programação.", Width: 0},
		{Flag: "--constantes", Description: "Detalha o uso de constantes em Go.", Width: 0},
		{Flag: "--iota", Description: "Explora o uso do identificador iota em Go.", Width: 0},
		{Flag: "--deslocamento-de-bits", Description: "Discute o conceito de deslocamento de bits em Go.", Width: 0},
	}

	fmt.Println("\nCapítulo 4: Fundamentos da Programação")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the documentation.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply the necessary formatting to the section
// located in the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
