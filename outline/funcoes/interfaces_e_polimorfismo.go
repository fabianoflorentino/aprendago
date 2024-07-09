package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func InterfacesEPolimorfismo() {
	topic := format.OutlineContent{
		Title: "Interfaces e Polimorfismo",
		Content: `
- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo *e também* o tipo da interface.
    `,
	}

	format.FormatOutlineTopic(topic)
}
