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
		Nome             string
		Idade            int
		Amigos           map[string]int
		ComidasFavoritas []string
	}{
		Nome:  "Fabiano",
		Idade: 39,
		Amigos: map[string]int{
			"Ale":   46,
			"Lucas": 38,
		},
		ComidasFavoritas: []string{
			"Pizza",
			"Lasanha",
			"Hamburguer",
		},
	}

	// mais para frente
	// jsonData, _ := json.Marshal(pessoa)
	// fmt.Println(string(jsonData))

	fmt.Println(pessoa)
}
