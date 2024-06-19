package structs

import "github.com/fabianoflorentino/aprendago/outline/format"

func StructsEmbutidos() {
	topic := format.OutlineContent{
		Title: "Structs Embutidos",
		Content: `
- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) *e tambem* um competidor (nome, equipe, pontos).
- Go Playground:
    `,
	}

	format.FormatOutlineTopic(topic)
}
