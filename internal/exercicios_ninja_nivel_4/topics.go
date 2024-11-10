package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_4"

func ExerciciosNinjaNivel4() {
	fmt.Printf("Capítulo 9: Exercícios Ninja Nível 4\n\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
	executeSection("Na prática: Exercício #8")
	executeSection("Na prática: Exercício #9")
	executeSection("Na prática: Exercício #10")
}

func MenuExerciciosNinjaNivel4([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--na-pratica-exercicio-1 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--na-pratica-exercicio-2 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--na-pratica-exercicio-3 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--na-pratica-exercicio-4 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-5 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--na-pratica-exercicio-5 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--na-pratica-exercicio-6 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--na-pratica-exercicio-6 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--na-pratica-exercicio-7 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--na-pratica-exercicio-7 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
		{Options: "--na-pratica-exercicio-8 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #8") }},
		{Options: "--na-pratica-exercicio-8 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio8() }},
		{Options: "--na-pratica-exercicio-9 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #9") }},
		{Options: "--na-pratica-exercicio-9 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio9() }},
		{Options: "--na-pratica-exercicio-10 --nivel-4", ExecFunc: func() { executeSection("Na prática: Exercício #10") }},
		{Options: "--na-pratica-exercicio-10 --nivel-4 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio10() }},
	}
}

func HelpMeExerciciosNinjaNivel4() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-4", Description: "Apresenta o primeiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-4 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-4", Description: "Apresenta o segundo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-4 --resolucao", Description: "Exibe a resolução do segundo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-4", Description: "Apresenta o terceiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-4 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-4", Description: "Apresenta o quarto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-4 --resolucao", Description: "Exibe a resolução do quarto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-4", Description: "Apresenta o quinto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-4 --resolucao", Description: "Exibe a resolução do quinto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-4", Description: "Apresenta o sexto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-6 --nivel-4 --resolucao", Description: "Exibe a resolução do sexto exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-4", Description: "Apresenta o sétimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-7 --nivel-4 --resolucao", Description: "Exibe a resolução do sétimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-8 --nivel-4", Description: "Apresenta o oitavo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-8 --nivel-4 --resolucao", Description: "Exibe a resolução do oitavo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-9 --nivel-4", Description: "Apresenta o nono exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-9 --nivel-4 --resolucao", Description: "Exibe a resolução do nono exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-10 --nivel-4", Description: "Apresenta o décimo exercício prático do Nível 4.", Width: 0},
		{Flag: "--na-pratica-exercicio-10 --nivel-4 --resolucao", Description: "Exibe a resolução do décimo exercício prático do Nível 4.", Width: 0},
	}

	fmt.Println("\nCapítulo 9: Exercícios Ninja Nível 4")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
