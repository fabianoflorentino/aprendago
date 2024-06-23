package structs

import "github.com/fabianoflorentino/aprendago/outline/format"

func StructsAnonimos() {
	topic := format.OutlineContent{
		Title: "Structs Anônimos",
		Content: `
- São structs sem identificadores.
- x := struct { name type }{ name: value }
- Go Playground: https://play.golang.org/p/xyhNnSCu1f
    `,
	}

	format.FormatOutlineTopic(topic)
}
