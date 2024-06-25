package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #4",
		Content: `
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio4() {
	pessoa := struct {
		nome             string
		idade            int
		amigos           map[string]int
		comidasFavoritas []string
	}{
		nome:  "Fabiano",
		idade: 33,
		amigos: map[string]int{
			"João":  30,
			"Maria": 25,
		},
		comidasFavoritas: []string{
			"Pizza",
			"Lasanha",
			"Hamburguer",
		},
	}

	fmt.Println(pessoa)
}
