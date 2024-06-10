package exercicios_ninja_nivel_1

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #4",
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

	var x ninja

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
