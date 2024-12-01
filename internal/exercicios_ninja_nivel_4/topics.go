// Package exercicios_ninja_nivel_4 provides functions to execute and display
// exercises from Chapter 9: Ninja Level 4. It includes functions to present
// the exercises, display their resolutions, and provide help information
// about the available commands and their descriptions.
package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir is a constant that holds the relative path to the directory
// containing the exercises for Ninja Level 4. This path is used to
// reference the location of the exercises within the internal package
// structure of the project.
const (
	rootDir = "internal/exercicios_ninja_nivel_4"
)

// ExerciciosNinjaNivel4 prints the title for Chapter 9: Exercícios Ninja Nível 4
// and sequentially executes sections labeled from "Na prática: Exercício #1" to "Na prática: Exercício #10".
// Each section is executed by calling the executeSection function with the respective section title as an argument.
func ExerciciosNinjaNivel4() {
	fmt.Printf("\n\nCapítulo 9: Exercícios Ninja Nível 4\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
	executeSection("Na prática: Exercício #8")
	executeSection("Na prática: Exercício #9")
	executeSection("Na prática: Exercício #10")
}

// MenuExerciciosNinjaNivel4 returns a slice of format.MenuOptions for the exercises in level 4.
// Each menu option includes a set of command-line options and an associated execution function.
// The execution functions either execute a section or call a specific resolution function for each exercise.
func MenuExerciciosNinjaNivel4([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--exercicio=1 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--exercicio=2 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--exercicio=3 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--exercicio=4 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--exercicio=5 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--exercicio=5 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--exercicio=6 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--exercicio=6 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--exercicio=7 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--exercicio=7 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
		{Options: "--exercicio=8 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #8") }},
		{Options: "--exercicio=8 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio8() }},
		{Options: "--exercicio=9 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #9") }},
		{Options: "--exercicio=9 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio9() }},
		{Options: "--exercicio=10 --nivel=4", ExecFunc: func() { executeSection("Na prática: Exercício #10") }},
		{Options: "--exercicio=10 --nivel=4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio10() }},
	}
}

// HelpMeExerciciosNinjaNivel4 provides a list of available commands and their descriptions
// for the practical exercises of Level 4 in Chapter 9. It prints out the help information
// for each exercise, including flags for displaying the exercise and its resolution.
func HelpMeExerciciosNinjaNivel4() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=4", Description: "Apresenta o primeiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=1 --nivel=4 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=2 --nivel=4", Description: "Apresenta o segundo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=2 --nivel=4 --resolucao", Description: "Exibe a resolução do segundo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=3 --nivel=4", Description: "Apresenta o terceiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=3 --nivel=4 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=4 --nivel=4", Description: "Apresenta o quarto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=4 --nivel=4 --resolucao", Description: "Exibe a resolução do quarto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=5 --nivel=4", Description: "Apresenta o quinto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=5 --nivel=4 --resolucao", Description: "Exibe a resolução do quinto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=6 --nivel=4", Description: "Apresenta o sexto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=6 --nivel=4 --resolucao", Description: "Exibe a resolução do sexto exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=7 --nivel=4", Description: "Apresenta o sétimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=7 --nivel=4 --resolucao", Description: "Exibe a resolução do sétimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=8 --nivel=4", Description: "Apresenta o oitavo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=8 --nivel=4 --resolucao", Description: "Exibe a resolução do oitavo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=9 --nivel=4", Description: "Apresenta o nono exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=9 --nivel=4 --resolucao", Description: "Exibe a resolução do nono exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=10 --nivel=4", Description: "Apresenta o décimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--exercicio=10 --nivel=4 --resolucao", Description: "Exibe a resolução do décimo exercício prático do Nível 4.", Width: 0},
	}

	fmt.Println("\nCapítulo 9: Exercícios Ninja Nível 4")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
