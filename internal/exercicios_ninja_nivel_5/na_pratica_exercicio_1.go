package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #1",
		Content: `
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
  - Nome
  - Sobrenome
  - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: https://play.golang.org/p/Pyp7vmTJfY
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	type pessoa struct {
		nome               string
		sobrenome          string
		sabores_de_sorvete []string
	}

	p1 := pessoa{
		nome:               "Fulano",
		sobrenome:          "de Tal",
		sabores_de_sorvete: []string{"Chocolate", "Morango", "Baunilha"},
	}

	p2 := pessoa{
		nome:               "Ciclano",
		sobrenome:          "da Silva",
		sabores_de_sorvete: []string{"Pistache", "Creme", "Coco"},
	}

	pessoas := []pessoa{p1, p2}

	for _, p := range pessoas {
		fmt.Println(p.nome, p.sobrenome)
		for _, s := range p.sabores_de_sorvete {
			fmt.Printf("\t%v\n", s)
		}
	}
}
