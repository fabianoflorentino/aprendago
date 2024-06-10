package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/outline/format"

func SliceLiteralComposta() {
	topic := format.OutlineContent{
		Title: "Slice: Literal Composta",
		Content: `
- O que são tipos de dados compostos?
  - Wikipedia: Composite_data_type
  - Effective Go: Composite literals
  - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}
- Go Playground: https://play.golang.org/p/W7Cxm8NPZC
    `,
	}

	format.FormatOutlineTopic(topic)
}
