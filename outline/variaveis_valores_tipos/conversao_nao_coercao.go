package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

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
