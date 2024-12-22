// Package exercicios_ninja_nivel_12 provides functions to execute and display
// exercises for level 12 of the ninja training program. It includes functions
// to present the exercises, display their resolutions, and provide help
// information for the exercises. The package relies on the format package
// for formatting and displaying the sections and help information.
package exercicios_ninja_nivel_12

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the directory containing
// the exercises for level 11 of the ninja training program.
const (
	rootDir = "internal/exercicios_ninja_nivel_12"
)

// ExerciciosNinjaNivel12 prints the title of the chapter and executes three sections of exercises.
// Each section is executed by calling the executeSection function with the section title as an argument.
func ExerciciosNinjaNivel12() {
	fmt.Printf("\n\nCapítulo 26: Exercicios: Ninja Nível 12\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
}

// MenuExerciciosNinjaNivel12 returns a slice of format.MenuOptions for the exercises in Ninja Level 12.
// Each menu option includes a command-line option string and an associated execution function.
// The execution functions either execute a section or call a resolution function for the respective exercise.
func MenuExerciciosNinjaNivel12([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-12", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-12 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-12", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-12 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-12", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-12 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
	}
}

// HelpMeExerciciosNinjaNivel12 provides help information for the exercises
// of Ninja Level 12. It lists the available flags and their descriptions
// for each practical exercise and its resolution in the course.
func HelpMeExerciciosNinjaNivel12() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-12", Description: "Apresenta o primeiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-12 --resolucao", Description: "Exibe a resolução do primeiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-12", Description: "Apresenta o segundo exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-12 --resolucao", Description: "Exibe a resolução do segundo exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-12", Description: "Apresenta o terceiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-12 --resolucao", Description: "Exibe a resolução do terceiro exercício prático.", Width: 0},
	}

	fmt.Printf("\nCapítulo 26: Exercicios Ninja Nível 12\n")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the project.
// It takes a section name as a string parameter and uses the FormatSection
// function from the format package to apply formatting to the specified section
// within the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
