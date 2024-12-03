// Package exercicios_ninja_nivel_9 provides functions to display and execute
// exercises from Chapter 20: Ninja Level 9. It includes functions to list
// exercises, display help information, and execute specific sections of the
// exercises. The package relies on the format package for formatting output.
package exercicios_ninja_nivel_9

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for Ninja Level 9 in the "aprendago" project.
// This constant is used to reference the directory within the project structure.
const (
	rootDir = "internal/exercicios_ninja_nivel_9"
)

// ExerciciosNinjaNivel9 prints the title for Chapter 20: Exercícios Ninja Nível 9
// and sequentially executes sections for each exercise from 1 to 7.
// Each exercise section is executed by calling the executeSection function
// with the corresponding exercise description as an argument.
func ExerciciosNinjaNivel9() {
	fmt.Printf("\n\nCapítulo 20: Exercícios Ninja Nível 9\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
}

// MenuExerciciosNinjaNivel9 returns a slice of format.MenuOptions for the exercises in Ninja Level 9.
// Each MenuOption contains a command-line option string and an associated execution function.
// The options include commands for executing and resolving each exercise from 1 to 7.
// The execution functions either call executeSection with the exercise description or call a specific resolution function for the exercise.
func MenuExerciciosNinjaNivel9([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-9", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

// HelpMeExerciciosNinjaNivel9 provides a list of help commands for the exercises
// in chapter 20, Ninja Level 9. It creates a slice of HelpMe structs, each containing
// a flag and a description for various exercises and their resolutions. The function
// then prints the chapter title and uses the format.PrintHelpMe function to display
// the help information.
func HelpMeExerciciosNinjaNivel9() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-9", Description: "Exibe o Exercício 1 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-1 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 1 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-2 --nivel-9", Description: "Exibe o Exercício 2 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-2 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 2 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-3 --nivel-9", Description: "Exibe o Exercício 3 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-3 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 3 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-4 --nivel-9", Description: "Exibe o Exercício 4 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-4 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 4 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-5 --nivel-9", Description: "Exibe o Exercício 5 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-5 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 5 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-6 --nivel-9", Description: "Exibe o Exercício 6 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-6 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 6 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-7 --nivel-9", Description: "Exibe o Exercício 7 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-7 --nivel-9 --resolucao", Description: "Exibe a resolução do Exercício 7 do capítulo 20"},
	}

	fmt.Printf("\nCapítulo 20: Exercícios Ninja Nível 9\n")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section within the root directory.
// It takes a section name as a string parameter and uses the FormatSection function
// from the format package to apply the formatting to the specified section.
//
// Parameters:
//   - section: A string representing the name of the section to be formatted.
//
// Example usage:
//
//	executeSection("chapter1")
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
