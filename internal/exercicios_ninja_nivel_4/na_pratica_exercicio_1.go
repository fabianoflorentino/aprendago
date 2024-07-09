package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #1",
		Content: `
- Usando uma literal composta:
  - Crie um array que suporte 5 valores to tipo int
  - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
- Solução: https://play.golang.org/p/tpWDzzlO2l
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	array := [5]int{1, 2, 3, 4, 5}
	var resolucao string

	for _, value := range array {
		resolucao += fmt.Sprintf("%v, ", value)
	}

	resolucao = strings.Trim(resolucao, ", ")
	format.FormatResolucaoExercicios(resolucao)
}
