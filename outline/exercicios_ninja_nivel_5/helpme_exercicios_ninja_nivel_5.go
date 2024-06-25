package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

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
