package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Sintaxe() {
	topic := format.OutlineContent{
		Title: "Sintaxe",
		Content: `
- Qual a utilidade de funções?
  - Abstrair funcionalidade
  - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
  - Funções são definidas com parâmetros
  - Funções são chamadas com argumentos
- Tudo em Go é *pass by value.*
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico.
- Exemplos:
  - Função básica.
    - Go Playground: https://play.golang.org/p/FebJblBenP
  - Função que aceita um argumento.
    - Go Playground: https://play.golang.org/p/CE6Ij3U4QB
  - Função com retorno.
    - Go Playground: https://play.golang.org/p/gKxwYe6btP
  - Função com múltiplos retornos e parâmetro variádico.
    - Go Playground: https://play.golang.org/p/OcQ1wXwM2c
  - Mais um: https://play.golang.org/p/8wc2TA9xH_
    `,
	}

	format.FormatOutlineTopic(topic)
}
