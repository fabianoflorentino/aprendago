package variaveis_valores_tipos

import "github.com/fabianoflorentino/aprendago/pkg/format"

func ValorZero() {
	topic := format.OutlineContent{
		Title: "Valor zero",
		Content: `
- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
  - ints: 0
  - floats: 0.0
  - booleans: false
  - strings: ""
  - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.
    `,
	}

	format.FormatOutlineTopic(topic)
}
