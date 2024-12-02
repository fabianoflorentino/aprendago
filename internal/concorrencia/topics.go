// Package concorrencia provides functions and utilities to demonstrate and explain
// various concepts related to concurrency in Go. It includes sections on the
// differences between concurrency and parallelism, the use of goroutines and
// waitgroups, race conditions, mutexes, and atomic operations. The package also
// provides a menu for selecting different concurrency topics and a help function
// to describe the available options.
package concorrencia

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where concurrency-related
// topics are stored within the project. This constant is used to reference the
// directory in a consistent manner throughout the codebase.
const (
	rootDir = "internal/concorrencia"
)

// Concorrencia is a function that prints a header for a section on concurrency
// and then sequentially executes several subsections related to concurrency topics.
// The subsections include:
// - "Concorrência vs Paralelismo": Discusses the difference between concurrency and parallelism.
// - "Goroutines & WaitGroups": Introduces goroutines and wait groups for managing concurrent execution.
// - "Discussão: Condição de corrida": Discusses race conditions and their implications.
// - "Condição de corrida": Provides a deeper dive into race conditions.
// - "Mutex": Explains the use of mutexes for synchronizing access to shared resources.
// - "Atomic": Covers atomic operations for low-level synchronization.
func Concorrencia() {
	fmt.Printf("\n\n18 - Concorrência\n")

	executeSection("Concorrência vs Paralelismo")
	executeSection("Goroutines & WaitGroups")
	executeSection("Discussão: Condição de corrida")
	executeSection("Condição de corrida")
	executeSection("Mutex")
	executeSection("Atomic")
}

// MenuConcorrencia returns a slice of format.MenuOptions, each representing a different concurrency topic.
// Each MenuOption contains an option string and an ExecFunc function that executes a section related to the option.
// The available options are:
// --concorrencia-vs-paralelismo: Executes the "Concorrência vs Paralelismo" section.
// --goroutines-waitgroups: Executes the "Goroutines & WaitGroups" section.
// --discussao-condicao-de-corrida: Executes the "Discussão: Condição de corrida" section.
// --condicao-de-corrida: Executes the "Condição de corrida" section.
// --mutex: Executes the "Mutex" section.
// --atomic: Executes the "Atomic" section.
func MenuConcorrencia([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--concorrencia-vs-paralelismo", ExecFunc: func() { executeSection("Concorrência vs Paralelismo") }},
		{Options: "--goroutines-waitgroups", ExecFunc: func() { executeSection("Goroutines & WaitGroups") }},
		{Options: "--discussao-condicao-de-corrida", ExecFunc: func() { executeSection("Discussão: Condição de corrida") }},
		{Options: "--condicao-de-corrida", ExecFunc: func() { executeSection("Condição de corrida") }},
		{Options: "--mutex", ExecFunc: func() { executeSection("Mutex") }},
		{Options: "--atomic", ExecFunc: func() { executeSection("Atomic") }},
	}
}

// HelpMeConcorrencia provides a list of topics related to concurrency in Go.
// It creates a slice of HelpMe structures, each containing a flag and a description
// of a specific concurrency topic. The topics include:
// - The difference between concurrency and parallelism
// - The use of goroutines and waitgroups
// - Discussion on race conditions
// - The concept of race conditions
// - The use of mutexes
// - The use of atomic operations
// The function then prints the chapter title and calls PrintHelpMe to display the list of topics.
func HelpMeConcorrencia() {
	hlp := []format.HelpMe{
		{Flag: "--concorrencia-vs-paralelismo", Description: "Apresenta a diferença entre concorrência e paralelismo.", Width: 0},
		{Flag: "--goroutines-waitgroups", Description: "Apresenta o uso de goroutines e waitgroups.", Width: 0},
		{Flag: "--discussao-condicao-de-corrida", Description: "Apresenta uma discussão sobre condição de corrida.", Width: 0},
		{Flag: "--condicao-de-corrida", Description: "Apresenta o conceito de condição de corrida.", Width: 0},
		{Flag: "--mutex", Description: "Apresenta o uso de mutex.", Width: 0},
		{Flag: "--atomic", Description: "Apresenta o uso de atomic.", Width: 0},
	}

	fmt.Println("\nCapítulo 18: Concorrência")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
