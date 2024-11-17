package exercicios_ninja_nivel_7

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_7"

func ExerciciosNinjaNivel7() {
	fmt.Printf("\n\nCapítulo 15: Exercícios Ninja Nível 7\n")

	executeSections("Na prática: exercício #1")
	executeSections("Na prática: exercício #2")
}

func MenuExercicioNinjaNivel7([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-7", ExecFunc: func() { executeSections("Na prática: exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-7 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-7", ExecFunc: func() { executeSections("Na prática: exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-7 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
	}
}

func HelpMeExerciciosNinjaNivel7() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-7", Description: "Apresenta o primeiro exercício prático do Nível 7.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-7 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 7.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-7", Description: "Apresenta o segundo exercício prático do Nível 7.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-7 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 7.", Width: 0},
	}

	format.PrintHelpMe(hlp)
}

func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
