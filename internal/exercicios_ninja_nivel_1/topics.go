// Package exercicios_ninja_nivel_1 provides functions to execute and display
// exercises for the Ninja Level 1 course. It includes functionalities to
// present exercises, show their resolutions, and provide help information
// about the exercises. The package interacts with the format package to
// format and display sections and help information.
package exercicios_ninja_nivel_1

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for Ninja Level 1 in the "aprendago" project.
// This constant is used to reference the directory where the
// exercise files are stored within the internal package structure.
const (
	rootDir = "internal/exercicios_ninja_nivel_1"
)

// ExerciciosNinjaNivel1 prints a header for "Ninja Level 1 Exercises" and
// executes a series of sections related to practical exercises. Each section
// is executed by calling the executeSection function with a specific exercise
// description as an argument. This function serves as an entry point to
// demonstrate and run various exercises for Ninja Level 1.
func ExerciciosNinjaNivel1() {
	fmt.Printf("\n\n03 - Exercicios: Ninja Nível 1\n")

	executeSection("Contribua seu código")
	executeSection("Na prática - Exercício #1")
	executeSection("Na prática - Exercício #2")
	executeSection("Na prática - Exercício #3")
	executeSection("Na prática - Exercício #4")
	executeSection("Na prática - Exercício #5")
	executeSection("Na prática - Exercício #6")
}

// MenuExerciciosNinjaNivel1 returns a slice of format.MenuOptions for the Ninja Level 1 exercises.
// Each menu option includes a command-line option string and an associated execution function.
// The available options include:
// - "--contribua-seu-codigo": Executes the "Contribua seu código" section.
// - "--na-pratica-exercicio-1 --nivel-1": Executes the "Na prática - Exercício #1" section.
// - "--na-pratica-exercicio-1 --nivel-1 --resolucao": Executes the resolution for "Na prática - Exercício #1".
// - "--na-pratica-exercicio-2 --nivel-1": Executes the "Na prática - Exercício #2" section.
// - "--na-pratica-exercicio-2 --nivel-1 --resolucao": Executes the resolution for "Na prática - Exercício #2".
// - "--na-pratica-exercicio-3 --nivel-1": Executes the "Na prática - Exercício #3" section.
// - "--na-pratica-exercicio-3 --nivel-1 --resolucao": Executes the resolution for "Na prática - Exercício #3".
// - "--na-pratica-exercicio-4 --nivel-1": Executes the "Na prática - Exercício #4" section.
// - "--na-pratica-exercicio-4 --nivel-1 --resolucao": Executes the resolution for "Na prática - Exercício #4".
// - "--na-pratica-exercicio-5 --nivel-1": Executes the "Na prática - Exercício #5" section.
// - "--na-pratica-exercicio-5 --nivel-1 --resolucao": Executes the resolution for "Na prática - Exercício #5".
// - "--na-pratica-exercicio-6 --nivel-1": Executes the "Na prática - Exercício #6" section.
// - "--na-pratica-exercicio-6 --nivel-1 --prova": Executes the "Responda a Prova" section.
func MenuExerciciosNinjaNivel1([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--contribua-seu-codigo", ExecFunc: func() { executeSection("Contribua seu código") }},
		{Options: "--na-pratica-exercicio-1 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-1 --prova", ExecFunc: func() { RespondaAProva() }},
	}
}

// HelpMeExerciciosNinjaNivel1 prints a list of available commands and their descriptions
// for the "Exercícios Ninja Nível 1" section of the course. Each command is associated
// with a flag and a description that explains what the command does. The function
// organizes these commands into a slice of HelpMe structs and then prints them using
// the format.PrintHelpMe function. This helps users understand what exercises and
// resolutions are available for this level of the course.
func HelpMeExerciciosNinjaNivel1() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-1", Description: "Apresenta o primeiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-1 --resolucao", Description: "Exibe a resolução do primeiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-1", Description: "Apresenta o segundo exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-1 --resolucao", Description: "Exibe a resolução do segundo exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-1", Description: "Apresenta o terceiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-1 --resolucao", Description: "Exibe a resolução do terceiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-1", Description: "Apresenta o quarto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-1 --resolucao", Description: "Exibe a resolução do quarto exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-1", Description: "Apresenta o quinto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-1 --resolucao", Description: "Exibe a resolução do quinto exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-1", Description: "Apresenta o sexto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-1 --prova", Description: "Exibe a prova do sexto exercício prático.", Width: 0},
	}

	fmt.Println("\nCapítulo 3: Exercícios Ninja Nível 1")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
