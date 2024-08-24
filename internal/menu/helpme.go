package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
)

const HEADER = `
Uso: go run main.go [opção]

Exemplo:
  go run main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.
	`

// ShowHelpMe exibe a lista de todas as opções disponíveis.
// Esta função é chamada quando o usuário passa a opção --help.
func HelpMe() {
	fmt.Println(HEADER)
	HelpMeCapituloOptions()
	HelpMeCapituloOutline()
	visao_geral_do_curso.HelpMeVisaoGeralDoCurso()
	variaveis_valores_tipos.HelpMeVariaveisValoresTipos()
	exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1()
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
}
