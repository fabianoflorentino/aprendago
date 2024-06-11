package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

const HEADER = `
Uso: go run main.go [opção]

Exemplo:
	go run main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
	`

// ShowHelpMe exibe a lista de todas as opções disponíveis.
// Esta função é chamada quando o usuário passa a opção --help.
func HelpMe() {
	fmt.Println(HEADER)
	visao_geral_do_curso.HelpMeVisaoGeralDoCurso()
	variaveis_valores_tipos.HelpMeVariaveisValoresTipos()
	exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1()
	fundamentos_da_programacao.HelpMeFundamentosDaProgramacao()
	exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2()
	fluxo_de_controle.HelpMeFluxoDeControle()
	exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3()
	agrupamento_de_dados.HelpMeAgrupamentoDeDados()
}
