// Package exercicios_ninja_nivel_5 provides functions to execute and display
// exercises for Ninja Level 5. It includes functions to present the exercises,
// display their resolutions, and provide help information for the exercises.
// The package utilizes the format package for formatting and displaying sections
// and help information.
package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory where the exercises
// for Ninja Level 5 are stored within the project. This constant is used to
// reference the directory in a consistent manner throughout the codebase,
// ensuring that any changes to the directory structure can be managed by
// updating this single constant.
const (
	rootDir = "internal/exercicios_ninja_nivel_5"
)

// ExerciciosNinjaNivel5 prints the title of Chapter 11 and executes a series of sections
// labeled as "Na prática: Exercício #1" through "Na prática: Exercício #4". Each section
// is executed by calling the executeSection function with the respective section title.
func ExerciciosNinjaNivel5() {
	fmt.Printf("\n\nCapítulo 11: Exercícios Ninja Nível 5\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
}

// MenuExerciciosNinjaNivel5 returns a slice of format.MenuOptions for the exercises in level 5.
// Each MenuOption contains a command-line option string and an associated function to execute.
// The options include commands for running specific exercises and their resolutions.
// The function takes a slice of strings as an argument, which can be used to pass additional parameters if needed.
func MenuExerciciosNinjaNivel5([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=5", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--exercicio=1 --nivel=5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=5", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--exercicio=2 --nivel=5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=5", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--exercicio=3 --nivel=5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=5", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--exercicio=4 --nivel=5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
	}
}

// HelpMeExerciciosNinjaNivel5 provides a list of help commands for the exercises in Ninja Level 5.
// It creates a slice of HelpMe structs, each containing a flag and its description.
// The flags correspond to different practical exercises and their resolutions in Level 5.
// The function then prints the chapter title and uses the format.PrintHelpMe function to display the help information.
func HelpMeExerciciosNinjaNivel5() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=5", Description: "Apresenta o primeiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=1 --nivel=5 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=2 --nivel=5", Description: "Apresenta o segundo exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=2 --nivel=5 --resolucao", Description: "Exibe a resolução do segundo exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=3 --nivel=5", Description: "Apresenta o terceiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=3 --nivel=5 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=4 --nivel=5", Description: "Apresenta o quarto exercício prático do Nível 5.", Width: 0},
		{Flag: "--exercicio=4 --nivel=5 --resolucao", Description: "Exibe a resolução do quarto exercício prático do Nível 5.", Width: 0},
	}

	fmt.Println("\nCapítulo 11: Exercícios Ninja Nível 5")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
