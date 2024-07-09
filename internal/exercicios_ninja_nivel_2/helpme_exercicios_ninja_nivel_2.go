package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

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
