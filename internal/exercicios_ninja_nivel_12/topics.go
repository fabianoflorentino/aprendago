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

func ExerciciosNinjaNivel12() {
	fmt.Printf("\n\nCapítulo 26: Exercicios: Ninja Nível 12\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
}

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

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
