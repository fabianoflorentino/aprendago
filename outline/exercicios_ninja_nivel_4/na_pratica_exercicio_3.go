package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #3",
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
	var newSlice []string

	for _, value := range slice {
		if value != len(slice) {
			newSlice = append(newSlice, fmt.Sprintf("%v, ", value))
		} else {
			newSlice = append(newSlice, fmt.Sprintf("%v", value))
		}
	}
	joinSlice := strings.Join(newSlice[0:4], "")
	trimSlice := strings.Trim(joinSlice, ", ")

	resolucao := fmt.Sprintf("\nSlice: %v\nTipo: %T", trimSlice, newSlice)

	format.FormatResolucaoExercicios(resolucao)
}
