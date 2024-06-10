package exercicios_ninja_nivel_2

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio6() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #6",
		Content: `
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre-a.
Solução: (https://play.golang.org/RkpqPpRWuo) ou use no menu do programa (--na-pratica-exercicio-6 --nivel-2 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio6() {
	const (
		ano = 2024 + iota
		v1
		v2
		v3
		v4
	)

	fmt.Printf("v1: %v\nv2: %v\nv3: %v\nv4: %v\n", v1, v2, v3, v4)
}
