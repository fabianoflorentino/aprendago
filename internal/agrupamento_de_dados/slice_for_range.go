package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/pkg/format"

func SliceForRange() {
	topic := format.OutlineContent{
		Title: "Slice: for range",
		Content: `
- Slices:
  - Tamanho: len(x)
  - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
- Go Playground:
  - https://play.golang.org/p/h5-RFJn-Fh
  - https://play.golang.org/p/2wj02m3-eM
    `,
	}

	format.FormatOutlineTopic(topic)
}
