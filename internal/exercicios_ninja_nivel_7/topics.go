// Package exercicios_ninja_nivel_7 provides functions and utilities for executing and displaying
// exercises from Chapter 15: Ninja Level 7. It includes functions to execute specific exercises,
// display menu options, and provide help descriptions for the exercises.
package exercicios_ninja_nivel_7

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir is a constant that holds the relative path to the directory
// containing the exercises for Ninja Level 7 within the project.
// This path is used to reference the location of the exercises
// and manage file operations related to them.
const (
	rootDir = "internal/exercicios_ninja_nivel_7"
)

// ExerciciosNinjaNivel7 prints the title for Chapter 15: Exercícios Ninja Nível 7
// and executes the sections for the exercises. It calls the executeSections function
// with the titles of the exercises as arguments.
func ExerciciosNinjaNivel7() {
	fmt.Printf("\n\nCapítulo 15: Exercícios Ninja Nível 7\n")

	executeSections("Na prática: exercício #1")
	executeSections("Na prática: exercício #2")
}

// MenuExercicioNinjaNivel7 returns a slice of format.MenuOptions for the exercises in Ninja Level 7.
// Each MenuOption contains a command-line option string and a corresponding function to execute.
// The available options are:
// --na-pratica-exercicio-1 --nivel-7: Executes the first practical exercise of level 7.
// --na-pratica-exercicio-1 --nivel-7 --resolucao: Executes the resolution function for the first practical exercise of level 7.
// --na-pratica-exercicio-2 --nivel-7: Executes the second practical exercise of level 7.
// --na-pratica-exercicio-2 --nivel-7 --resolucao: Executes the resolution function for the second practical exercise of level 7.
func MenuExercicioNinjaNivel7([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=7", ExecFunc: func() { executeSections("Na prática: exercício #1") }},
		{Options: "--exercicio=1 --nivel=7 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=7", ExecFunc: func() { executeSections("Na prática: exercício #2") }},
		{Options: "--exercicio=2 --nivel=7 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
	}
}

// HelpMeExerciciosNinjaNivel7 prints a help message for the exercises of Ninja Level 7.
// It creates a slice of HelpMe structs, each containing a flag and its description,
// and then prints the help message for each exercise and its resolution.
func HelpMeExerciciosNinjaNivel7() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=7", Description: "Apresenta o primeiro exercício prático do Nível 7.", Width: 0},
		{Flag: "--exercicio=1 --nivel=7 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 7.", Width: 0},
		{Flag: "--exercicio=2 --nivel=7", Description: "Apresenta o segundo exercício prático do Nível 7.", Width: 0},
		{Flag: "--exercicio=2 --nivel=7 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 7.", Width: 0},
	}

	fmt.Println("\nCapítulo 15: Exercícios Ninja - Nível 7")
	format.PrintHelpMe(hlp)
}

// executeSections formats and processes a specific section of the project.
// It takes a section name as a parameter and applies formatting to that section
// using the FormatSection function from the format package.
//
// Parameters:
//   - section: A string representing the name of the section to be formatted.
func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
