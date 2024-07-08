package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func Metodos() {
	topic := format.OutlineContent{
		Title: "Métodos",
		Content: `
- Um método é uma função anexada a um tipo.
- Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
- Pode-se anexar uma função a um tipo utilizando seu receiver.
- Utilização: valor.método()
- Exemplo: o tipo "pessoa" pode ter um método oibomdia()
- Go Playground: https://play.golang.org/p/tQtoqUBpY5
    `,
	}

	format.FormatOutlineTopic(topic)
}
