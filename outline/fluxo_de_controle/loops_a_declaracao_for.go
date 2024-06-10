package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/outline/format"

func LoopsADeclaracaoFor() {
	topic := format.OutlineContent{
		Title: "Loops: A declaração for",
		Content: `
- For: inicialização, condição, pós
- For: condição ("while")
- For: ...ever? (http servers)
- For: break
- https://golang.org/ref/spec#For_statements, Effective Go
- (Range vem mais pra frente.)
    `,
	}

	format.FormatOutlineTopic(topic)
}
