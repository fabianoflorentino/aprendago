package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func HelMeExerciciosNinjaNivel6() {
	hlp := []format.HelpMe{
		{Flag: "--na-pratica-exercicio-1 --nivel-6", Description: "Apresenta o primeiro exercício prático do Nível 6.", Width: 0},
		{Flag: "--na-pratica-exercicio-1 --nivel-6 --resolucao", Description: "Apresenta a resolução do primeiro exercício prático do Nível 6.", Width: 0},
	}

	fmt.Println("\nCapítulo 12: Exercícios Ninja Nível 6")
	format.PrintHelpMe(hlp)
}
