// Package variaveis_valores_tipos provides functions and constants to demonstrate
// various concepts related to variables, values, and types in Go programming language.
// It includes sections on Go Playground, Hello World, short declaration operator,
// reserved word var, exploring types, zero value, fmt package, creating custom types,
// and type conversion (not coercion).
package variaveis_valores_tipos

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const (
	rootDir = "internal/variaveis_valores_tipos"
)

// VariaveisValoresETipos demonstrates various concepts related to variables, values, and types in Go.
// It prints section titles and executes corresponding sections to illustrate:
// - Go Playground
// - Hello, World!
// - Short variable declaration operator
// - The var keyword
// - Exploring types
// - Zero value
// - The fmt package
// - Creating your own type
// - Conversion, not coercion
func VariaveisValoresETipos() {
	fmt.Printf("\n\nCapítulo 02 - Variáveis, Valores & Tipos\n")

	executeSection("Go Playground")
	executeSection("Hello, World!")
	executeSection("Operador curto de declaração")
	executeSection("A palavra reservada var")
	executeSection("Explorando tipos")
	executeSection("Valor zero")
	executeSection("O pacote fmt")
	executeSection("Criando seu próprio tipo")
	executeSection("Conversão, não coerção")
}

// MenuVariaveisValoresTipos returns a slice of MenuOptions, each representing a different
// section of the "Variáveis, Valores e Tipos" topic. Each MenuOption contains an option
// string and an ExecFunc function that executes the corresponding section when called.
func MenuVariaveisValoresTipos([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--go-playground", ExecFunc: func() { executeSection("Go Playground") }},
		{Options: "--hello-world", ExecFunc: func() { executeSection("Hello, World!") }},
		{Options: "--operador-curto-de-declaracao", ExecFunc: func() { executeSection("Operador curto de declaração") }},
		{Options: "--a-palavra-reservada-var", ExecFunc: func() { executeSection("A palavra reservada var") }},
		{Options: "--explorando-tipos", ExecFunc: func() { executeSection("Explorando tipos") }},
		{Options: "--valor-zero", ExecFunc: func() { executeSection("Valor zero") }},
		{Options: "--o-pacote-fmt", ExecFunc: func() { executeSection("O pacote fmt") }},
		{Options: "--criando-seu-proprio-tipo", ExecFunc: func() { executeSection("Criando seu próprio tipo") }},
		{Options: "--conversao-nao-coercao", ExecFunc: func() { executeSection("Conversão, não coerção") }},
	}
}

// HelpMeVariaveisValoresTipos prints a list of help topics related to variables, values, and types in Go.
// Each topic includes a flag and a description in Portuguese.
// The function displays the chapter title and then prints the help topics using the format.PrintHelpMe function.
func HelpMeVariaveisValoresTipos() {
	hlp := []format.HelpMe{
		{Flag: "--go-playground", Description: "Exibe o Go Playground.", Width: 0},
		{Flag: "--hello-world", Description: "Exibe o Hello World.", Width: 0},
		{Flag: "--operador-curto-de-declaracao", Description: "Exibe o operador curto de declaração.", Width: 0},
		{Flag: "--a-palavra-reservada-var", Description: "Exibe a palavra reservada var.", Width: 0},
		{Flag: "--explorando-tipos", Description: "Exibe como explorar tipos.", Width: 0},
		{Flag: "--valor-zero", Description: "Exibe o valor zero.", Width: 0},
		{Flag: "--o-pacote-fmt", Description: "Exibe o pacote fmt.", Width: 0},
		{Flag: "--criando-seu-proprio-tipo", Description: "Exibe como criar seu próprio tipo.", Width: 0},
		{Flag: "--conversao-nao-coercao", Description: "Exibe a conversão não coerção.", Width: 0},
	}

	fmt.Println("\nCapítulo 2: Variáveis, Valores e Tipos")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to format the section within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
