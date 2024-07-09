package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func FuncoesAnonimas() {
	topic := format.OutlineContent{
		Title: "Funções Anônimas",
		Content: `
- Anonymous self-executing functions → Funções anônimas auto-executáveis.
- func(p params) { ... }()
- Vamos ver bastante quando falarmos de goroutines.
- Go Playground: https://play.golang.org/p/Rnqmo6X6jh
    `,
	}

	format.FormatOutlineTopic(topic)
}
