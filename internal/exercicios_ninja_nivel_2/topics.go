package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_2"

func ExerciciosNinjaNivel2() {
	fmt.Printf("05 - Exercicios: Ninja Nível 2\n\n")

	executeSection("Na prática exercício #1")
	executeSection("Na prática exercício #2")
	executeSection("Na prática exercício #3")
	executeSection("Na prática exercício #4")
	executeSection("Na prática exercício #5")
	executeSection("Na prática exercício #6")
	executeSection("Na prática exercício #7")
}

func MenuExerciciosNinjaNivel2([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-2 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-2", ExecFunc: func() { executeSection("Na prática exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-2 --prova", ExecFunc: func() { RespondaAProva() }},
	}
}

func HelpMeExerciciosNinjaNivel2() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-2", Description: "Apresenta o primeiro exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-2 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-2", Description: "Apresenta o segundo exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-2 --resolucao", Description: "Exibe a resolução do segundo exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-2", Description: "Apresenta o terceiro exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-2 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-2", Description: "Apresenta o quarto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-2 --resolucao", Description: "Exibe a resolução do quarto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-2", Description: "Apresenta o quinto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-2 --resolucao", Description: "Exibe a resolução do quinto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-2", Description: "Apresenta o sexto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-2 --resolucao", Description: "Exibe a resolução do sexto exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-2", Description: "Apresenta o sétimo exercício prático do nível 2.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-2 --prova", Description: "Exibe a prova do sétimo exercício prático do nível 2.", Width: 0},
	}

	fmt.Println("\nCapítulo 5: Exercícios Ninja - Nível 2")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
