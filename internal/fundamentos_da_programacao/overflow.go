package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Overflow() {
	topic := format.OutlineContent{
		Title: "Overflow",
		Content: `
- Um uint16, por exemplo, vai de 0 a 65535.
- Que acontece se a gente tentar usar 65536?
- Ou se a gente estiver em 65535 e tentar adicionar mais 1?
- Playground: https://play.golang.org/p/t7Z4m127F2t
    `,
	}

	format.FormatOutlineTopic(topic)
}
