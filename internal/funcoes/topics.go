// Package funcoes provides various examples and explanations of functions in Go.
// It covers topics such as function syntax, slices, defer, methods, interfaces,
// anonymous functions, function expressions, returning functions, callbacks,
// closures, and recursion. This package is part of the "aprendago" project,
// which aims to teach Go programming concepts.
package funcoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the internal functions directory
// within the project. It is used to reference the location of function-related
// files and resources.
const (
	rootDir = "internal/funcoes"
)

// Funcoes prints the chapter title and executes a series of sections related to functions in Go.
// Each section is executed by calling the executeSection function with the section name as an argument.
func Funcoes() {
	fmt.Printf("\n\nCapítulo 12: Funções\n")

	executeSection("Síntaxe")
	executeSection("Desenrolando (enumerando) uma slice")
	executeSection("Defer")
	executeSection("Métodos")
	executeSection("Interfaces e Polimorfismo")
	executeSection("Funções anônimas")
	executeSection("Funções como expressões")
	executeSection("Retornando uma função")
	executeSection("Callback")
	executeSection("Closure")
	executeSection("Recursividade")
}

// MenuFuncoes returns a slice of format.MenuOptions, each representing a menu option
// for different sections related to Go functions. Each menu option includes an
// option string and an associated function to execute the corresponding section.
// The sections covered include syntax, enumerating slices, defer, methods, interfaces
// and polymorphism, anonymous functions, functions as expressions, returning functions,
// callbacks, closure, and recursion.
func MenuFuncoes([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--sintaxe", ExecFunc: func() { executeSection("Síntaxe") }},
		{Options: "--desenroland-enumerando-uma-slice", ExecFunc: func() { executeSection("Desenrolando (enumerando) uma slice") }},
		{Options: "--defer", ExecFunc: func() { executeSection("Defer") }},
		{Options: "--metodos", ExecFunc: func() { executeSection("Métodos") }},
		{Options: "--interfaces-e-polimorfismo", ExecFunc: func() { executeSection("Interfaces e Polimorfismo") }},
		{Options: "--funcoes-anonimas", ExecFunc: func() { executeSection("Funções anônimas") }},
		{Options: "--func-como-expressa", ExecFunc: func() { executeSection("Funções como expressões") }},
		{Options: "--retornando-uma-funcao", ExecFunc: func() { executeSection("Retornando uma função") }},
		{Options: "--callback", ExecFunc: func() { executeSection("Callback") }},
		{Options: "--resolucao-desafio-callback", ExecFunc: func() { ResolucaoDesafioCallback() }},
		{Options: "--closure", ExecFunc: func() { executeSection("Closure") }},
		{Options: "--recursividade", ExecFunc: func() { executeSection("Recursividade") }},
	}
}

// HelpMeFuncoes provides a list of help topics related to functions in Go.
// Each topic includes a flag and a description explaining various concepts
// such as function syntax, iterating over slices, defer statement, methods,
// interfaces and polymorphism, anonymous functions, function expressions,
// returning functions, callbacks, closures, and recursion.
func HelpMeFuncoes() {
	hlp := []format.HelpMe{
		{Flag: "--sintaxe", Description: "Sintaxe de declaração de função", Width: 0},
		{Flag: "--desenroland-enumerando-uma-slice", Description: "Descreve como iterar (enumerar) uma slice", Width: 0},
		{Flag: "--defer", Description: "Descreve o uso da declaração defer", Width: 0},
		{Flag: "--metodos", Description: "Descreve o que são métodos em Go", Width: 0},
		{Flag: "--interfaces-e-polimorfismo", Description: "Descreve o que são interfaces e polimorfismo em Go", Width: 0},
		{Flag: "--funcoes-anonimas", Description: "Descreve o que são funções anônimas em Go", Width: 0},
		{Flag: "--func-como-expressa", Description: "Descreve como declarar funções como expressões", Width: 0},
		{Flag: "--retornando-uma-funcao", Description: "Descreve como retornar uma função em Go", Width: 0},
		{Flag: "--callback", Description: "Descreve o que é um callback em Go", Width: 0},
		{Flag: "--resolucao-desafio-callback", Description: "Mostra a resolução do desafio de callback", Width: 0},
		{Flag: "--closure", Description: "Descreve o que é um closure em Go", Width: 0},
		{Flag: "--recursividade", Description: "Descreve o que é recursividade em Go", Width: 0},
	}

	fmt.Println("\nCapítulo 12: Funções")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
