package fluxo_de_controle

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func LoopsUtilizandoAscii() {
	topic := format.OutlineContent{
		Title: "Loops: Utilizando ASCII",
		Content: `
- Desafio surpresa!
- Format printing:
	- Decimal       %d
	- Hexadecimal   %#x
	- Unicode       %#U
	- Tab           \t
	- Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
- Solução: https://play.golang.org/p/REm2WHyzzz
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoLoopsUtilizandoAscii() {
	for i := 33; i <= 122; i++ {
		if i%10 == 0 {
			println()
		}

		fmt.Printf("%d: %c\t", i, i)
	}
}
