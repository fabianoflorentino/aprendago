package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na pratica: Exercício #3",
		Content: `
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
  - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
  - Do quinto ao último item do slice (incluindo o último item!)
  - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
  - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
  - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
- Solução: https://play.golang.org/p/1aPXVeR1mf
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {
		if value != len(slice) {
			fmt.Printf("%v, ", value)
		} else {
			fmt.Printf("%v", value)
		}
	}

	fmt.Printf("\nTipo: %T", slice)
	fmt.Printf("\n\n %v", slice[0:4])
}
