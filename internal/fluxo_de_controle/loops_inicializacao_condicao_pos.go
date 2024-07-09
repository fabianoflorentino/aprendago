package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/pkg/format"

func LoopsInicializacaoCondicaoPos() {
	topic := format.OutlineContent{
		Title: "Loops: Inicialização, Condição e Pós",
		Content: `
- For
  - Inicialização, condição, pós
  - Ponto e vírgula?
  - https://gobyexample.com
  - Não existe while!
    `,
	}

	format.FormatOutlineTopic(topic)
}
