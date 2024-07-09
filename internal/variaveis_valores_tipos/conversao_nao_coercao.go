package variaveis_valores_tipos

import "github.com/fabianoflorentino/aprendago/pkg/format"

func ConversaoNaoCoercao() {
	topic := format.OutlineContent{
		Title: "Conversão não coerção",
		Content: `
- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.
		`,
	}

	format.FormatOutlineTopic(topic)
}
