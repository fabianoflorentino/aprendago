package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio5() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #5",
		Content: `
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
- Solução: https://play.golang.org/p/zcEsXqnBr8 ou use no menu do programa (--na-pratica-exercicio-5 --nivel-3 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio5() {
	fmt.Printf("O resto da divisão por 4 de todos os números entre 10 e 100\n\n")
	for i := 10; i <= 100; i++ {
		if i == 100 {
			fmt.Printf("%d: %d ", i, i%4)
		} else {
			fmt.Printf("%d: %d, ", i, i%4)
		}
	}
}
