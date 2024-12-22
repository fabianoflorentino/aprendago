// Package documentacao provides functionalities to handle documentation topics
// for the "aprendago" project. It includes functions to display documentation
// sections, generate menu options for documentation topics, and print help
// information related to documentation.
package documentacao

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the root directory for documentation files within the project.
const (
	rootDir = "internal/documentacao"
)

// Package documentacao provides functions to execute and display various sections
// related to Go documentation. It includes topics such as introduction to documentation,
// usage of go doc, godoc, and pkg.go.dev, as well as writing documentation.
func Documentacao() {
	fmt.Printf("\n\nCapítulo 25: Documentação\n")

	executeSections("Introdução")
	executeSections("go doc")
	executeSections("godoc")
	executeSections("https://pkg.go.dev/")
	executeSections("Escrevendo documentação")
}

// Package documentacao provides functionalities related to documentation topics in Go.
// It includes a menu with various documentation options and their respective execution functions.

// MenuDocumentacao returns a slice of MenuOptions containing different documentation topics
// and their associated execution functions. Each option is represented by a string and
// an ExecFunc that executes the corresponding section.
func MenuDocumentacao([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--introducao", ExecFunc: func() { executeSections("Introdução") }},
		{Options: "--go-doc", ExecFunc: func() { executeSections("go doc") }},
		{Options: "--godoc", ExecFunc: func() { executeSections("godoc") }},
		{Options: "--pkg-go-dev", ExecFunc: func() { executeSections("https://pkg.go.dev/") }},
		{Options: "--escrevendo-documentacao", ExecFunc: func() { executeSections("Escrevendo documentação") }},
	}
}

// HelpMeDocumentacao provides a list of documentation topics and their descriptions.
// It prints a formatted help message for Chapter 25: Documentation.
//
// The topics include:
// - Introdução: Introduction to the topic
// - go doc: Information about the 'go doc' command
// - godoc: Information about the 'godoc' command
// - https://pkg.go.dev/: Link to the Go package documentation website
// - Escrevendo documentação: Tips on writing documentation
//
// This function is part of the 'documentacao' package, which contains utilities and helpers
// for generating and managing documentation in Go projects.
func HelpMeDocumentacao() {
	hlp := []format.HelpMe{
		{Flag: "--introducao", Description: "Introdução"},
		{Flag: "--go-doc", Description: "go doc"},
		{Flag: "--godoc", Description: "godoc"},
		{Flag: "--pkg-go-dev", Description: "https://pkg.go.dev/"},
		{Flag: "--escrevendo-documentacao", Description: "Escrevendo documentação"},
	}

	fmt.Printf("\nCapítulo 25: Documentação")
	format.PrintHelpMe(hlp)
}

// executeSections formats and processes a given section of documentation.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting to the section within the root directory.
func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
