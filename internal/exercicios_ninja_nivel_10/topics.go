// Package exercicios_ninja_nivel_10 provides functions to execute and display
// exercises for Ninja Level 10 from the "Aprenda Go" project. It includes
// functions to execute specific sections of exercises, generate menu options
// for command-line interaction, and display help information for the exercises.
package exercicios_ninja_nivel_10

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for Ninja Level 10 within the project structure.
const (
	rootDir = "internal/exercicios_ninja_nivel_10"
)

// ExerciciosNinjaNivel10 prints the title of Chapter 22 and executes a series of sections
// labeled as "Na prática: Exercício #1" through "Na prática: Exercício #7". Each section
// is executed by calling the executeSection function with the respective section title.
func ExerciciosNinjaNivel10() {
	fmt.Printf("\n\nCapítulo 22: Exercícios Ninja Nível 10\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
}

// MenuExerciciosNinjaNivel10 returns a slice of format.MenuOptions for the exercises in Ninja Level 10.
// Each menu option includes a command-line option string and a corresponding execution function.
// The execution functions either execute a section or call a resolution function for the respective exercise.
func MenuExerciciosNinjaNivel10([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

// HelpMeExerciciosNinjaNivel10 provides a list of available commands and their descriptions
// for the "Exercícios Ninja Nível 10" section. It prints out the chapter title and uses the
// format.PrintHelpMe function to display the help information for each exercise.
func HelpMeExerciciosNinjaNivel10() {
	help := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-10", Description: "Exibe a descrição do Exercício #1"},
		{Flag: "--na-pratica-exercicio-1 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #1"},
		{Flag: "--na-pratica-exercicio-2 --nivel-10", Description: "Exibe a descrição do Exercício #2"},
		{Flag: "--na-pratica-exercicio-2 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #2"},
		{Flag: "--na-pratica-exercicio-3 --nivel-10", Description: "Exibe a descrição do Exercício #3"},
		{Flag: "--na-pratica-exercicio-3 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #3"},
		{Flag: "--na-pratica-exercicio-4 --nivel-10", Description: "Exibe a descrição do Exercício #4"},
		{Flag: "--na-pratica-exercicio-4 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #4"},
		{Flag: "--na-pratica-exercicio-5 --nivel-10", Description: "Exibe a descrição do Exercício #5"},
		{Flag: "--na-pratica-exercicio-5 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #5"},
		{Flag: "--na-pratica-exercicio-6 --nivel-10", Description: "Exibe a descrição do Exercício #6"},
		{Flag: "--na-pratica-exercicio-6 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #6"},
		{Flag: "--na-pratica-exercicio-7 --nivel-10", Description: "Exibe a descrição do Exercício #7"},
		{Flag: "--na-pratica-exercicio-7 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #7"},
	}

	fmt.Printf("\n\nCapítulo 22: Exercícios Ninja Nível 10\n")
	format.PrintHelpMe(help)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting based on the root directory
// and the specified section.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
