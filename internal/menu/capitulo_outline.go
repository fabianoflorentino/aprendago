// Package menu provides functionality to display and handle the menu options
// for different chapters of the "aprendago" course. It includes functions to
// generate menu options and display help information for each chapter.
package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/documentacao"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_11"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_12"
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
	"github.com/fabianoflorentino/aprendago/internal/tratamento_de_erro"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// MenuCapituloOutline returns a slice of format.MenuOptions, each representing
// a chapter outline with an associated command-line option and execution function.
// The options are structured as "--cap=<chapter_number> --overview" and the
// corresponding ExecFunc executes a specific function related to that chapter.
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
		{Options: "--cap=21 --overview", ExecFunc: func() { canais.Canais() }},
		{Options: "--cap=22 --overview", ExecFunc: func() { exercicios_ninja_nivel_10.ExerciciosNinjaNivel10() }},
		{Options: "--cap=23 --overview", ExecFunc: func() { tratamento_de_erro.TratamentoDeErro() }},
		{Options: "--cap=24 --overview", ExecFunc: func() { exercicios_ninja_nivel_11.ExerciciosNinjaNivel11() }},
		{Options: "--cap=25 --overview", ExecFunc: func() { documentacao.Documentacao() }},
		{Options: "--cap=26 --overview", ExecFunc: func() { exercicios_ninja_nivel_12.ExerciciosNinjaNivel12() }},
	}
}

// HelpMeCapituloOutline prints the outline of the course chapters.
// Each chapter is represented by a flag and a description, which are stored in a slice of HelpMe structs.
// The function then prints the course outline using the PrintHelpMe function from the format package.
func HelpMeCapituloOutline() {
	hlp := []format.HelpMe{
		{Flag: "--cap=1 --overview", Description: "Visão Geral do Curso"},
		{Flag: "--cap=2 --overview", Description: "Variáveis, Valores & Tipos"},
		{Flag: "--cap=3 --overview", Description: "Exercícios Ninja Nível 1"},
		{Flag: "--cap=4 --overview", Description: "Fundamentos da Programação"},
		{Flag: "--cap=5 --overview", Description: "Exercícios Ninja Nível 2"},
		{Flag: "--cap=6 --overview", Description: "Fluxo de Controle"},
		{Flag: "--cap=7 --overview", Description: "Exercícios Ninja Nível 3"},
		{Flag: "--cap=8 --overview", Description: "Agrupamento de Dados"},
		{Flag: "--cap=9 --overview", Description: "Exercícios Ninja Nível 4"},
		{Flag: "--cap=10 --overview", Description: "Structs"},
		{Flag: "--cap=11 --overview", Description: "Exercícios Ninja Nível 5"},
		{Flag: "--cap=12 --overview", Description: "Funções"},
		{Flag: "--cap=13 --overview", Description: "Exercícios Ninja Nível 6"},
		{Flag: "--cap=14 --overview", Description: "Ponteiros"},
		{Flag: "--cap=15 --overview", Description: "Exercícios Ninja Nível 7"},
		{Flag: "--cap=16 --overview", Description: "Aplicações"},
		{Flag: "--cap=17 --overview", Description: "Exercícios Ninja Nível 8"},
		{Flag: "--cap=18 --overview", Description: "Concorrência"},
		{Flag: "--cap=19 --overview", Description: "Seu Ambiente de Desenvolvimento"},
		{Flag: "--cap=20 --overview", Description: "Exercícios Ninja Nível 9"},
		{Flag: "--cap=21 --overview", Description: "Canais"},
		{Flag: "--cap=22 --overview", Description: "Exercícios Ninja Nível 10"},
		{Flag: "--cap=23 --overview", Description: "Tratamento de Erro"},
		{Flag: "--cap=24 --overview", Description: "Exercícios Ninja Nível 11"},
		{Flag: "--cap=25 --overview", Description: "Documentação"},
		{Flag: "--cap=26 --overview", Description: "Exercícios Ninja Nível 12"},
	}

	fmt.Printf("\nOutline do Curso por Capítulo\n")
	format.PrintHelpMe(hlp)
}
