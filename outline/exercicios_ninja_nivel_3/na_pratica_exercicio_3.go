package outline

import (
	"fmt"
	"time"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #3",
		Content: `
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/qnFjiDJzLor
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
