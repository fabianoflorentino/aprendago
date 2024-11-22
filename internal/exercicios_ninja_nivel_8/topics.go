package exercicios_ninja_nivel_8

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_8"

func ExerciciosNinjaNivel8() {
	fmt.Printf("\n\nCapítulo 17: Exercícios Ninja Nível 8\n")

	executeSections("Na prática: exercício #1")
	executeSections("Na prática: exercício #2")
	executeSections("Na prática: exercício #3")
	executeSections("Na prática: exercício #4")
	executeSections("Na prática: exercício #5")
}

func MenuExerciciosNinjaNivel8([]string) []format.MenuOptions {

	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-8", ExecFunc: func() { executeSections("Na prática: exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-8", ExecFunc: func() { executeSections("Na prática: exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-8", ExecFunc: func() { executeSections("Na prática: exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-8", ExecFunc: func() { executeSections("Na prática: exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-8", ExecFunc: func() { executeSections("Na prática: exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-8 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
	}
}

func HelpMeExerciciosNinjaNivel8() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-8", Description: "Apresenta o primeiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-8 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-8", Description: "Apresenta o segundo exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-8 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-8", Description: "Apresenta o terceiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-8 --resolucao", Description: "Apresenta a resolução do terceiro exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-8", Description: "Apresenta o quarto exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-8 --resolucao", Description: "Apresenta a resolução do quarto exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-8", Description: "Apresenta o quinto exercício prático do Nível 8.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-8 --resolucao", Description: "Apresenta a resolução do quinto exercício prático do Nível 8.", Width: 0},
	}

	fmt.Printf("\nCapítulo 17: Exercícios Ninja Nível 8\n")
	format.PrintHelpMe(hlp)
}

func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
