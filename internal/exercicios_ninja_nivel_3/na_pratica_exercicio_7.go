package exercicios_ninja_nivel_3

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio7() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #7",
		Content: `
- Utilizando a solução anterior, adicione as opções else if e else.
- Solução: https://play.golang.org/p/_cO7E-yL0o
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio7() {
	fmt.Printf("10 é maior que 5? ")

	if 10 > 5 {
		fmt.Print("Sim")
	} else if 5 > 10 {
		fmt.Print("Não")
	} else {
		fmt.Print("Não sei")
	}
}
