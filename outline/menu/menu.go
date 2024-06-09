package outline

import (
	"fmt"

	outline "github.com/fabianoflorentino/aprendago/outline"
	exercicios_ninja_nivel_1 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	exercicios_ninja_nivel_2 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	exercicios_ninja_nivel_3 "github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_3"
	fluxo_de_controle "github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	format "github.com/fabianoflorentino/aprendago/outline/format"
	fundamentos_da_programacao "github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Menu(args []string) {
	fmt.Printf("Aprenda GO\n\n")

	buildMenu(
		args,
		generalMenu(args),
		variaveis_valores_tipos.MenuVariaveisValoresTipos(args),
		visao_geral_do_curso.MenuVisaoGeralDoCurso(args),
		exercicios_ninja_nivel_1.MenuExerciciosNinjaNivel1(args),
		fundamentos_da_programacao.MenuFundamentosDaProgramcao(args),
		exercicios_ninja_nivel_2.MenuExerciciosNinjaNivel2(args),
		fluxo_de_controle.MenuFluxoDeControle(args),
		exercicios_ninja_nivel_3.MenuExerciciosNinjaNivel3(args),
	)
}

func generalMenu([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--help", ExecFunc: func() { ShowHelpMe() }},
		{Options: "--outline", ExecFunc: func() { outline.Outline() }},
	}
}

func buildMenu(args []string, options ...[]format.MenuOptions) {
	var opt []format.MenuOptions

	for _, o := range options {
		opt = append(opt, o...)
	}

	format.FormatMenuOptions(args, opt)
}
