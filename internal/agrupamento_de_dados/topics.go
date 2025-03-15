/*
Package agrupamento_de_dados provides functionalities for demonstrating and executing
various data grouping topics in Go. It includes functions to print headers, execute
specific sections related to arrays, slices, and maps, and generate menu options and
help information for these topics. The package leverages the format package to format
and display the sections and help information.
*/
package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where data grouping topics are stored.
// This constant is used to reference the internal directory structure within the project.
const (
	rootDir = "internal/agrupamento_de_dados"
)

// AgrupamentoDeDados prints a header and executes a series of sections related to data grouping in Go.
// Each section is executed by calling the executeSection function with a specific topic name.
// The topics covered include arrays, slices (with various operations), and maps.
func AgrupamentoDeDados() {
	fmt.Printf("\n\n08 - Agrupamento de Dados\n")
	showTopics()
}

// MenuAgrupamentoDeDados returns a slice of format.MenuOptions, each representing a menu option
// for different data grouping topics in Go. Each menu option includes a command-line option
// string and an associated function to execute a specific section related to arrays, slices,
// and maps. The function does not take any parameters and returns a slice of format.MenuOptions.
func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { executeSection("Array") }},
		{Options: "--slice-literal-composta", ExecFunc: func() { executeSection("Slice: literal composta") }},
		{Options: "--slice-for-range", ExecFunc: func() { executeSection("Slice: for range") }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { executeSection("Slice: fatiando ou deletando de uma fatia") }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { executeSection("Slice: anexando a uma slice") }},
		{Options: "--slice-make", ExecFunc: func() { executeSection("Slice: make") }},
		{Options: "--slice-multi-dimensional", ExecFunc: func() { executeSection("Slice: multi dimensional") }},
		{Options: "--slice-a-surpresa-do-array-subjacente", ExecFunc: func() { executeSection("Slice: a surpresa do array subjacente") }},
		{Options: "--maps-introducao", ExecFunc: func() { executeSection("Maps: introdução") }},
		{Options: "--maps-range-e-deletando", ExecFunc: func() { executeSection("Maps: range e deletando") }},
	}
}

// HelpMeAgrupamentoDeDados prints a list of topics related to data grouping in Go.
// It creates a slice of HelpMe structs, each containing a flag and a description
// of a specific topic. The topics include arrays, slices (with various operations),
// and maps. The function then prints the chapter title and uses the PrintHelpMe
// function from the format package to display the list of topics.
func HelpMeAgrupamentoDeDados() {
	hlp := []format.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array.", Width: 0},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta.", Width: 0},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-anexando-a-uma-slice", Description: "Apresenta o tópico Slice: anexando a uma slice.", Width: 0},
		{Flag: "--slice-make", Description: "Apresenta o tópico Slice: Make.", Width: 0},
		{Flag: "--slice-multi-dimensional", Description: "Apresenta o tópico Slice: Multi Dimensional.", Width: 0},
		{Flag: "--slice-a-surpresa-do-array-subjacente", Description: "Apresenta o tópico Slice: a surpresa do array subjacente.", Width: 0},
		{Flag: "--maps-introducao", Description: "Apresenta o tópico Maps: introdução.", Width: 0},
		{Flag: "--maps-range-e-deletando", Description: "Apresenta o tópico Maps: Range e Deletando.", Width: 0},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(hlp)
}

// showTopics iterates over a list of topics and executes a section for each topic.
// It retrieves the list of topics from the listOfTopics function and processes each topic
// using the executeSection function.
func showTopics() {
	for _, topic := range listOfTopics() {
		executeSection(topic)
	}
}

// listOfTopics returns a list of topics related to Go programming.
// The topics include various aspects of arrays, slices, and maps.
func listOfTopics() []string {
	return []string{
		"Array",
		"Slice: literal composta",
		"Slice: for range",
		"Slice: fatiando ou deletando de uma fatia",
		"Slice: anexando a uma slice",
		"Slice: make",
		"Slice: multi dimensional",
		"Slice: a surpresa do array subjacente",
		"Maps: introdução",
		"Maps: range e deletando",
	}
}

// executeSection formats and processes a specific section of data.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to format the section within the root directory.
//
// Parameters:
//   - section: A string representing the name of the section to be formatted.
//
// Example usage:
//
//	executeSection("introduction")
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
