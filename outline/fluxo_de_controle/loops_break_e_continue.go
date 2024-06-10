package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/outline/format"

func LoopsBreakEContinue() {
	topic := format.OutlineContent{
		Title: "Loops: break e continue",
		Content: `
- Operação módulo: %
- For: break
- For: continue
- Go Playground: https://play.golang.org/p/gpKMP1wAEM & https://play.golang.org/p/8erMGEbZQix
    `,
	}

	format.FormatOutlineTopic(topic)
}
