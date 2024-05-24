package outline

import (
	"fmt"

	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

const HELPME = `
Opções:

  --bem-vindo                 Exibe a mensagem de boas-vindas do curso
  --porque-go                 Exibe a mensagem sobre por que aprender Go
  --sucesso                   Exibe a mensagem sobre sucesso
  --recursos                  Exibe os recursos do curso
  --como-esse-curso-funciona  Exibe como esse curso funciona
  --outline                   Exibe o outline do curso
  --help                      Exibe a lista de opções
`

func Menu(args string) {
	fmt.Println("Aprenda GO")

	options := map[string]func(){
		"--bem-vindo": func() { visao_geral_do_curso.BemVindo() },
		"--porque-go": func() { visao_geral_do_curso.PorQueGo() },
		"--sucesso":   func() { visao_geral_do_curso.Sucesso() },
		"--recursos":  func() { visao_geral_do_curso.Recursos() },
		"--outline":   func() { Outline() },
		"--help":      func() { fmt.Print(HELPME) },
	}

	if _, ok := options[args]; ok {
		options[args]()
	} else {
		fmt.Println("\nOpção inválida")
		fmt.Print(HELPME)
	}
}
