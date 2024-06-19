package structs

import "github.com/fabianoflorentino/aprendago/outline/format"

func Struct() {
	topic := format.OutlineContent{
		Title: "Structs",
		Content: `
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
- Go Playground: https://play.golang.org/p/5i0DqxuBp1
    `,
	}

	format.FormatOutlineTopic(topic)
}
