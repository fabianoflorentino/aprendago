package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MapRangeEDeletando() {
	topic := format.OutlineContent{
		Title: "Map Range e Deletando",
		Content: `
- Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
- Go Playground: https://play.golang.org/p/6zEMfIP-AE
- delete(map, key)
- Deletar uma key não-existente não retorna erros!
- Go Playground: https://play.golang.org/p/0uuIicU3Zz
    `,
	}

	format.FormatOutlineTopic(topic)
}
