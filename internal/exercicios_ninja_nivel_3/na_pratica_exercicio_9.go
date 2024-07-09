package exercicios_ninja_nivel_3

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio9() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #9",
		Content: `
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio9() {
	var esporteFavorito string = "futebol americano"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Seu esporte favorito é futebol")
	case "futebol americano":
		fmt.Println("Seu esporte favorito é futebol americano")
	case "basquete":
		fmt.Println("Seu esporte favorito é basquete")
	default:
		fmt.Println("Seu esporte favorito não é futebol, futebol americano ou basquete")
	}
}
