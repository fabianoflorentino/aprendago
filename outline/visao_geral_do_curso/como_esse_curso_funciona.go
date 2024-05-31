package outline

import "fmt"

func ComoEsseCursoFunciona() {
	como_esse_curso_funciona := `
- Velocidade de playback.
- Repetição.
- Erros.
- Português vs. inglês
  - Obviamente esse é um curso em português.
  - Mas a língua semi-oficial da programação e do mundo da tecnologia em geral é a língua inglesa.
  - Por isso vamos utilizar bastante termos em inglês ao longo do curso.
  - As explicações serão em português, claro, mas quero que todos fiquem bem confortáveis com os termos em inglês.
  - Desta maneira você será auto-suficiente e poderá procurar por documentação e ajuda por toda a internet, não somente nos sites em português.
  - (Alem de que eu não sei e nem quero saber o nome em português de boa parte dessas coisas :P)
- Constantemente "em progresso."
- Seu feedback é super importante!
  `

	fmt.Printf("Como esse curso funciona?\n %s", como_esse_curso_funciona)
}
