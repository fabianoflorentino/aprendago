package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func RetornandoUmaFuncao() {
	topic := format.OutlineContent{
		Title: "Retornando uma função",
		Content: `
- Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
  - ????: fmt.Println(f()())
- Go Playground: https://play.golang.org/p/zPjoWNrCJF
    `,
	}

	format.FormatOutlineTopic(topic)
}
