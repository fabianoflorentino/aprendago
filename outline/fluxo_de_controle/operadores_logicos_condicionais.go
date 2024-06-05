package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func OperadoresLogicosCondicionais() {
	topic := format.OutlineContent{
		Title: "Operadores lógicos condicionais",
		Content: `
- &&
- ||
- !
- Go Playground: https://play.golang.org/p/MFwrt93xlc
- Qual o resultado de fmt.Println...
  - true && true
  - true && false
  - true || true
  - true || false
  - !true
    `,
	}

	format.FormatOutlineTopic(topic)
}
