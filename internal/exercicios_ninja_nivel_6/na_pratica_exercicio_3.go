package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício 3",
		Content: `
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: https://play.golang.org/p/b5Ua2kNN8a
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	defer fmt.Println("Execução do defer ocorreu ao final do contexto.")

	fmt.Println("Execução do defer ocorre ao final do contexto ao qual ela pertence.")
}
