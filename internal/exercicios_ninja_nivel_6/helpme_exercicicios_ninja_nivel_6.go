package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func HelMeExerciciosNinjaNivel6() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-6", Description: "Apresenta o primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-6 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-6", Description: "Apresenta o segundo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-2 --nivel-6 --resolucao", Description: "Apresenta a resolução do segundo exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-6", Description: "Apresenta o terceiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-3 --nivel-6 --resolucao", Description: "Apresenta a resolução do terceiro exercício prático do Nível 6.", Width: 0},
	}

	fmt.Println("\nCapítulo 13: Exercícios Ninja Nível 6")
	format.PrintHelpMe(hlp)
}
