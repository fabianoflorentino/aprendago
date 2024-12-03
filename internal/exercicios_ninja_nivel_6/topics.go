// Package exercicios_ninja_nivel_6 provides functions to execute and display
// exercises from Chapter 13: Ninja Level 6. It includes functions to execute
// each exercise, display a menu of exercises, and provide help descriptions
// for each exercise option. The exercises are formatted and displayed using
// the format package.
package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir is a constant that holds the relative path to the directory
// containing the exercises for Ninja Level 6 within the project.
// This path is used to reference the location of the exercises
// within the internal package structure.
const (
	rootDir = "internal/exercicios_ninja_nivel_6"
)

// ExerciciosNinjaNivel6 prints the title for Chapter 13: Exercícios Ninja Nível 6
// and sequentially executes a series of sections labeled from "Na Prática - Exercício #1"
// to "Na Prática - Exercício #11". Each section is executed by calling the executeSection
// function with the respective section title as an argument.
func ExerciciosNinjaNivel6() {
	fmt.Printf("\n\nCapítulo 13: Exercícios Ninja Nível 6\n")

	executeSection("Na Prática - Exercício #1")
	executeSection("Na Prárica - Exercício #2")
	executeSection("Na Prática - Exercício #3")
	executeSection("Na Prática - Exercício #4")
	executeSection("Na Prática - Exercício #5")
	executeSection("Na Prática - Exercício #6")
	executeSection("Na Prática - Exercício #7")
	executeSection("Na Prática - Exercício #8")
	executeSection("Na Prática - Exercício #9")
	executeSection("Na Prática - Exercício #10")
	executeSection("Na Prática - Exercício #11")
}

// MenuExerciciosNinjaNivel6 returns a slice of format.MenuOptions for the exercises of Ninja Level 6.
// Each menu option includes a command line option string and an associated execution function.
// The options are structured to handle both the execution and resolution of each exercise.
// The function does not take any parameters and returns a slice of format.MenuOptions.
func MenuExerciciosNinjaNivel6([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
		{Options: "--na-pratica-exercicio-8 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #8") }},
		{Options: "--na-pratica-exercicio-8 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio8() }},
		{Options: "--na-pratica-exercicio-9 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #9") }},
		{Options: "--na-pratica-exercicio-9 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio9() }},
		{Options: "--na-pratica-exercicio-10 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #10") }},
		{Options: "--na-pratica-exercicio-10 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio10() }},
		{Options: "--na-pratica-exercicio-11 --nivel-6", ExecFunc: func() { executeSection("Na Prática - Exercício #11") }},
		{Options: "--na-pratica-exercicio-11 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio11() }},
	}
}

// HelMeExerciciosNinjaNivel6 prints a help menu for the exercises of Ninja Level 6.
// It creates a slice of HelpMe structs, each containing a flag and description for
// various practical exercises and their resolutions. The function then prints the
// chapter title and uses the format.PrintHelpMe function to display the help menu.
func HelMeExerciciosNinjaNivel6() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-6", Description: "Apresenta o primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-6 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-6", Description: "Apresenta o segundo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-6 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-6", Description: "Apresenta o terceiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-6 --resolucao", Description: "Apresenta a resolução do terceiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-6", Description: "Apresenta o quarto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-6 --resolucao", Description: "Apresenta a resolução do quarto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-6", Description: "Apresenta o quinto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-6 --resolucao", Description: "Apresenta a resolução do quinto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-6", Description: "Apresenta o sexto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-6 --resolucao", Description: "Apresenta a resolução do sexto exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-6", Description: "Apresenta o sétimo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-6 --resolucao", Description: "Apresenta a resolução do sétimo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-8 --nivel-6", Description: "Apresenta o oitavo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-8 --nivel-6 --resolucao", Description: "Apresenta a resolução do oitavo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-9 --nivel-6", Description: "Apresenta o nono exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-9 --nivel-6 --resolucao", Description: "Apresenta a resolução do nono exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-10 --nivel-6", Description: "Apresenta o décimo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-10 --nivel-6 --resolucao", Description: "Apresenta a resolução do décimo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-11 --nivel-6", Description: "Apresenta o décimo primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-11 --nivel-6 --resolucao", Description: "Apresenta a resolução do décimo primeiro exercício prático do Nível 6.", Width: 0},
	}

	fmt.Println("\nCapítulo 13: Exercícios Ninja Nível 6")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section within the root directory.
// It takes a section name as a string parameter and uses the FormatSection function
// from the format package to apply the formatting to the specified section.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
