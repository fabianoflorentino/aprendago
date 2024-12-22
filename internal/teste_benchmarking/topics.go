// Package teste_benchmarking provides functionalities to demonstrate and execute
// various testing and benchmarking techniques in Go. It includes sections on
// introduction to testing, table-driven tests, example-based tests, code formatting
// and linting tools (go fmt, govet, golint), benchmarking, and code coverage.
// The package also offers menu options and help descriptions for each section.
package teste_benchmarking

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where benchmarking tests are located.
const (
	rootDir = "internal/teste_benchmarking"
)

// TesteEBenchmarking prints the title of Chapter 27 and executes a series of sections
// related to testing and benchmarking in Go. Each section is executed by calling the
// executeSections function with the section name as an argument.
func TesteEBenchmarking() {
	fmt.Printf("\n\nCapítulo 27: Teste e Benchmarking\n")

	executeSections("Introdução")
	executeSections("Testes em tabelas")
	executeSections("Testes como exemplo")
	executeSections("go fmt, govet e golint")
	executeSections("Benchmark")
	executeSections("Cobertura")
}

// MenuTesteEBenchmarking returns a slice of MenuOptions, each representing a different
// testing and benchmarking topic. Each MenuOption contains an option string and an
// associated function to execute the corresponding section. The available options are:
// --introducao, --testes-em-tabelas, --testes-como-exemplo, --go-fmt-govet-golint,
// --benchmark, and --cobertura.
func MenuTesteEBenchmarking([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--introducao", ExecFunc: func() { executeSections("Introdução") }},
		{Options: "--testes-em-tabelas", ExecFunc: func() { executeSections("Testes em tabelas") }},
		{Options: "--testes-como-exemplo", ExecFunc: func() { executeSections("Testes como exemplo") }},
		{Options: "--go-fmt-govet-golint", ExecFunc: func() { executeSections("go fmt, govet e golint") }},
		{Options: "--benchmark", ExecFunc: func() { executeSections("Benchmark") }},
		{Options: "--cobertura", ExecFunc: func() { executeSections("Cobertura") }},
	}
}

// HelpMeTesteEBenchmarking prints a help message for chapter 27: Testing and Benchmarking.
// It provides a list of flags and their descriptions related to various testing and benchmarking topics.
func HelpMeTesteEBenchmarking() {
	hlp := []format.HelpMe{
		{Flag: "--introducao", Description: "Introdução"},
		{Flag: "--testes-em-tabelas", Description: "Testes em tabelas"},
		{Flag: "--testes-como-exemplo", Description: "Testes como exemplo"},
		{Flag: "--go-fmt-govet-golint", Description: "go fmt, govet e golint"},
		{Flag: "--benchmark", Description: "Benchmark"},
		{Flag: "--cobertura", Description: "Cobertura"},
	}

	fmt.Printf("\nCapítulo 27: Teste e Benchmarking\n")
	format.PrintHelpMe(hlp)
}

// executeSections formats and processes a given section of the project.
// It takes a section name as a string parameter and applies formatting
// to that section using the FormatSection function from the format package.
func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
