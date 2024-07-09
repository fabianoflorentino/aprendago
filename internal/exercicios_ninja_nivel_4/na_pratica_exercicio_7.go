package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio7() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #7",
		Content: `
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
  - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio7() {
	var resolucao string

	pessoas := [][]string{
		{"Fábio", "Florentino", "Programar"},
		{"Fulano", "de Tal", "Jogar bola"},
		{"Ciclano", "da Silva", "Assistir filmes"},
	}

	for _, pessoa := range pessoas {
		resolucao += fmt.Sprintf("Nome: %v, Sobrenome: %v, Hobby favorito: %v\n", pessoa[0], pessoa[1], pessoa[2])
	}

	format.FormatResolucaoExercicios(resolucao)
}
