package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio5() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #5",
		Content: `
- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]
- Solução: https://play.golang.org/p/26bT-UKmJH
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio5() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	novoSlice := append(slice[:3], slice[6:]...)

	resolucao := fmt.Sprintf("%v", novoSlice)

	format.FormatResolucaoExercicios(resolucao)
}
