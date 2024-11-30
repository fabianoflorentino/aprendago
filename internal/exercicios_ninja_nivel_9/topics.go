package exercicios_ninja_nivel_9

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_9"

func ExerciciosNinjaNivel9() {
	fmt.Printf("\n\nCapítulo 20: Exercícios Ninja Nível 9\n")

	executeSection("Na prática: exercício #1")
	executeSection("Na prática: exercício #2")
	executeSection("Na prática: exercício #3")
	executeSection("Na prática: exercício #4")
	executeSection("Na prática: exercício #5")
	executeSection("Na prática: exercício #6")
	executeSection("Na prática: exercício #7")
}

func MenuExerciciosNinjaNivel9([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-9", ExecFunc: func() { executeSection("Na prática: exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

func HelpMeExerciciosNinjaNivel9() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-9", Description: "Exibe o exercício 1 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-1 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 1 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-2 --nivel-9", Description: "Exibe o exercício 2 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-2 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 2 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-3 --nivel-9", Description: "Exibe o exercício 3 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-3 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 3 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-4 --nivel-9", Description: "Exibe o exercício 4 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-4 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 4 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-5 --nivel-9", Description: "Exibe o exercício 5 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-5 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 5 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-6 --nivel-9", Description: "Exibe o exercício 6 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-6 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 6 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-7 --nivel-9", Description: "Exibe o exercício 7 do capítulo 20"},
		{Flag: "--na-pratica-exercicio-7 --nivel-9 --resolucao", Description: "Exibe a resolução do exercício 7 do capítulo 20"},
	}

	fmt.Printf("\nCapítulo 20: Exercícios Ninja Nível 9\n")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
