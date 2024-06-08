package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #1",
		Content: `
- Põe na tela: todos os números de 1 a 10000.
- Solução: https://play.golang.org/p/MkdZiDW8SQ ou use no menu do programa (--na-pratica-exercicio-1 --nivel-3 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	for i := 1; i <= 60; i++ {
		fmt.Printf("%d ", i)
	}
}
