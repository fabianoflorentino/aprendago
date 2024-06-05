package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func CondicionaisIfElseIfElse() {
	topic := format.OutlineContent{
		Title: "Condicionais: A declaração if else if else",
		Content: `
- If, else.
- If, else if, else.
- If, else if, else if, ..., else.
- Go Playground: https://play.golang.org/p/18VrRX2pec
    `,
	}

	format.FormatOutlineTopic(topic)
}
