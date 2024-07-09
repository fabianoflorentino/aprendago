package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/pkg/format"

func TipoString() {
	topic := format.OutlineContent{
		Title: "Tipo string",
		Content: `
- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
  - %v %T
  - Raw string literals
  - Conversão para slice of bytes: []byte(x)
  - %#U, %#x
  - Go Playground: https://play.golang.org/p/dt2x1ies5b & https://play.golang.org/p/PpDnspiyA_7
- https://blog.golang.org/string
    `,
	}

	format.FormatOutlineTopic(topic)
}
