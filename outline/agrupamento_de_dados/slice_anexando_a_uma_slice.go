package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/outline/format"

func SliceAnexandoASlice() {
	topic := format.OutlineContent{
		Title: "Slice: Anexando a uma slice",
		Content: `
- Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → desdobrar, desenrolar
- Nome oficial: enumeration
- Go Playground: https://play.golang.org/p/RpkDCTumpT
    `,
	}

	format.FormatOutlineTopic(topic)
}
