package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func Constantes() {
	topic := format.OutlineContent{
		Title: "Constantes",
		Content: `
- São valores imutáveis.
- Podem ser tipadas ou não:
  - const oi = "Bom dia"
  - const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
  - Ex. qual o tipo de 42? int? uint? float64?
  - Ou seja, é uma flexibilidade conveniente.
- Na prática: int, float, string.
  - const x = y
  - const ( x = y )
    `,
	}

	format.FormatOutlineTopic(topic)
}
