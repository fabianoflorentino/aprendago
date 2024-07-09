package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/pkg/format"

func CondicionaisADeclaracaoIf() {
	topic := format.OutlineContent{
		Title: "Condicionais: A declaração if",
		Content: `
- If: bool
- If: o operador não → "!"
- If: declaração de inicialização
- Go Playground: https://play.golang.org/p/6nq2Tjb07i
    `,
	}

	format.FormatOutlineTopic(topic)
}
