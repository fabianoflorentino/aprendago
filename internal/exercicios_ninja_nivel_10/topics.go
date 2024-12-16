package exercicios_ninja_nivel_10

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const (
	rootDir = "internal/exercicios_ninja_nivel_10"
)

func ExerciciosNinjaNivel10() {
	fmt.Printf("\n\nCapítulo 22: Exercícios Ninja Nível 10\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
}

func MenuExerciciosNinjaNivel10([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-10", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-10 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

func HelpMeExerciciosNinjaNivel10() {
	help := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-10", Description: "Exibe a descrição do Exercício #1"},
		{Flag: "--na-pratica-exercicio-1 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #1"},
		{Flag: "--na-pratica-exercicio-2 --nivel-10", Description: "Exibe a descrição do Exercício #2"},
		{Flag: "--na-pratica-exercicio-2 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #2"},
		{Flag: "--na-pratica-exercicio-3 --nivel-10", Description: "Exibe a descrição do Exercício #3"},
		{Flag: "--na-pratica-exercicio-3 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #3"},
		{Flag: "--na-pratica-exercicio-4 --nivel-10", Description: "Exibe a descrição do Exercício #4"},
		{Flag: "--na-pratica-exercicio-4 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #4"},
		{Flag: "--na-pratica-exercicio-5 --nivel-10", Description: "Exibe a descrição do Exercício #5"},
		{Flag: "--na-pratica-exercicio-5 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #5"},
		{Flag: "--na-pratica-exercicio-6 --nivel-10", Description: "Exibe a descrição do Exercício #6"},
		{Flag: "--na-pratica-exercicio-6 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #6"},
		{Flag: "--na-pratica-exercicio-7 --nivel-10", Description: "Exibe a descrição do Exercício #7"},
		{Flag: "--na-pratica-exercicio-7 --nivel-10 --resolucao", Description: "Exibe a resolução do Exercício #7"},
	}

	fmt.Printf("\n\nCapítulo 22: Exercícios Ninja Nível 10\n")
	format.PrintHelpMe(help)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
