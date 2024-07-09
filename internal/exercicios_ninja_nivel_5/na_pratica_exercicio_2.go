package exercicios_ninja_nivel_5

import (
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #2",
		Content: `
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio2() {
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

	pessoas := map[string]pessoa{
		p1.sobrenome: p1,
		p2.sobrenome: p2,
	}

	for k, p := range pessoas {
		println(k)
		for _, s := range p.sabores_de_sorvete {
			println("\t", s)
		}
	}
}
