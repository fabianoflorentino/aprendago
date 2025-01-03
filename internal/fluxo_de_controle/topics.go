// Package fluxo_de_controle provides functions and utilities to demonstrate
// and explain control flow concepts in Go. It includes examples and explanations
// for various control flow structures such as loops, conditionals, and logical operators.
package fluxo_de_controle

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the root directory path for the "fluxo_de_controle" module within the internal package.
const (
	rootDir = "internal/fluxo_de_controle"
)

// FluxoDeControle demonstrates various control flow constructs in Go.
// It prints section headers and executes corresponding sections to illustrate
// different control flow mechanisms such as loops, conditionals, and logical operators.
func FluxoDeControle() {
	fmt.Printf("\n\n06 - Fluxo de Controle\n")

	executeSection("Entendendo Fluxo de Controle")
	executeSection("Loops: Inicialização, Condição e Pós")
	executeSection("Loops: Nested Loop (repetição hierárquica)")
	executeSection("Loops: A Declaração for")
	executeSection("Loops: Break e Continue")
	executeSection("Loops: Utilizando ASCII")
	executeSection("Condicionais: A Declaração if")
	executeSection("Condicionais: A Declaração if else if else")
	executeSection("Condicionais: A Declaração switch")
	executeSection("Condicionais: A Declaração switch pt. 2 & documentação")
	executeSection("Operadores lógicos condicionais")
}

// MenuFluxoDeControle returns a slice of format.MenuOptions, each representing a menu option
// for different sections of control flow topics in Go. Each menu option has an associated
// execution function that calls executeSection with a specific section title.
func MenuFluxoDeControle([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--entendendo-fluxo-de-controle", ExecFunc: func() { executeSection("Entendendo Fluxo de Controle") }},
		{Options: "--loops-inicializacao-condicao-pos", ExecFunc: func() { executeSection("Loops: Inicialização, Condição e Pós") }},
		{Options: "--loops-nested-loop", ExecFunc: func() { executeSection("Loops: Nested Loop (repetição hierárquica)") }},
		{Options: "--loops-a-declaracao-for", ExecFunc: func() { executeSection("Loops: A Declaração for") }},
		{Options: "--loops-break-e-continue", ExecFunc: func() { executeSection("Loops: Break e Continue") }},
		{Options: "--loops-utilizando-ascii", ExecFunc: func() { executeSection("Loops: Utilizando ASCII") }},
		{Options: "--condicionais-a-declaracao-if", ExecFunc: func() { executeSection("Condicionais: A Declaração if") }},
		{Options: "--condicionais-if-else-if-else", ExecFunc: func() { executeSection("Condicionais: A Declaração if else if else") }},
		{Options: "--condicionais-a-declaracao-switch", ExecFunc: func() { executeSection("Condicionais: A Declaração switch") }},
		{Options: "--condicionais-a-declaracao-switch-pt2", ExecFunc: func() { executeSection("Condicionais: A Declaração switch pt. 2 & documentação") }},
		{Options: "--operadores-logicos-condicionais", ExecFunc: func() { executeSection("Operadores lógicos condicionais") }},
	}
}

// HelpMeFluxoDeControle provides a list of help topics related to control flow in Go.
// It includes explanations and details about various control flow concepts such as loops,
// conditional statements, and logical operators. The function prints a formatted help
// message to guide users through these topics.
func HelpMeFluxoDeControle() {
	hlp := []format.HelpMe{
		{Flag: "--entendendo-fluxo-de-controle", Description: "Explica o conceito de fluxo de controle em Go.", Width: 0},
		{Flag: "--loops-inicializacao-condicao-pos", Description: "Detalha o uso de loops com inicialização, condição e pós em Go.", Width: 0},
		{Flag: "--loops-nested-loop", Description: "Explora o conceito de loops aninhados em Go.", Width: 0},
		{Flag: "--loops-a-declaracao-for", Description: "Apresenta a declaração for em Go.", Width: 0},
		{Flag: "--loops-break-e-continue", Description: "Discute as instruções break e continue em loops em Go.", Width: 0},
		{Flag: "--loops-utilizando-ascii", Description: "Desafio surpresa: utilize ASCII para exibir texto em Go.", Width: 0},
		{Flag: "--loops-utilizando-ascii --resolucao", Description: "Desafio surpresa: utilize ASCII para exibir texto em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-if", Description: "Apresenta a declaração if em Go.", Width: 0},
		{Flag: "--condicionais-if-else-if-else", Description: "Detalha a declaração if-else-if-else em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-switch", Description: "Apresenta a declaração switch em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-switch-pt2", Description: "Detalha a declaração switch em Go.", Width: 0},
		{Flag: "--operadores-logicos-condicionais", Description: "Explora os operadores lógicos condicionais em Go.", Width: 0},
	}

	fmt.Println("\nCapípulo 6: Fluxo de Controle")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and applies the FormatSection function
// from the format package to the specified section within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
