package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #3",
		Content: `
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
- Solução: https://play.golang.org/p/eWnKI59ual ou use no menu do programa (--na-pratica-exercicio-3 --nivel-2 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	const (
		untypedConst     = 10
		typedConst   int = 20
	)

	fmt.Printf("untypedConst: %v \ntypedConst: %v", untypedConst, typedConst)
}
