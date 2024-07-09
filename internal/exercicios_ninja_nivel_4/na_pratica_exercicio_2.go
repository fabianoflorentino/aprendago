package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #2",
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
	var newSlice []string

	for _, value := range slice {
		if value != len(slice) {
			newSlice = append(newSlice, fmt.Sprintf("%v, ", value))
		} else {
			newSlice = append(newSlice, fmt.Sprintf("%v", value))
		}
	}

	joinSlice := strings.Join(newSlice, "")
	resolucao := fmt.Sprintf("\nSlice: %v\nTipo: %T", joinSlice, newSlice)

	format.FormatResolucaoExercicios(resolucao)
}
