// Package exercicios_ninja_nivel_8 provides functions to execute and display
// exercises for Ninja Level 8 from the "aprendago" project. It includes
// functions to execute specific exercises, display a menu of exercise options,
// and provide help descriptions for each exercise.
package exercicios_ninja_nivel_8

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const (
	rootDir = "internal/exercicios_ninja_nivel_8"
)

// ExerciciosNinjaNivel8 prints the header for Chapter 17: Exercícios Ninja Nível 8
// and executes a series of sections labeled as "Na prática: exercício #1" to "Na prática: exercício #5".
// Each section is executed by calling the executeSections function with the respective section label.
func ExerciciosNinjaNivel8() {
	fmt.Printf("\n\nCapítulo 17: Exercícios Ninja Nível 8\n")

	executeSections("Na prática: exercício #1")
	executeSections("Na prática: exercício #2")
	executeSections("Na prática: exercício #3")
	executeSections("Na prática: exercício #4")
	executeSections("Na prática: exercício #5")
}

// MenuExerciciosNinjaNivel8 returns a slice of format.MenuOptions for the exercises in level 8.
// Each MenuOption contains a command-line option string and an associated execution function.
// The options include commands for executing and resolving exercises 1 through 5.
// The ExecFunc for each option either executes a section or calls a resolution function specific to the exercise.
func MenuExerciciosNinjaNivel8([]string) []format.MenuOptions {

	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=8", ExecFunc: func() { executeSections("Na prática: exercício #1") }},
		{Options: "--exercicio=1 --nivel=8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=8", ExecFunc: func() { executeSections("Na prática: exercício #2") }},
		{Options: "--exercicio=2 --nivel=8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=8", ExecFunc: func() { executeSections("Na prática: exercício #3") }},
		{Options: "--exercicio=3 --nivel=8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=8", ExecFunc: func() { executeSections("Na prática: exercício #4") }},
		{Options: "--exercicio=4 --nivel=8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--exercicio=5 --nivel=8", ExecFunc: func() { executeSections("Na prática: exercício #5") }},
		{Options: "--exercicio=5 --nivel=8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
	}
}

// HelpMeExerciciosNinjaNivel8 provides a list of available commands and their descriptions
// for the practical exercises of Ninja Level 8. It prints out the chapter title and then
// uses the format.PrintHelpMe function to display the help information for each exercise.
// Each exercise has a flag and a description, and some exercises also have a resolution flag
// to show the solution for that exercise.
func HelpMeExerciciosNinjaNivel8() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=8", Description: "Apresenta o primeiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=1 --nivel=8 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=2 --nivel=8", Description: "Apresenta o segundo exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=2 --nivel=8 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=3 --nivel=8", Description: "Apresenta o terceiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=3 --nivel=8 --resolucao", Description: "Apresenta a resolução do terceiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=4 --nivel=8", Description: "Apresenta o quarto exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=4 --nivel=8 --resolucao", Description: "Apresenta a resolução do quarto exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=5 --nivel=8", Description: "Apresenta o quinto exercício prático do Nível 8.", Width: 0},
		{Flag: "--exercicio=5 --nivel=8 --resolucao", Description: "Apresenta a resolução do quinto exercício prático do Nível 8.", Width: 0},
	}

	fmt.Printf("\nCapítulo 17: Exercícios Ninja Nível 8\n")
	format.PrintHelpMe(hlp)
}

// executeSections formats and processes a specific section of the project.
// It takes a section name as a string argument and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
