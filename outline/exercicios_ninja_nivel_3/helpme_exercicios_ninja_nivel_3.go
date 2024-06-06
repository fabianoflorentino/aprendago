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
	}

	fmt.Println("\nCapítulo 7: Exercícios Ninja Nível 3")
	helpme.PrintHelpMe(hlp)
}
