package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio6() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #6",
		Content: `
- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
      "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
      "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
      "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*
- Solução: https://play.golang.org/p/cGYBphlyCE
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio6() {
	estados := make([]string, 26)

	copy(estados, []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins",
	})

	editEstados := strings.Join(estados, ", ")

	resolucao := fmt.Sprintf("len: %v cap: %v\nEstados: %v", len(estados), cap(estados), editEstados)

	format.FormatResolucaoExercicios(resolucao)
}
