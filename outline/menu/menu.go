package outline

import (
	"fmt"

	outline "github.com/fabianoflorentino/aprendago/outline"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Menu(args string) {
	fmt.Println("Aprenda GO")

	options := map[string]func(){
		"--bem-vindo":                    func() { visao_geral_do_curso.BemVindo() },
		"--porque-go":                    func() { visao_geral_do_curso.PorQueGo() },
		"--sucesso":                      func() { visao_geral_do_curso.Sucesso() },
		"--recursos":                     func() { visao_geral_do_curso.Recursos() },
		"--go-playground":                func() { variaveis_valores_tipos.GoPlayground() },
		"--hello-world":                  func() { variaveis_valores_tipos.HelloWorld() },
		"--operador-curto-de-declaracao": func() { variaveis_valores_tipos.OperadorCurtoDeDeclaracao() },
		"--a-palavra-reservada-var":      func() { variaveis_valores_tipos.ApalavraReservadaVar() },
		"--explorando-tipos":             func() { variaveis_valores_tipos.ExplorandoTipos() },
		"--valor-zero":                   func() { variaveis_valores_tipos.ValorZero() },
		"--o-pacote-fmt":                 func() { variaveis_valores_tipos.OpacoteFmt() },
		"--criando-seu-proprio-tipo":     func() { variaveis_valores_tipos.CriandoSeuProprioTipo() },
		"--conversao-nao-coercao":        func() { variaveis_valores_tipos.ConversaoNaoCoercao() },
		"--outline":                      func() { outline.Outline() },
		"--help":                         func() { PrintHelpMe() },
	}

	if action, ok := options[args]; ok {
		action()
	} else {
		fmt.Println("\nOpção inválida")
		PrintHelpMe()
	}
}
