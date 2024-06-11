package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/outline/format"

func SliceMultiDimensional() {
	topic := format.OutlineContent{
		Title: "Slice: Multi Dimensional",
		Content: `
- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
- Go Playground: https://play.golang.org/p/vKyHiG1GtM
- Só pra sacanear: https://play.golang.org/p/ZSU_8eJ9Yp
    `,
	}

	format.FormatOutlineTopic(topic)
}
