package exercicios_ninja_nivel_1

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_1"

func ExerciciosNinjaNivel1() {
	fmt.Printf("\n\n03 - Exercicios: Ninja Nível 1\n")

	executeSection("Contribua seu código")
	executeSection("Na prática - Exercício #1")
	executeSection("Na prática - Exercício #2")
	executeSection("Na prática - Exercício #3")
	executeSection("Na prática - Exercício #4")
	executeSection("Na prática - Exercício #5")
	executeSection("Na prática - Exercício #6")
}

func MenuExerciciosNinjaNivel1([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--contribua-seu-codigo", ExecFunc: func() { executeSection("Contribua seu código") }},
		{Options: "--na-pratica-exercicio-1 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-1 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-1", ExecFunc: func() { executeSection("Na prática - Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-1 --prova", ExecFunc: func() { RespondaAProva() }},
	}
}

func HelpMeExerciciosNinjaNivel1() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-1", Description: "Apresenta o primeiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-1 --resolucao", Description: "Exibe a resolução do primeiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-1", Description: "Apresenta o segundo exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-1 --resolucao", Description: "Exibe a resolução do segundo exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-1", Description: "Apresenta o terceiro exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-1 --resolucao", Description: "Exibe a resolução do terceiro exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-1", Description: "Apresenta o quarto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-1 --resolucao", Description: "Exibe a resolução do quarto exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-1", Description: "Apresenta o quinto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-1 --resolucao", Description: "Exibe a resolução do quinto exercício prático.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-1", Description: "Apresenta o sexto exercício prático do curso.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-1 --prova", Description: "Exibe a prova do sexto exercício prático.", Width: 0},
	}

	fmt.Println("\nCapítulo 3: Exercícios Ninja Nível 1")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
