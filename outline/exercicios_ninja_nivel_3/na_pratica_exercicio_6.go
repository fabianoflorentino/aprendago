package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio6() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #6",
		Content: `
- Crie um programa que demonstre o funcionamento da declaração if.
- Solução: https://play.golang.org/p/OGPgTJZbiy
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio6() {
	fmt.Printf("10 é maior que 5? ")
	if 10 > 5 {
		fmt.Print("Sim")
	}
}
