package outline

import (
	"fmt"

	exercicios_ninja_nivel_1 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	exercicios_ninja_nivel_2 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	fluxo_de_controle "github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	fumdamentos_da_programacao "github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

const HEADER = `
Uso: go run main.go [opção]

Exemplo:
	go run main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
	`

func ShowHelpMe() {
	fmt.Println(HEADER)
	visao_geral_do_curso.HelpMeVisaoGeralDoCurso()
	variaveis_valores_tipos.HelpMeVariaveisValoresTipos()
	exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1()
	fumdamentos_da_programacao.HelpMeFundamentosDaProgramacao()
	exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2()
	fluxo_de_controle.HelpMeFluxoDeControle()
}
