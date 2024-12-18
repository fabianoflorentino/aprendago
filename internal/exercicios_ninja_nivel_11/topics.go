// Package exercicios_ninja_nivel_11 provides functions to execute and display
// exercises from the Ninja Level 11 course. It includes functions to present
// the exercises, show their resolutions, and display help information for
// the available commands. The package utilizes the format package for
// formatting and displaying sections and help information.
package exercicios_ninja_nivel_11

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for level 11 of the ninja training program.
const (
	rootDir = "internal/exercicios_ninja_nivel_11"
)

// ExercicicioNinjaNivel11 prints the title of the chapter and executes a series of sections
// corresponding to exercises in Ninja Level 11. Each section is executed by calling the
// executeSection function with the name of the exercise as an argument.
func ExerciciosNinjaNivel11() {
	fmt.Printf("\n\nCapítulo 24: Exercicios: Ninja Nível 11\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
}

// MenuExerciciosNinjaNivel11 returns a slice of format.MenuOptions for the exercises of Ninja Level 11.
// Each MenuOption contains a command-line option string and an associated execution function.
// The options include commands for executing and resolving exercises 1 through 5.
func MenuExerciciosNinjaNivel11([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-11", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-11 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-11", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-11 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-11", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-11 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-11", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-11 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-11", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-11 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
	}
}

// HelpMeExerciciosNinjaNivel11 provides a list of available commands and their descriptions
// for the practical exercises of level 11 in the Ninja course. It prints the chapter title
// and uses the format.PrintHelpMe function to display the help information.
func HelpMeExerciciosNinjaNivel11() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-11", Description: "Apresenta o primeiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-11 --resolucao", Description: "Exibe a resolução do primeiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-11", Description: "Apresenta o segundo exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-11 --resolucao", Description: "Exibe a resolução do segundo exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-11", Description: "Apresenta o terceiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-11 --resolucao", Description: "Exibe a resolução do terceiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-11", Description: "Apresenta o quarto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-11 --resolucao", Description: "Exibe a resolução do quarto exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-11", Description: "Apresenta o quinto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-11 --resolucao", Description: "Exibe a resolução do quinto exercício prático.", Width: 0},
	}

	fmt.Printf("\nCapítulo 24: Exercicios Ninja Nível 11\n")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
