package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_5"

func ExerciciosNinjaNivel5() {
	fmt.Printf("\n\nCapítulo 11: Exercícios Ninja Nível 5\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
}

func MenuExerciciosNinjaNivel5([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-5", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-5", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-5", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-5", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
	}
}

func HelpMeExerciciosNinjaNivel5() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-5", Description: "Apresenta o primeiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-5 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-5", Description: "Apresenta o segundo exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-5 --resolucao", Description: "Exibe a resolução do segundo exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-5", Description: "Apresenta o terceiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-5 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-5", Description: "Apresenta o quarto exercício prático do Nível 5.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-5 --resolucao", Description: "Exibe a resolução do quarto exercício prático do Nível 5.", Width: 0},
	}

	fmt.Println("\nCapítulo 11: Exercícios Ninja Nível 5")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
