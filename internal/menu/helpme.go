// Package menu provides functionality to display help and outline information
// for the "aprendago" application. It includes a function to show all available
// options and detailed help for each chapter of the course.
package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/documentacao"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_11"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_12"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_7"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_8"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_9"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/ponteiros"
	"github.com/fabianoflorentino/aprendago/internal/seu_ambiente_de_desenvolvimento"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/teste_benchmarking"
	"github.com/fabianoflorentino/aprendago/internal/tratamento_de_erro"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
)

// HEADER is a constant string that provides usage instructions for running the Go program.
// It includes an example command and descriptions of available options such as --outline, --help, and --caps.
const HEADER = `
Uso: go run cmd/aprendago/main.go [opção]

Exemplo:
  go run cmd/aprendago/main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.
`

// HelpMe provides a comprehensive guide through various chapters and exercises
// of the Go programming course. It prints the header and sequentially calls
// helper functions for each chapter and exercise level, offering an overview
// and detailed explanations of variables, control flow, data grouping, structs,
// functions, pointers, concurrency, error handling, and more.
func HelpMe() {
	fmt.Printf("%s\n", HEADER)
	HelpMeCapituloOptions()
	HelpMeCapituloOutline()
	visao_geral_do_curso.HelpMeVisaoGeralDoCurso()
	variaveis_valores_tipos.HelpMeVariaveisValoresTipos()
	exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1()
	fundamentos_da_programacao.HelpMeFundamentosDaProgramacao()
	exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2()
	fluxo_de_controle.HelpMeFluxoDeControle()
	exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3()
	agrupamento_de_dados.Help()
	exercicios_ninja_nivel_4.HelpMeExerciciosNinjaNivel4()
	structs.HelpMeStructs()
	exercicios_ninja_nivel_5.HelpMeExerciciosNinjaNivel5()
	funcoes.HelpMeFuncoes()
	exercicios_ninja_nivel_6.HelMeExerciciosNinjaNivel6()
	ponteiros.HelpMePonteiros()
	exercicios_ninja_nivel_7.HelpMeExerciciosNinjaNivel7()
	aplicacoes.HelpMeAplicacoes()
	exercicios_ninja_nivel_8.HelpMeExerciciosNinjaNivel8()
	concorrencia.HelpMeConcorrencia()
	seu_ambiente_de_desenvolvimento.HelpMeSeuAmbienteDeDesenvolvimento()
	exercicios_ninja_nivel_9.HelpMeExerciciosNinjaNivel9()
	canais.HelpMeCanais()
	exercicios_ninja_nivel_10.HelpMeExerciciosNinjaNivel10()
	tratamento_de_erro.HelpMeTratamentoDeErro()
	exercicios_ninja_nivel_11.HelpMeExerciciosNinjaNivel11()
	documentacao.HelpMeDocumentacao()
	exercicios_ninja_nivel_12.HelpMeExerciciosNinjaNivel12()
	teste_benchmarking.HelpMeTesteEBenchmarking()
}
