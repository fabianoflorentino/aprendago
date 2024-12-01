// Package exercicios_ninja_nivel_2 provides functions to execute and display
// exercises for the Ninja Level 2 course. It includes functions to execute
// specific sections of the exercises, display a menu of exercise options,
// and provide help descriptions for each exercise. The package relies on
// the format package for formatting and displaying sections and help information.
package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for Ninja Level 2 in the "aprendago" project.
// This constant is used to reference the directory where the
// specific exercise files are stored within the internal package.
const (
	rootDir = "internal/exercicios_ninja_nivel_2"
)

// ExerciciosNinjaNivel2 prints a header for the Ninja Level 2 exercises
// and sequentially executes sections labeled from "Na prática exercício #1" to "Na prática exercício #7".
// Each section is executed by calling the executeSection function with the respective section label as an argument.
func ExerciciosNinjaNivel2() {
	fmt.Printf("\n\n05 - Exercicios: Ninja Nível 2\n")

	executeSection("Na prática exercício #1")
	executeSection("Na prática exercício #2")
	executeSection("Na prática exercício #3")
	executeSection("Na prática exercício #4")
	executeSection("Na prática exercício #5")
	executeSection("Na prática exercício #6")
	executeSection("Na prática exercício #7")
}

// MenuExerciciosNinjaNivel2 returns a slice of format.MenuOptions for various exercises
// in "Ninja Level 2". Each menu option includes a command-line option string and an
// associated function to execute when that option is selected.
//
// The available options are:
// - "--na-pratica-exercicio-1 --nivel-2": Executes the section "Na prática exercício #1".
// - "--na-pratica-exercicio-1 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #1".
// - "--na-pratica-exercicio-2 --nivel-2": Executes the section "Na prática exercício #2".
// - "--na-pratica-exercicio-2 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #2".
// - "--na-pratica-exercicio-3 --nivel-2": Executes the section "Na prática exercício #3".
// - "--na-pratica-exercicio-3 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #3".
// - "--na-pratica-exercicio-4 --nivel-2": Executes the section "Na prática exercício #4".
// - "--na-pratica-exercicio-4 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #4".
// - "--na-pratica-exercicio-5 --nivel-2": Executes the section "Na prática exercício #5".
// - "--na-pratica-exercicio-5 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #5".
// - "--na-pratica-exercicio-6 --nivel-2": Executes the section "Na prática exercício #6".
// - "--na-pratica-exercicio-6 --nivel-2 --resolucao": Executes the resolution function for "Na prática exercício #6".
// - "--na-pratica-exercicio-7 --nivel-2": Executes the section "Na prática exercício #7".
// - "--na-pratica-exercicio-7 --nivel-2 --prova": Executes the function to respond to the test for "Na prática exercício #7".
func MenuExerciciosNinjaNivel2([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #1") }},
		{Options: "--exercicio=1 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #2") }},
		{Options: "--exercicio=2 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #3") }},
		{Options: "--exercicio=3 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #4") }},
		{Options: "--exercicio=4 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--exercicio=5 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #5") }},
		{Options: "--exercicio=5 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--exercicio=6 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #6") }},
		{Options: "--exercicio=6 --nivel=2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--exercicio=7 --nivel=2", ExecFunc: func() { executeSection("Na prática exercício #7") }},
		{Options: "--exercicio=7 --nivel=2 --prova", ExecFunc: func() { RespondaAProva() }},
	}
}

// HelpMeExerciciosNinjaNivel2 provides a list of available commands and their descriptions
// for the "Exercícios Ninja - Nível 2" section. Each command is represented by a flag and
// a description, indicating what the command does. The function prints out the chapter title
// and then uses the format.PrintHelpMe function to display the list of commands and their
// respective descriptions.
func HelpMeExerciciosNinjaNivel2() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=2", Description: "Apresenta o primeiro exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=1 --nivel=2 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=2 --nivel=2", Description: "Apresenta o segundo exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=2 --nivel=2 --resolucao", Description: "Exibe a resolução do segundo exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=3 --nivel=2", Description: "Apresenta o terceiro exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=3 --nivel=2 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=4 --nivel=2", Description: "Apresenta o quarto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=4 --nivel=2 --resolucao", Description: "Exibe a resolução do quarto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=5 --nivel=2", Description: "Apresenta o quinto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=5 --nivel=2 --resolucao", Description: "Exibe a resolução do quinto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=6 --nivel=2", Description: "Apresenta o sexto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=6 --nivel=2 --resolucao", Description: "Exibe a resolução do sexto exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=7 --nivel=2", Description: "Apresenta o sétimo exercício prático do nível 2.", Width: 0},
		{Flag: "--exercicio=7 --nivel=2 --prova", Description: "Exibe a prova do sétimo exercício prático do nível 2.", Width: 0},
	}

	fmt.Println("\nCapítulo 5: Exercícios Ninja - Nível 2")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
