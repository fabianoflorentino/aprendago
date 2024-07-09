package exercicios_ninja_nivel_3

import (
	"fmt"
	"time"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #3",
		Content: `
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/qnFjiDJzLor ou use no menu do programa (--na-pratica-exercicio-3 --nivel-3 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	var birdthYear int = 1985
	currentYear := time.Now().Year()

	// fmt.Printf("\nDigite o ano de seu nascimento: ")
	// fmt.Scan(&birdthYear)

	for year := birdthYear; year <= currentYear; year++ {
		if year == currentYear {
			fmt.Printf("%d\n", year)
		} else {
			fmt.Printf("%d, ", year)
		}
	}

	fmt.Printf("Sua idade: %v", currentYear-birdthYear)
}
