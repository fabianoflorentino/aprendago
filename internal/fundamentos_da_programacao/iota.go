package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Iota() {
	topic := format.OutlineContent{
		Title: "Iota",
		Content: `
- golang.org/ref/spec
- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
  - iota, iota + 1, a = iota b c, reinicia em cada const, _
- Go Playground: https://play.golang.org/p/eSrwoQjuYR
    `,
	}

	format.FormatOutlineTopic(topic)
}
