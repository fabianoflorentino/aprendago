package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/outline/format"

func LoopsNestedLoop() {
	topic := format.OutlineContent{
		Title: "Loops: Nested Loop (repetição hierárquica)",
		Content: `
- For
  - Repetição hierárquica
  - Exemplos: relógio, calendário
    `,
	}

	format.FormatOutlineTopic(topic)
}
