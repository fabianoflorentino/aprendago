package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #3",
		Content: `
- Utilizando a solução do exercício anterior:
  1. Em package-level scope, atribua os seguintes valores às variáveis:
    1. para "x" atribua 42
    2. para "y" atribua "James Bond"
    3. para "z" atribua true
  2. Na função main:
    1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
      Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
    2. Demonstre a variável "s".
- Solução: https://play.golang.org/p/QFctSQB_h3
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)
}
