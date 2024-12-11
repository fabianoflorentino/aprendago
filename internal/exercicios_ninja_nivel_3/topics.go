// Package exercicios_ninja_nivel_3 provides a set of exercises for Ninja Level 3.
// It includes functions to execute and display various practical exercises,
// as well as functions to generate menu options and help documentation for these exercises.
package exercicios_ninja_nivel_3

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for Ninja Level 3 in the "aprendago" project.
// This constant is used to reference the specific directory within
// the internal package structure of the project.
const (
	rootDir = "internal/exercicios_ninja_nivel_3"
)

// ExerciciosNinjaNivel3 prints a header for the "Ninja Level 3" exercises
// and sequentially executes a series of sections labeled from "Na prática - Exercício #1"
// to "Na prática - Exercício #10". Each section is executed by calling the executeSection function
// with the corresponding section name as an argument.
func ExerciciosNinjaNivel3() {
	fmt.Printf("\n\n07 - Exercicios: Ninja Nível 3\n")

	executeSection("Na prática - Exercício #1")
	executeSection("Na prática - Exercício #2")
	executeSection("Na prática - Exercício #3")
	executeSection("Na prática - Exercício #4")
	executeSection("Na prática - Exercício #5")
	executeSection("Na prática - Exercício #6")
	executeSection("Na prática - Exercício #7")
	executeSection("Na prática - Exercício #8")
	executeSection("Na prática - Exercício #9")
	executeSection("Na prática - Exercício #10")
}

// MenuExerciciosNinjaNivel3 returns a slice of format.MenuOptions for the exercises in Ninja Level 3.
// Each menu option includes a command-line option string and a corresponding execution function.
// The execution functions either execute a section or call a resolution function for each exercise.
// The options are structured to handle both the exercise execution and its resolution for exercises 1 to 10.
func MenuExerciciosNinjaNivel3([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #1") }},
		{Options: "--exercicio=1 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #2") }},
		{Options: "--exercicio=2 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #3") }},
		{Options: "--exercicio=3 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #4") }},
		{Options: "--exercicio=4 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--exercicio=5 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #5") }},
		{Options: "--exercicio=5 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--exercicio=6 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #6") }},
		{Options: "--exercicio=6 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--exercicio=7 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #7") }},
		{Options: "--exercicio=7 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
		{Options: "--exercicio=8 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #8") }},
		{Options: "--exercicio=8 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio8() }},
		{Options: "--exercicio=9 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #9") }},
		{Options: "--exercicio=9 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio9() }},
		{Options: "--exercicio=10 --nivel=3", ExecFunc: func() { executeSection("Na prática - Exercício #10") }},
		{Options: "--exercicio=10 --nivel=3 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio10() }},
	}
}

// HelpMeExerciciosNinjaNivel3 provides a list of available commands and their descriptions
// for the practical exercises of level 3 in Chapter 7. It prints out the help information
// for each exercise, including flags for viewing the exercise and its resolution.
func HelpMeExerciciosNinjaNivel3() {
	hlp := []format.HelpMe{
		// {Flag: "", Description: "", Width: 0},
		{Flag: "--exercicio=1 --nivel=3", Description: "Apresenta o primeiro exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=1 --nivel=3 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=2 --nivel=3", Description: "Apresenta o segundo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=2 --nivel=3 --resolucao", Description: "Exibe a resolução do segundo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=3 --nivel=3", Description: "Apresenta o terceiro exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=3 --nivel=3 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=4 --nivel=3", Description: "Apresenta o quarto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=4 --nivel=3 --resolucao", Description: "Exibe a resolução do quarto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=5 --nivel=3", Description: "Apresenta o quinto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=5 --nivel=3 --resolucao", Description: "Exibe a resolução do quinto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=6 --nivel=3", Description: "Apresenta o sexto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=6 --nivel=3 --resolucao", Description: "Exibe a resolução do sexto exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=7 --nivel=3", Description: "Apresenta o sétimo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=7 --nivel=3 --resolucao", Description: "Exibe a resolução do sétimo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=8 --nivel=3", Description: "Apresenta o oitavo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=8 --nivel=3 --resolucao", Description: "Exibe a resolução do oitavo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=9 --nivel=3", Description: "Apresenta o nono exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=9 --nivel=3 --resolucao", Description: "Exibe a resolução do nono exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=10 --nivel=3", Description: "Apresenta o décimo exercício prático do nível 3.", Width: 0},
		{Flag: "--exercicio=10 --nivel=3 --resolucao", Description: "Exibe a resolução do décimo exercício prático do nível 3.", Width: 0},
	}

	fmt.Println("\nCapítulo 7: Exercícios Ninja Nível 3")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string and uses the FormatSection function
// from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
