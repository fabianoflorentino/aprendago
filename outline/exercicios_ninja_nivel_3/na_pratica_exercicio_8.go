package exercicios_ninja_nivel_3

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio8() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #8",
		Content: `
- Crie um programa que utilize a declaração switch, sem switch statement especificado.
- Solução: https://play.golang.org/p/TyIGp4Hi8B
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio8() {
	statement := true

	switch {
	case statement:
		fmt.Println("Não deve ser impresso")
	case !statement:
		fmt.Println("Deve ser impresso")
	}
}
