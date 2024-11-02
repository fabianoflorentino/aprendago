package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func MenuCapituloOutline([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap-1 --overview", ExecFunc: func() { visao_geral_do_curso.VisaoGeralDoCurso() }},
		{Options: "--cap-2 --overview", ExecFunc: func() { variaveis_valores_tipos.VariaveisValoresETipos() }},
		{Options: "--cap-3 --overview", ExecFunc: func() { exercicios_ninja_nivel_3.ExerciciosNinjaNivel3() }},
		{Options: "--cap-4 --overview", ExecFunc: func() { fundamentos_da_programacao.FundamentosDaProgramacao() }},
		{Options: "--cap-5 --overview", ExecFunc: func() { exercicios_ninja_nivel_2.ExerciciosNinjaNivel2() }},
		{Options: "--cap-6 --overview", ExecFunc: func() { fluxo_de_controle.FluxoDeControle() }},
	}
}

func HelpMeCapituloOutline() {
	hlp := []format.HelpMe{
		{Flag: "--cap-1 --overview", Description: "Visão Geral do Curso", Width: 0},
		{Flag: "--cap-2 --overview", Description: "Variáveis, Valores & Tipos", Width: 0},
		{Flag: "--cap-3 --overview", Description: "Exercícios Ninja Nível 1", Width: 0},
		{Flag: "--cap-4 --overview", Description: "Fundamentos da Programação", Width: 0},
		{Flag: "--cap-5 --overview", Description: "Exercícios Ninja Nível 2", Width: 0},
		{Flag: "--cap-6 --overview", Description: "Fluxo de Controle", Width: 0},
	}

	fmt.Printf("\nOutline do Curso por Capítulo\n")
	format.PrintHelpMe(hlp)
}
