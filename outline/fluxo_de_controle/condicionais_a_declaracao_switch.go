package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/outline/format"

func CondicionaisADeclaracaoSwitch() {
	topic := format.OutlineContent{
		Title: "Condicionais: A declaração switch",
		Content: `
- Switch:
- pode avaliar uma expressão
  - switch statement == case (value)
  - default switch statement == true (bool)
- não há fall-through por padrão
- criando fall-through
- default
- cases compostos
    `,
	}

	format.FormatOutlineTopic(topic)
}
