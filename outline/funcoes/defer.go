package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func Defer() {
	topic := format.OutlineContent{
		Title: "Defer",
		Content: `
- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
- Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
- "Deixa pra última hora!"
- ref/spec
- Sempre usamos para fechar um arquivo após abri-lo.
- Go Playground: https://play.golang.org/p/sFj8arw0E_
    `,
	}

	format.FormatOutlineTopic(topic)
}
