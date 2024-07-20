package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Closure() {
	topic := format.OutlineContent{
		Title: "Closure",
		Content: `
- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
  - Package-level scope
  - Function-level scope
  - Code-block-in-code-block scope
- Exemplo de closure:
  - func i() func() int { x := 0; return func() int { x++; return x } }
  - Quando fizermos a := i() teremos um scope, um valor para x.
  - Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
- Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.
- Go Playground: https://play.golang.org/p/AdFciYwI2Z
    `,
	}

	format.FormatOutlineTopic(topic)
}
