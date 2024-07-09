package exercicios_ninja_nivel_1

import "github.com/fabianoflorentino/aprendago/pkg/format"

func ContribuaSeuCodigo() {
	topic := format.OutlineContent{
		Title: "Contribua seu código",
		Content: `
- Nesse curso a gente vai fazer um monte de exercícios.
  - Talvez você queira contribuir suas próprias soluções.
  - Talvez você tenha exemplos melhores que os que estamos mostrando aqui.
- Para compartilhar me manda o link no twitter do seu código no Go Playground, twitter.com/ellenkorbes!
- "@ellenkorbes Olha essa solução pro exercício Ninja nível 5, exercício 2: <link> O que vc acha?"
  `,
	}

	format.FormatOutlineTopic(topic)
}
