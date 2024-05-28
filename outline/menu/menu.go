package outline

import (
	"fmt"
	"strings"

	outline "github.com/fabianoflorentino/aprendago/outline"
	exercicios_ninja_nivel_1 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Menu(args []string) {
	fmt.Println("Aprenda GO")

	options := map[string]func(){
		"--bem-vindo":                          func() { visao_geral_do_curso.BemVindo() },
		"--porque-go":                          func() { visao_geral_do_curso.PorQueGo() },
		"--sucesso":                            func() { visao_geral_do_curso.Sucesso() },
		"--recursos":                           func() { visao_geral_do_curso.Recursos() },
		"--go-playground":                      func() { variaveis_valores_tipos.GoPlayground() },
		"--hello-world":                        func() { variaveis_valores_tipos.HelloWorld() },
		"--operador-curto-de-declaracao":       func() { variaveis_valores_tipos.OperadorCurtoDeDeclaracao() },
		"--a-palavra-reservada-var":            func() { variaveis_valores_tipos.ApalavraReservadaVar() },
		"--explorando-tipos":                   func() { variaveis_valores_tipos.ExplorandoTipos() },
		"--valor-zero":                         func() { variaveis_valores_tipos.ValorZero() },
		"--o-pacote-fmt":                       func() { variaveis_valores_tipos.OpacoteFmt() },
		"--criando-seu-proprio-tipo":           func() { variaveis_valores_tipos.CriandoSeuProprioTipo() },
		"--conversao-nao-coercao":              func() { variaveis_valores_tipos.ConversaoNaoCoercao() },
		"--contribua-seu-codigo":               func() { exercicios_ninja_nivel_1.ContribuaSeuCodigo() },
		"--na-pratica-exercicio-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio1() },
		"--na-pratica-exercicio-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio1() },
		"--na-pratica-exercicio-2":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio2() },
		"--na-pratica-exercicio-2 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio2() },
		"--na-pratica-exercicio-3":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio3() },
		"--na-pratica-exercicio-3 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio3() },
		"--na-pratica-exercicio-4":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio4() },
		"--na-pratica-exercicio-4 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio4() },
		"--na-pratica-exercicio-5":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio5() },
		"--na-pratica-exercicio-5 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio5() },
		"--na-pratica-exercicio-6":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio6() },
		"--na-pratica-exercicio-6 --prova":     func() { exercicios_ninja_nivel_1.QuestionarioProva() },
		"--outline":                            func() { outline.Outline() },
		"--help":                               func() { PrintHelpMe() },
	}

	argsStr := strings.Join(args, " ")

	if action, ok := options[argsStr]; ok {
		action()
	} else {
		fmt.Println("\nOpção inválida")
		PrintHelpMe()
	}
}
