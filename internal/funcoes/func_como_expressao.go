package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func FuncComoExpressa() {
	topic := format.OutlineContent{
		Title: "Funções como expressões",
		Content: `
- f := func(p params){ ... }
- f()
- Go Playground: https://play.golang.org/p/cPxhPUbfLy
    `,
	}

	format.FormatOutlineTopic(topic)
}
