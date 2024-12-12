// Package canais provides functions to demonstrate and execute various concepts related to channels in Go.
// It includes sections on understanding channels, directional channels, range and close, select statement,
// comma ok idiom, convergence, divergence, and context. The package also provides menu options and help
// descriptions for each section.
package canais

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir is a constant that holds the relative path to the "canais" directory
// within the "internal" directory. This path is used as a base directory for
// various operations related to the "canais" module in the project.
const (
	rootDir = "internal/canais"
)

// Canais is a function that prints the title of Chapter 21: Channels
// and executes several sections related to the topic of channels in Go.
// The sections covered are:
// - Understanding Channels
// - Directional Channels & Using Channels
// - Range and Close
// - Select
// - The comma ok idiom
// - Convergence
// - Divergence
// - Context
func Canais() {
	fmt.Printf("\n\nCapítulo 21: Canais\n")

	executeSection("Entendendo Canais")
	executeSection("Canais direcionais & utilizando canais")
	executeSection("Range e Close")
	executeSection("Select")
	executeSection("A expressão comma ok")
	executeSection("Convergência")
	executeSection("Divergência")
	executeSection("Context")
}

// MenuCanais returns a slice of format.MenuOptions, each representing a menu option
// for different topics related to Go channels. Each menu option includes a string
// that represents the option and a function to execute when the option is selected.
// The topics covered include understanding channels, directional channels, range and close,
// select statement, the comma ok expression, convergence, divergence, and context.
func MenuCanais([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--entendendo-canais", ExecFunc: func() { executeSection("Entendendo Canais") }},
		{Options: "--canais-direcionais-utilizando-canais", ExecFunc: func() { executeSection("Canais direcionais & utilizando canais") }},
		{Options: "--range-e-close", ExecFunc: func() { executeSection("Range e Close") }},
		{Options: "--select", ExecFunc: func() { executeSection("Select") }},
		{Options: "--a-expressao-comma-ok", ExecFunc: func() { executeSection("A expressão comma ok") }},
		{Options: "--convergencia", ExecFunc: func() { executeSection("Convergência") }},
		{Options: "--convergencia --example", ExecFunc: func() { UsingConverge() }},
		{Options: "--convergencia --example --chan-string", ExecFunc: func() { UsingConvergeString() }},
		{Options: "--divergencia", ExecFunc: func() { executeSection("Divergência") }},
		{Options: "--context", ExecFunc: func() { executeSection("Context") }},
	}
}

// HelpMeCanais prints a list of topics related to channels in Go.
// It creates a slice of HelpMe structs, each containing a flag and a description
// of a specific topic about channels. The topics include understanding channels,
// directional channels, using channels, range and close, select statement,
// the comma ok idiom, convergence, divergence, and context.
// The function then prints the chapter title and calls PrintHelpMe to display the topics.
func HelpMeCanais() {
	hlp := []format.HelpMe{
		{Flag: "--entendendo-canais", Description: "Entendendo Canais"},
		{Flag: "--canais-direcionais-utilizando-canais", Description: "Canais direcionais & utilizando canais"},
		{Flag: "--range-e-close", Description: "Range e Close"},
		{Flag: "--select", Description: "Select"},
		{Flag: "--a-expressao-comma-ok", Description: "A expressão comma ok"},
		{Flag: "--convergencia", Description: "Convergência"},
		{Flag: "--convergencia --example", Description: "Convergência - Exemplo"},
		{Flag: "--convergencia --example --chan-string", Description: "Convergência de Strings - Exemplo"},
		{Flag: "--divergencia", Description: "Divergência"},
		{Flag: "--context", Description: "Context"},
	}

	fmt.Println("\nCapítulo 21: Canais")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
