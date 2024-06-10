package exercicios_ninja_nivel_2

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #4",
		Content: `
- Crie um programa que:
- Atribua um valor int a uma variável
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
- Demonstre esta outra variável em decimal, binário e hexadecimal
- Solução: (https://play.golang.org/p/IiwgT0v3Mp) ou use no menu do programa (--na-pratica-exercicio-4 --nivel-2 --resolucao
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio4() {
	v1 := 10

	fmt.Printf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n\n", v1, v1, v1)

	v2 := v1 << 1

	fmt.Printf("Decimal: %d \nBinário: %b \nHexadecimal: %#x", v2, v2, v2)
}
