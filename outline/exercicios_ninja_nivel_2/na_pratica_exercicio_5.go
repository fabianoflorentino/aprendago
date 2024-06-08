package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio5() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #5",
		Content: `
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
- Solução: https://play.golang.org/p/RkpqPpRWuo ou use no menu do programa (--na-pratica-exercicio-5 --nivel-2 --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio5() {
	v1 := `Aprenda Go!`

	fmt.Println(v1)
}
