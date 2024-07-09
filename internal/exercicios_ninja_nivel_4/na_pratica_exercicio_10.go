package exercicios_ninja_nivel_4

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio10() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #10",
		Content: `
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
- Solução:
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio10() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
		"de_tal_ciclano":   {"Programar 3", "Jogar bola 3", "Assistir filmes 3"},
	}

	// Removendo uma entrada do map
	delete(pessoas, "da_silva_ciclano")

	// Exibindo o map inteiro utilizando range loop
	// O range loop retorna a chave e o valor do map em cada iteração do loop
	// A chave é o sobrenome_nome e o valor é o slice de hobbies
	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}
