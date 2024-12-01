package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
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
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const HEADER = `
Uso: go run cmd/aprendago/main.go [opção]

Exemplo:
  go run cmd/aprendago/main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.
`

// ShowHelpMe exibe a lista de todas as opções disponíveis.
// Esta função é chamada quando o usuário passa a opção --help.
func HelpMe() {
	fmt.Printf("%s\n", HEADER)
	HelpMeCapituloOptions()
	HelpMeCapituloOutline()
	format.PrintChapterHelpMe(1)
	format.PrintChapterHelpMe(2)
	format.PrintChapterHelpMe(3)
	fundamentos_da_programacao.HelpMeFundamentosDaProgramacao()
	exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2()
	fluxo_de_controle.HelpMeFluxoDeControle()
	exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3()
	agrupamento_de_dados.HelpMeAgrupamentoDeDados()
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
}
