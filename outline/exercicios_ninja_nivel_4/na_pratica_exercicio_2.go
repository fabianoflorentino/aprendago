package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na pratica: Exercício #2",
		Content: `
- Usando uma literal composta:
  - Crie uma slice de tipo int
  - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
- Solução: https://play.golang.org/p/ST3TKusuOd
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {
		if value == len(slice) {
			fmt.Printf("%v", value)
		} else {
			fmt.Printf("%v, ", value)
		}
	}

	fmt.Printf("\nTipo: %T", slice)
}
