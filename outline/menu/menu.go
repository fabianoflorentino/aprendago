package outline

import (
	"fmt"
	"strings"

	outline "github.com/fabianoflorentino/aprendago/outline"
	exercicios_ninja_nivel_1 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	exercicios_ninja_nivel_2 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	exercicios_ninja_nivel_3 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_3"
	fluxo_de_controle "github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	fundamentos_da_programacao "github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Menu(args []string) {
	fmt.Printf("Aprenda GO\n\n")

	options := map[string]func(){
		"--bem-vindo":                                    func() { visao_geral_do_curso.BemVindo() },
		"--porque-go":                                    func() { visao_geral_do_curso.PorQueGo() },
		"--sucesso":                                      func() { visao_geral_do_curso.Sucesso() },
		"--recursos":                                     func() { visao_geral_do_curso.Recursos() },
		"--como-esse-curso-funciona":                     func() { visao_geral_do_curso.ComoEsseCursoFunciona() },
		"--go-playground":                                func() { variaveis_valores_tipos.GoPlayground() },
		"--hello-world":                                  func() { variaveis_valores_tipos.HelloWorld() },
		"--operador-curto-de-declaracao":                 func() { variaveis_valores_tipos.OperadorCurtoDeDeclaracao() },
		"--a-palavra-reservada-var":                      func() { variaveis_valores_tipos.ApalavraReservadaVar() },
		"--explorando-tipos":                             func() { variaveis_valores_tipos.ExplorandoTipos() },
		"--valor-zero":                                   func() { variaveis_valores_tipos.ValorZero() },
		"--o-pacote-fmt":                                 func() { variaveis_valores_tipos.OpacoteFmt() },
		"--criando-seu-proprio-tipo":                     func() { variaveis_valores_tipos.CriandoSeuProprioTipo() },
		"--conversao-nao-coercao":                        func() { variaveis_valores_tipos.ConversaoNaoCoercao() },
		"--contribua-seu-codigo":                         func() { exercicios_ninja_nivel_1.ContribuaSeuCodigo() },
		"--na-pratica-exercicio-1 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio1() },
		"--na-pratica-exercicio-1 --nivel-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio1() },
		"--na-pratica-exercicio-2 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio2() },
		"--na-pratica-exercicio-2 --nivel-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio2() },
		"--na-pratica-exercicio-3 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio3() },
		"--na-pratica-exercicio-3 --nivel-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio3() },
		"--na-pratica-exercicio-4 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio4() },
		"--na-pratica-exercicio-4 --nivel-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio4() },
		"--na-pratica-exercicio-5 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio5() },
		"--na-pratica-exercicio-5 --nivel-1 --resolucao": func() { exercicios_ninja_nivel_1.ResolucaoNaPraticaExercicio5() },
		"--na-pratica-exercicio-6 --nivel-1":             func() { exercicios_ninja_nivel_1.NaPraticaExercicio6() },
		"--na-pratica-exercicio-6 --nivel-1 --prova":     func() { exercicios_ninja_nivel_1.RespondaAProva() },
		"--tipo-booleano":                                func() { fundamentos_da_programacao.TipoBooleano() },
		"--como-os-computadores-funcionam":               func() { fundamentos_da_programacao.ComoOsComputadoresFuncionam() },
		"--tipos-numericos":                              func() { fundamentos_da_programacao.TiposNumericos() },
		"--overflow":                                     func() { fundamentos_da_programacao.Overflow() },
		"--tipo-string":                                  func() { fundamentos_da_programacao.TipoString() },
		"--sistemas-numericos":                           func() { fundamentos_da_programacao.SistemasNumericos() },
		"--constantes":                                   func() { fundamentos_da_programacao.Constantes() },
		"--iota":                                         func() { fundamentos_da_programacao.Iota() },
		"--deslocamento-de-bits":                         func() { fundamentos_da_programacao.DeslocamentoDeBits() },
		"--na-pratica-exercicio-1 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio1() },
		"--na-pratica-exercicio-1 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio1() },
		"--na-pratica-exercicio-2 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio2() },
		"--na-pratica-exercicio-2 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio2() },
		"--na-pratica-exercicio-3 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio3() },
		"--na-pratica-exercicio-3 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio3() },
		"--na-pratica-exercicio-4 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio4() },
		"--na-pratica-exercicio-4 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio4() },
		"--na-pratica-exercicio-5 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio5() },
		"--na-pratica-exercicio-5 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio5() },
		"--na-pratica-exercicio-6 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio6() },
		"--na-pratica-exercicio-6 --nivel-2 --resolucao": func() { exercicios_ninja_nivel_2.ResolucaoNaPraticaExercicio6() },
		"--na-pratica-exercicio-7 --nivel-2":             func() { exercicios_ninja_nivel_2.NaPraticaExercicio7() },
		"--na-pratica-exercicio-7 --nivel-2 --prova":     func() { exercicios_ninja_nivel_2.RespondaAProva() },
		"--entendendo-fluxo-de-controle":                 func() { fluxo_de_controle.EntendendoFluxoDeControle() },
		"--loops-inicializacao-condicao-pos":             func() { fluxo_de_controle.LoopsInicializacaoCondicaoPos() },
		"--loops-nested-loop":                            func() { fluxo_de_controle.LoopsNestedLoop() },
		"--loops-a-declaracao-for":                       func() { fluxo_de_controle.LoopsADeclaracaoFor() },
		"--loops-break-e-continue":                       func() { fluxo_de_controle.LoopsBreakEContinue() },
		"--loops-utilizando-ascii":                       func() { fluxo_de_controle.LoopsUtilizandoAscii() },
		"--loops-utilizando-ascii --resolucao":           func() { fluxo_de_controle.ResolucaoLoopsUtilizandoAscii() },
		"--condicionais-a-declaracao-if":                 func() { fluxo_de_controle.CondicionaisADeclaracaoIf() },
		"--condicionais-if-else-if-else":                 func() { fluxo_de_controle.CondicionaisIfElseIfElse() },
		"--condicionais-a-declaracao-switch":             func() { fluxo_de_controle.CondicionaisADeclaracaoSwitch() },
		"--condicionais-a-declaracao-switch-pt2":         func() { fluxo_de_controle.CondicionaisADeclaracaoSwitchPt2() },
		"--operadores-logicos-condicionais":              func() { fluxo_de_controle.OperadoresLogicosCondicionais() },
		"--na-pratica-exercicio-1 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio1() },
		"--na-pratica-exercicio-1 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio1() },
		"--na-pratica-exercicio-2 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio2() },
		"--na-pratica-exercicio-2 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio2() },
		"--na-pratica-exercicio-3 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio3() },
		"--na-pratica-exercicio-3 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio3() },
		"--na-pratica-exercicio-4 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio4() },
		"--na-pratica-exercicio-4 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio4() },
		"--na-pratica-exercicio-5 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio5() },
		"--na-pratica-exercicio-5 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio5() },
		"--na-pratica-exercicio-6 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio6() },
		"--na-pratica-exercicio-6 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio6() },
		"--na-pratica-exercicio-7 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio7() },
		"--na-pratica-exercicio-7 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio7() },
		"--na-pratica-exercicio-8 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio8() },
		"--na-pratica-exercicio-8 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio8() },
		"--na-pratica-exercicio-9 --nivel-3":             func() { exercicios_ninja_nivel_3.NaPraticaExercicio9() },
		"--na-pratica-exercicio-9 --nivel-3 --resolucao": func() { exercicios_ninja_nivel_3.ResolucaoNaPraticaExercicio9() },
		"--outline": func() { outline.Outline() },
		"--help":    func() { ShowHelpMe() },
	}

	argsStr := strings.Join(args, " ")

	if action, ok := options[argsStr]; ok {
		action()
	} else {
		fmt.Println("\nOpção inválida")
		ShowHelpMe()
	}
}
