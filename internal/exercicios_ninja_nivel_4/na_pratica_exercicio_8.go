package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio8() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #8",
		Content: `
- Crie um map com key tipo string e value tipo []string.
  - Key deve conter nomes no formato sobrenome_nome
  - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio8() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
	}

	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}
