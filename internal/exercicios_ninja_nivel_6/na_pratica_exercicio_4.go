package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio4() {
	topic := format.OutlineContent{
		Title: "Na Prática: Exercício #4",
		Content: `
- Crie um tipo struct "pessoa" que contenha os campos:
  - nome
  - sobrenome
  - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: https://play.golang.org/p/GBZcnu0Ajp
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio4() {

	p := pessoa{
		nome:      "Fabiano",
		sobrenome: "Florentino",
		idade:     39,
	}

	p.MostrarPessoa()
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) MostrarPessoa() {
	fmt.Printf("Nome: %s %s\nIdade: %d\n", p.nome, p.sobrenome, p.idade)
}
