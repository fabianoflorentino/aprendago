package exercicios_ninja_nivel_1

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio5() {
	topic := format.OutlineContent{
		Title: "Na prática - Exercício #5",
		Content: `
- Utilizando a solução do exercício anterior:
  1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
  2. Na função main:
    1. Isto já deve estar feito:
      1. Demonstre o valor da variável "x"
      2. Demonstre o tipo da variável "x"
      3. Atribua 42 à variável "x" utilizando o operador "="
      4. Demonstre o valor da variável "x"
    2. Agora faça tambem:
      1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
      2. Demonstre o valor de "y"
      3. Demonstre o tipo de "y"
- Solução: https://play.golang.org/p/uq81T_fasP ou use no menu do programa (--na-pratica-exercicio-5 --nivel-1 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio5() {
	type ninja int

	var x ninja
	var y int

	x = 42
	y = int(x)

	resolucao := fmt.Sprintf("%T\n%T", x, y)

	format.FormatResolucaoExercicios(resolucao)
}
