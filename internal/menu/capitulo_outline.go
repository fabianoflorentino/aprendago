package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func MenuCapituloOutline([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap-1 --outline", ExecFunc: func() { visao_geral_do_curso.VisaoGeralDoCurso() }},
		{Options: "--cap-2 --variaveis-valores-tipos", ExecFunc: func() { variaveis_valores_tipos.VariaveisValoresETipos() }},
		{Options: "--cap-3 --exercicios-ninja-nivlel-1", ExecFunc: func() { exercicios_ninja_nivel_3.ExerciciosNinjaNivel3() }},
		{Options: "--cap-4 --fundamentos-da-programacao", ExecFunc: func() { fundamentos_da_programacao.FundamentosDaProgramacao() }},
	}
}

func HelpMeCapituloOutline() {
	hlp := []format.HelpMe{
		{Flag: "--cap-1 --outline", Description: "Visão Geral do Curso", Width: 0},
		{Flag: "--cap-2 --variaveis-valores-tipos", Description: "Variáveis, Valores & Tipos", Width: 0},
		{Flag: "--cap-3 --exercicios-ninja-nivlel-1", Description: "Exercícios Ninja Nível 1", Width: 0},
		{Flag: "--cap-4 --fundamentos-da-programacao", Description: "Fundamentos da Programação", Width: 0},
	}

	fmt.Printf("\nOutline do Curso por Capítulo\n")
	format.PrintHelpMe(hlp)
}
