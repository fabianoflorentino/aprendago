package outline

import "fmt"

func ConversaoNaoCoercao() {

	conversao_nao_coercao := `
- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.
  `

	fmt.Println("Conversão não coerção")
	fmt.Println(conversao_nao_coercao)
}
