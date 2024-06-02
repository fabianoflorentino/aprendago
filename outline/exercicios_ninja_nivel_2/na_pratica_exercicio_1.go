package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #1",
		Content: `
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
- Solução: https://play.golang.org/p/X7qm3aWSa6
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	numero := 10000000

	fmt.Printf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", numero, numero, numero)
}
