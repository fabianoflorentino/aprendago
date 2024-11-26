// OUTLINE: https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md
package internal

import (
	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_7"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/ponteiros"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
)

func Outline() {
	visao_geral_do_curso.VisaoGeralDoCurso()
	variaveis_valores_tipos.VariaveisValoresETipos()
	exercicios_ninja_nivel_1.ExerciciosNinjaNivel1()
	fundamentos_da_programacao.FundamentosDaProgramacao()
	exercicios_ninja_nivel_2.ExerciciosNinjaNivel2()
	fluxo_de_controle.FluxoDeControle()
	exercicios_ninja_nivel_3.ExerciciosNinjaNivel3()
	agrupamento_de_dados.AgrupamentoDeDados()
	exercicios_ninja_nivel_4.ExerciciosNinjaNivel4()
	structs.TopicStructs()
	exercicios_ninja_nivel_5.ExerciciosNinjaNivel5()
	funcoes.Funcoes()
	exercicios_ninja_nivel_6.ExerciciosNinjaNivel6()
	ponteiros.Ponteiros()
	exercicios_ninja_nivel_7.ExerciciosNinjaNivel7()
	aplicacoes.Aplicacoes()
	concorrencia.Concorrencia()
}
