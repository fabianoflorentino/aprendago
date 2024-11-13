package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_6"

func ExerciciosNinjaNivel6() {
	fmt.Printf("Capítulo 13: Exercícios Ninja Nível 6\n\n")

	executeSection("Na Prática - Exercício #1")
	executeSection("Na Prárica - Exercício #2")
	executeSection("Na Prática - Exercício #3")
	executeSection("Na Prática - Exercício #4")
	executeSection("Na Prática - Exercício #5")
	executeSection("Na Prática - Exercício #6")
	executeSection("Na Prática - Exercício #7")
}

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
		{Options: "--na-pratica-exercicio-7 --nivel-6", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
		{Options: "--na-pratica-exercicio-7 --nivel-6 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

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
	}

	fmt.Println("\nCapítulo 13: Exercícios Ninja Nível 6")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
