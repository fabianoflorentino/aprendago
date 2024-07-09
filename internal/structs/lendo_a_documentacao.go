package structs

import "github.com/fabianoflorentino/aprendago/pkg/format"

func LendoADocumentacao() {
	topic := format.OutlineContent{
		Title: "Lendo a documentação",
		Content: `
- É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
  - ref/spec
    - Já vimos mais da metade dos tipos em Go!
    - Struct types.
      - x, y int
      - anonymous fields
      - promoted fields
- Go Playground: https://play.golang.org/p/z9UQej4IQT
    `,
	}

	format.FormatOutlineTopic(topic)
}
