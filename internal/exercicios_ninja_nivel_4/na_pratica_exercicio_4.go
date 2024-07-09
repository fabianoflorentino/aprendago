package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #4",
		Content: `
- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.
- Solução: https://play.golang.org/p/6WNJ0Otpy0
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio4() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	append52 := append(slice, 52)
	append53to55 := append(append52, 53, 54, 55)

	sliceY := []int{56, 57, 58, 59, 60}
	appendSliceY := append(append53to55, sliceY...)

	resolucao := fmt.Sprintf("append52: %v\nappend53to55: %v\nappendSliceY: %v", append52, append53to55, appendSliceY)

	format.FormatResolucaoExercicios(resolucao)
}
