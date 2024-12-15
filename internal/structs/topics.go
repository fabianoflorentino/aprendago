// Package structs provides functions to demonstrate and execute various sections
// related to structs in Go. It includes functions to print topics, generate menu
// options, and display help information for different struct-related sections.
package structs

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where internal structures are stored.
const (
	rootDir = "internal/structs"
)

// TopicStructs prints the chapter title for structs and executes sections related to structs.
// It covers the following sections:
// - Structs
// - Embedded Structs
// - Reading Documentation
// - Anonymous Structs
func TopicStructs() {
	fmt.Print("\n\nCapítulo 10 - Structs\n")

	executeSection("Structs")
	executeSection("Structs embutidos")
	executeSection("Lendo a documentação")
	executeSection("Structs anônimos")
}

// MenuStructs returns a slice of MenuOptions, each representing a menu item
// related to different topics about structs in Go. Each menu item has an option
// string and an associated function to execute when the option is selected.
func MenuStructs([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--structs", ExecFunc: func() { executeSection("Structs") }},
		{Options: "--structs-embutidos", ExecFunc: func() { executeSection("Structs embutidos") }},
		{Options: "--lendo-a-documentacao", ExecFunc: func() { executeSection("Lendo a documentação") }},
		{Options: "--structs-anonimos", ExecFunc: func() { executeSection("Structs anônimos") }},
	}
}

// HelpMeStructs prints a help message for various struct-related topics.
// It defines a slice of HelpMe structs with flags and descriptions for each topic,
// and then prints the chapter title followed by the formatted help message.
func HelpMeStructs() {
	hlp := []format.HelpMe{
		{Flag: "--structs", Description: "Structs", Width: 0},
		{Flag: "--structs-embutidos", Description: "Structs Embutidos", Width: 0},
		{Flag: "--lendo-a-documentacao", Description: "Lendo a documentação", Width: 0},
		{Flag: "--structs-anonimos", Description: "Structs Anônimos", Width: 0},
	}

	fmt.Println("\n Capítulo 10: Structs")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to format the section within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
