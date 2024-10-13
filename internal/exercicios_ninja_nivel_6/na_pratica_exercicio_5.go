package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio5() {
	topic := format.OutlineContent{
		Title: "Na Prática: Exercício #5",
		Content: `
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
  - Área do círculo: 2 * π * raio
  - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
- Solução: https://play.golang.org/p/qLY-q3vffQ
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio5() {

	q := quadrado{
		lado: 10,
	}

	c := circulo{
		raio: 5,
	}

	fmt.Printf("\nÁrea do quadrado: %f", info(q))
	fmt.Printf("\nÁrea do círculo: %f", info(c))
}

func info(f figura) float64 {
	return f.area()
}

type figura interface {
	area() float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return 2 * 3.14 * c.raio
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}
