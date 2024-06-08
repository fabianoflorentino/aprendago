package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #2",
		Content: `
- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
	65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
	66
		U+0042 'B'
		U+0042 'B'
		U+0042 'B'
	...e por aí vai.
- Solução: https://play.golang.org/p/QlP6nVchq- ou use no menu do programa (--na-pratica-exercicio-2 --nivel-3 --resolucao)
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio2() {
	for upperLetter := 65; upperLetter <= 90; upperLetter++ {
		fmt.Printf("%d: ", upperLetter)

		for threeTimes := 0; threeTimes < 3; threeTimes++ {
			fmt.Printf(" %#U '%c'", threeTimes, upperLetter)
		}
		fmt.Printf("\n")
	}
}
