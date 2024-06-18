// Pacote exercicios_ninja_nivel_4 contém ferramentas e utilitários
// para os exercícios práticos do Nível 4 do curso Ninja.
package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

// HelpMeExerciciosNinjaNivel4 exibe um menu de ajuda com uma lista
// de flags disponíveis para acessar os exercícios práticos do Nível 4
// e suas resoluções.
func HelpMeExerciciosNinjaNivel4() {
	// Define o menu de ajuda com flags e suas descrições para cada exercício.
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
	}

	// Imprime o menu de ajuda para os exercícios do Nível 4.
	fmt.Println("\nCapítulo 9: Exercícios Ninja Nível 4")
	format.PrintHelpMe(hlp)
}
