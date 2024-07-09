package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio9() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #9",
		Content: `
- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
- Solução: https://play.golang.org/p/3fcvHlt8Lm
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio9() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
		"de_tal_ciclano":   {"Programar 3", "Jogar bola 3", "Assistir filmes 3"},
	}

	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}
