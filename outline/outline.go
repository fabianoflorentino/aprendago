// OUTLINE: https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md
package outline

import (
	exercicios_ninja_nivel_1 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	exercicios_ninja_nivel_2 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	exercicios_ninja_nivel_3 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_3"
	fluxo_de_controle "github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	fundamentos_da_programacao "github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Outline() {
	visao_geral_do_curso.VisaoGeralDoCurso()
	variaveis_valores_tipos.VariaveisValoresETipos()
	exercicios_ninja_nivel_1.ExerciciosNinjaNivel1()
	fundamentos_da_programacao.FundamentosDaProgramacao()
	exercicios_ninja_nivel_2.ExerciciosNinjaNivel2()
	fluxo_de_controle.FluxoDeControle()
	exercicios_ninja_nivel_3.ExerciciosNinjaNivel3()
}
