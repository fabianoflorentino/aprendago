package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
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
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func MenuCapituloOutline([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap=1 --overview", ExecFunc: func() { visao_geral_do_curso.VisaoGeralDoCurso() }},
		{Options: "--cap=2 --overview", ExecFunc: func() { variaveis_valores_tipos.VariaveisValoresETipos() }},
		{Options: "--cap=3 --overview", ExecFunc: func() { exercicios_ninja_nivel_3.ExerciciosNinjaNivel3() }},
		{Options: "--cap=4 --overview", ExecFunc: func() { fundamentos_da_programacao.FundamentosDaProgramacao() }},
		{Options: "--cap=5 --overview", ExecFunc: func() { exercicios_ninja_nivel_2.ExerciciosNinjaNivel2() }},
		{Options: "--cap=6 --overview", ExecFunc: func() { fluxo_de_controle.FluxoDeControle() }},
		{Options: "--cap=7 --overview", ExecFunc: func() { exercicios_ninja_nivel_3.ExerciciosNinjaNivel3() }},
		{Options: "--cap=8 --overview", ExecFunc: func() { agrupamento_de_dados.AgrupamentoDeDados() }},
		{Options: "--cap=9 --overview", ExecFunc: func() { exercicios_ninja_nivel_4.ExerciciosNinjaNivel4() }},
		{Options: "--cap=10 --overview", ExecFunc: func() { structs.TopicStructs() }},
		{Options: "--cap=11 --overview", ExecFunc: func() { exercicios_ninja_nivel_5.ExerciciosNinjaNivel5() }},
		{Options: "--cap=12 --overview", ExecFunc: func() { funcoes.Funcoes() }},
		{Options: "--cap=13 --overview", ExecFunc: func() { exercicios_ninja_nivel_6.ExerciciosNinjaNivel6() }},
		{Options: "--cap=14 --overview", ExecFunc: func() { ponteiros.Ponteiros() }},
		{Options: "--cap=15 --overview", ExecFunc: func() { exercicios_ninja_nivel_7.ExerciciosNinjaNivel7() }},
		{Options: "--cap=16 --overview", ExecFunc: func() { aplicacoes.Aplicacoes() }},
		{Options: "--cap=17 --overview", ExecFunc: func() { exercicios_ninja_nivel_8.ExerciciosNinjaNivel8() }},
		{Options: "--cap=18 --overview", ExecFunc: func() { concorrencia.Concorrencia() }},
		{Options: "--cap=19 --overview", ExecFunc: func() { seu_ambiente_de_desenvolvimento.SeuAmbienteDeDesenvolvimento() }},
		{Options: "--cap=20 --overview", ExecFunc: func() { exercicios_ninja_nivel_9.ExerciciosNinjaNivel9() }},
	}
}

func HelpMeCapituloOutline() {
	hlp := []format.HelpMe{
		{Flag: "--cap=1 --overview", Description: "Visão Geral do Curso", Width: 0},
		{Flag: "--cap=2 --overview", Description: "Variáveis, Valores & Tipos", Width: 0},
		{Flag: "--cap=3 --overview", Description: "Exercícios Ninja Nível 1", Width: 0},
		{Flag: "--cap=4 --overview", Description: "Fundamentos da Programação", Width: 0},
		{Flag: "--cap=5 --overview", Description: "Exercícios Ninja Nível 2", Width: 0},
		{Flag: "--cap=6 --overview", Description: "Fluxo de Controle", Width: 0},
		{Flag: "--cap=7 --overview", Description: "Exercícios Ninja Nível 3", Width: 0},
		{Flag: "--cap=8 --overview", Description: "Agrupamento de Dados", Width: 0},
		{Flag: "--cap=9 --overview", Description: "Exercícios Ninja Nível 4", Width: 0},
		{Flag: "--cap=10 --overview", Description: "Structs", Width: 0},
		{Flag: "--cap=11 --overview", Description: "Exercícios Ninja Nível 5", Width: 0},
		{Flag: "--cap=12 --overview", Description: "Funções", Width: 0},
		{Flag: "--cap=13 --overview", Description: "Exercícios Ninja Nível 6", Width: 0},
		{Flag: "--cap=14 --overview", Description: "Ponteiros", Width: 0},
		{Flag: "--cap=15 --overview", Description: "Exercícios Ninja Nível 7", Width: 0},
		{Flag: "--cap=16 --overview", Description: "Aplicações", Width: 0},
		{Flag: "--cap=17 --overview", Description: "Exercícios Ninja Nível 8", Width: 0},
		{Flag: "--cap=18 --overview", Description: "Concorrência", Width: 0},
		{Flag: "--cap=19 --overview", Description: "Seu Ambiente de Desenvolvimento", Width: 0},
		{Flag: "--cap=20 --overview", Description: "Exercícios Ninja Nível 9", Width: 0},
	}

	fmt.Printf("\nOutline do Curso por Capítulo\n")
	format.PrintHelpMe(hlp)
}
