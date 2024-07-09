package exercicios_ninja_nivel_1

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na prática - Exercício #4",
		Content: `
- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
  1. Demonstre o valor da variável "x"
  2. Demonstre o tipo da variável "x"
  3. Atribua 42 à variável "x" utilizando o operador "="
  4. Demonstre o valor da variável "x"
- Para os aventureiros: https://golang.org/ref/spec#Types
- Agora já somos quase ninjas nível 1!
- Solução: https://play.golang.org/p/snm4WuuYmG ou use no menu do programa (--na-pratica-exercicio-4 --nivel-1 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio4() {
	type ninja int

	x := ninja(42)

	resolucao := fmt.Sprintf("%v", x)
	format.FormatResolucaoExercicios(resolucao)

	fmt.Printf("\n%T\n", x)

}
