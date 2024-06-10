package variaveis_valores_tipos

import "github.com/fabianoflorentino/aprendago/outline/format"

func ApalavraReservadaVar() {
	topic := format.OutlineContent{
		Title: "A palavra reservada var",
		Content: `
- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos "var"
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses
		`,
	}

	format.FormatOutlineTopic(topic)
}
