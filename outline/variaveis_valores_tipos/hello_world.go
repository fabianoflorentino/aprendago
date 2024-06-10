package variaveis_valores_tipos

import format "github.com/fabianoflorentino/aprendago/outline/format"

func HelloWorld() {
	topic := format.OutlineContent{
		Title: "Hello, World!",
		Content: `
- Estrutura básica:
  - package main.
  - func main: é aqui que tudo começa, é aqui que tudo acaba.
  - import.
- Packages:
  - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
  - Notação: pacote.Identificador. Exemplo: fmt.Println()
  - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.
    `,
	}

	format.FormatOutlineTopic(topic)
}
