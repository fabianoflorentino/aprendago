// Package ponteiros provides functions and utilities to understand and work with pointers in Go.
// It includes sections on what pointers are, when to use them, and provides a menu and help options
// for navigating through the different topics related to pointers.
package ponteiros

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where the pointer-related topics are stored within the project.
const (
	rootDir = "internal/ponteiros"
)

// Ponteiros prints the chapter title and executes sections related to pointers.
// It covers topics such as what pointers are and when to use them.
func Ponteiros() {
	fmt.Printf("\n\nCapítulo 14: Ponteiros\n")

	executeSection("O que são ponteiros?")
	executeSection("Quando usar ponteiros")
}

// MenuPonteiros returns a slice of MenuOptions for the "ponteiros" topic.
// Each MenuOption contains an option string and an associated execution function.
// The options provided are:
// --o-que-sao-ponteiros: Executes the section explaining what pointers are.
// --quando-usar-ponteiros: Executes the section explaining when to use pointers.
func MenuPonteiros([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--o-que-sao-ponteiros", ExecFunc: func() { executeSection("O que são ponteiros?") }},
		{Options: "--quando-usar-ponteiros", ExecFunc: func() { executeSection("Quando usar ponteiros") }},
	}
}

// HelpMePonteiros prints a help message for the "Ponteiros" (Pointers) chapter.
// It provides descriptions of what pointers are in Go and when to use them.
func HelpMePonteiros() {
	hlp := []format.HelpMe{
		{Flag: "--o-que-sao-ponteiros", Description: "Descreve o que são ponteiros em Go", Width: 0},
		{Flag: "--quando-usar-ponteiros", Description: "Descreve quando usar ponteiros em Go", Width: 0},
	}

	fmt.Println("\nCapítulo 14: Ponteiros")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a specific section of the documentation.
// It takes a section name as a string parameter and uses the FormatSection function
// from the format package to apply the formatting to the section located in the rootDir.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
