package outline

import (
	"fmt"

	helpme "github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeExerciciosNinjaNivel3() {
	hlp := []helpme.HelpMe{
		// {Flag: "", Description: "", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-3", Description: "Apresenta o primeiro exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-3 --resolucao", Description: "Exibe a resolução do primeiro exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-3", Description: "Apresenta o segundo exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-3 --resolucao", Description: "Exibe a resolução do segundo exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-3", Description: "Apresenta o terceiro exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-3 --resolucao", Description: "Exibe a resolução do terceiro exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-3", Description: "Apresenta o quarto exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-4 --nivel-3 --resolucao", Description: "Exibe a resolução do quarto exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-3", Description: "Apresenta o quinto exercício prático do nível 3.", Width: 0},
		{Flag: "--na-pratica-exercicio-5 --nivel-3 --resolucao", Description: "Exibe a resolução do quinto exercício prático do nível 3.", Width: 0},
	}

	fmt.Println("\nCapítulo 7: Exercícios Ninja Nível 3")
	helpme.PrintHelpMe(hlp)
}
