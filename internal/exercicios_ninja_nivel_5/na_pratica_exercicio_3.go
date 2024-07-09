package exercicios_ninja_nivel_5

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio3() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #3",
		Content: `
- Crie um novo tipo: veículo
  - O tipo subjacente deve ser struct
  - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
  - Os tipos subjacentes devem ser struct
  - Ambos devem conter "veículo" como struct embutido
  - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
  - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
  - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
  - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
- Solução: https://play.golang.org/p/3eoGb0kxzT
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio3() {
	type veiculo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}

	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	cnt := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preto",
		},

		tracaoNasQuatro: true,
	}

	sdn := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "Branco",
		},

		modeloLuxo: true,
	}

	fmt.Printf("Caminhonete: %+v \n", cnt)
	fmt.Printf("Sedan: %+v \n", sdn)

	fmt.Println("Portas da caminhonete:", cnt.portas)
	fmt.Println("Cor do sedan:", sdn.cor)
}
