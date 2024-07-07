package funcoes

import "github.com/fabianoflorentino/aprendago/outline/format"

func DesenrolandoEnumerandoUmaSlice() {
	topic := format.OutlineContent{
		Title: "Desenrolando (enumerando) uma slice",
		Content: `
- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
  - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
    - Go Playground: https://play.golang.org/p/k8O3__8UDa
  - Pode-se passar *zero* ou mais valores
    - Go Playground: https://play.golang.org/p/C238I9n7Vs
  - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
    - Go Playground: https://play.golang.org/p/8wc2TA9xH_
    - Não roda: https://play.golang.org/p/2qTAnLWfgB
    `,
	}

	format.FormatOutlineTopic(topic)
}
