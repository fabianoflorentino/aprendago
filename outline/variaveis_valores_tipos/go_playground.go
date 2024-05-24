package outline

import "fmt"

func GoPlayground() {
	go_playground := `
- É online, funciona sem instalar nem configurar nada.
- Assim você pode começar a programar o mais rápido possível.
- Mais pra frente no curso vou explicar direitinho como configurar tudo no seu computador.
- [Go Playground](https://play.golang.org/).
  - Função share. Use para compartilhar código, por exemplo pra fazer uma pergunta em um fórum.
  - Função imports.
  - Função format.
    - Maneira idiomática: a gente fala da mesma maneira que a comunidade onde estamos.
  - Função run.
  `

	fmt.Println("\nGo Playground")
	fmt.Println(go_playground)
}
