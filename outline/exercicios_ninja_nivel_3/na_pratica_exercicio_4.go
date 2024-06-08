package outline

import (
	"fmt"
	"time"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #4",
		Content: `
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/dIbfdiuw0ms ou use no menu do programa (--na-pratica-exercicio-4 --nivel-3 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

// ResolucaoNaPraticaExercicio4 - Exercício #4
func ResolucaoNaPraticaExercicio4() {
	var birdthYear int = 1985
	currentYear := time.Now().Year()

	// Nesse formato de loop, é importante quebrar o loop com o break
	for {
		if birdthYear >= currentYear {
			fmt.Printf("%d ", birdthYear)
			// nessário para quebrar o loop
			break
		} else {
			fmt.Printf("%d, ", birdthYear)
			birdthYear++
		}
	}

	fmt.Printf("\nSua idade: %v", time.Now().Year()-1985)
}
