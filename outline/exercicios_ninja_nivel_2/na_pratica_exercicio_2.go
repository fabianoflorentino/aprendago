package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #1",
		Content: `
- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
- ==
- !=
- <=
- <
- >=
- >
- Demonstre estes valores.
- Solução: (https://play.golang.org/p/BMYEch6_s8) ou use no menu do programa (--na-pratica-exercicio-2 --nivel-2 --resolucao)
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio2() {
	a := 10
	b := 20

	fmt.Printf("a == b: %v \n", a == b)
	fmt.Printf("a != b: %v \n", a != b)
	fmt.Printf("a <= b: %v \n", a <= b)
	fmt.Printf("a < b: %v \n", a < b)
	fmt.Printf("a >= b: %v \n", a >= b)
	fmt.Printf("a > b: %v \n", a > b)
}
