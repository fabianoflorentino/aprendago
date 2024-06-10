package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/outline/format"

func EntendendoFluxoDeControle() {
	topic := format.OutlineContent{
		Title: "Entendendo Fluxo de Controle",
		Content: `
- Computadores lêem programas de uma certa maneira, do mesmo jeito que nós lemos livros, por exemplo, de uma certa maneira.
- Quando nós ocidentais lemos livros, lemos da frente pra trás, da esquerda pra direito, de cima pra baixo.
- Computadores lêem de cima pra baixo.
- Ou seja, sua leitura é sequencial. Isso chama-se fluxo de controle sequencial.
- Alem do fluxo de controle sequencial, há duas declarações que podem afetar como o computador lê o código:
  - Uma delas é o fluxo de controle de repetição (loop). Nesse caso, o computador vai repetir a leitura de um mesmo código de uma maneira específica. O fluxo de controle de repetição tambem é conhecido como fluxo de controle iterativo.
  - E o outro é o fluxo de controle condicional, ou fluxo de controle de seleção. Nesse caso o computador encontra uma condição e, através de uma declaração if ou switch, toma um curso ou outro dependendo dessa condição.
- Ou seja, há três tipos de fluxo de controle: sequencial, de repetição e condicional.
- Nesse capítulo:
    - Sequencial
    - Iterativo (loop)
      - for: inicialização, condição, pós
      - for: hierarquicamente
      - for: condição ("while")
      - for: ...ever?
      - for: break
      - for: continue
    - Condicional
      - declarações switch/case/default
        - não há fall-through por padrão
        - criando fall-through
        - default
        - múltiplos casos
        - casos podem ser expressões
            - se resultarem em true, rodam
        - tipo
        - if
          - bool
          - o operador "!"
          - declaração de inicialização
          - if, else
          - if, else if, else
          - if, else if, else if, ..., else
    `,
	}

	format.FormatOutlineTopic(topic)
}
