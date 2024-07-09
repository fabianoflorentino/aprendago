package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na prática - Exercício #2",
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

	resolucao := fmt.Sprintf("a: %v \nb: %v \n\na == b: %v \na != b: %v \na <= b: %v \na < b: %v \na >= b: %v \na > b: %v \n", a, b, a == b, a != b, a <= b, a < b, a >= b, a > b)

	format.FormatResolucaoExercicios(resolucao)
}
