package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MapsIntroducao() {
	topic := format.OutlineContent{
		Title: "Maps: introdução",
		Content: `
- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
  - v, ok := m[key]
  - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps *não tem ordem.*
- Go Playground: https://play.golang.org/p/JXDdJan8Ev
    `,
	}

	format.FormatOutlineTopic(topic)
}
