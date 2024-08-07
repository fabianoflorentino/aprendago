package exercicios_ninja_nivel_1

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

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
