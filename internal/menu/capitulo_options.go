// Package menu provides functionality to display and handle menu options for different chapters of a Go programming course.
// It includes functions to generate menu options and display help information for each chapter.
package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/documentacao"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_11"
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

// MenuCapituloOptions returns a slice of format.MenuOptions, each representing a chapter option with associated topics and execution functions.
// The options are structured with a specific format "--cap=<chapter_number> --topics" and each option has an ExecFunc that calls a corresponding help function.
func MenuCapituloOptions([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap=1 --topics", ExecFunc: func() { visao_geral_do_curso.HelpMeVisaoGeralDoCurso() }},
		{Options: "--cap=2 --topics", ExecFunc: func() { variaveis_valores_tipos.HelpMeVariaveisValoresTipos() }},
		{Options: "--cap=3 --topics", ExecFunc: func() { exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1() }},
		{Options: "--cap=4 --topics", ExecFunc: func() { fundamentos_da_programacao.HelpMeFundamentosDaProgramacao() }},
		{Options: "--cap=5 --topics", ExecFunc: func() { exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2() }},
		{Options: "--cap=6 --topics", ExecFunc: func() { fluxo_de_controle.HelpMeFluxoDeControle() }},
		{Options: "--cap=7 --topics", ExecFunc: func() { exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3() }},
		{Options: "--cap=8 --topics", ExecFunc: func() { agrupamento_de_dados.HelpMeAgrupamentoDeDados() }},
		{Options: "--cap=9 --topics", ExecFunc: func() { exercicios_ninja_nivel_4.HelpMeExerciciosNinjaNivel4() }},
		{Options: "--cap=10 --topics", ExecFunc: func() { structs.HelpMeStructs() }},
		{Options: "--cap=11 --topics", ExecFunc: func() { exercicios_ninja_nivel_5.HelpMeExerciciosNinjaNivel5() }},
		{Options: "--cap=12 --topics", ExecFunc: func() { funcoes.HelpMeFuncoes() }},
		{Options: "--cap=13 --topics", ExecFunc: func() { exercicios_ninja_nivel_6.HelMeExerciciosNinjaNivel6() }},
		{Options: "--cap=14 --topics", ExecFunc: func() { ponteiros.HelpMePonteiros() }},
		{Options: "--cap=15 --topics", ExecFunc: func() { exercicios_ninja_nivel_7.HelpMeExerciciosNinjaNivel7() }},
		{Options: "--cap=16 --topics", ExecFunc: func() { aplicacoes.HelpMeAplicacoes() }},
		{Options: "--cap=17 --topics", ExecFunc: func() { exercicios_ninja_nivel_8.HelpMeExerciciosNinjaNivel8() }},
		{Options: "--cap=18 --topics", ExecFunc: func() { concorrencia.HelpMeConcorrencia() }},
		{Options: "--cap=19 --topics", ExecFunc: func() { seu_ambiente_de_desenvolvimento.HelpMeSeuAmbienteDeDesenvolvimento() }},
		{Options: "--cap=20 --topics", ExecFunc: func() { exercicios_ninja_nivel_9.HelpMeExerciciosNinjaNivel9() }},
		{Options: "--cap=21 --topics", ExecFunc: func() { canais.HelpMeCanais() }},
		{Options: "--cap=22 --topics", ExecFunc: func() { exercicios_ninja_nivel_10.HelpMeExerciciosNinjaNivel10() }},
		{Options: "--cap=23 --topics", ExecFunc: func() { tratamento_de_erro.HelpMeTratamentoDeErro() }},
		{Options: "--cap=24 --topics", ExecFunc: func() { exercicios_ninja_nivel_11.HelpMeExerciciosNinjaNivel11() }},
		{Options: "--cap=25 --topics", ExecFunc: func() { documentacao.HelpMeDocumentacao() }},
	}
}

// HelpMeCapituloOptions prints a list of course chapters and their descriptions.
// Each chapter is represented by a flag and a description, which are printed
// in a formatted manner to the console. This function is useful for providing
// an overview of the course content and helping users navigate through different
// topics and exercises.
func HelpMeCapituloOptions() {
	hlp := []format.HelpMe{
		{Flag: "--cap=1 --topics", Description: "Visão Geral do Curso"},
		{Flag: "--cap=2 --topics", Description: "Variáveis, Valores e Tipos"},
		{Flag: "--cap=3 --topics", Description: "Exercícios Ninja: Nível 1"},
		{Flag: "--cap=4 --topics", Description: "Fundamentos da Programação"},
		{Flag: "--cap=5 --topics", Description: "Exercícios Ninja: Nível 2"},
		{Flag: "--cap=6 --topics", Description: "Fluxo de Controle"},
		{Flag: "--cap=7 --topics", Description: "Exercícios Ninja: Nível 3"},
		{Flag: "--cap=8 --topics", Description: "Agrupamento de Dados"},
		{Flag: "--cap=9 --topics", Description: "Exercícios Ninja: Nível 4"},
		{Flag: "--cap=10 --topics", Description: "Structs"},
		{Flag: "--cap=11 --topics", Description: "Exercícios Ninja: Nível 5"},
		{Flag: "--cap=12 --topics", Description: "Funções"},
		{Flag: "--cap=13 --topics", Description: "Exercícios Ninja: Nível 6"},
		{Flag: "--cap=14 --topics", Description: "Ponteiros"},
		{Flag: "--cap=15 --topics", Description: "Exercícios Ninja: Nível 7"},
		{Flag: "--cap=16 --topics", Description: "Aplicações"},
		{Flag: "--cap=17 --topics", Description: "Exercícios Ninja: Nível 8"},
		{Flag: "--cap=18 --topics", Description: "Concorrência"},
		{Flag: "--cap=19 --topics", Description: "Seu Ambiente de Desenvolvimento"},
		{Flag: "--cap=20 --topics", Description: "Exercícios Ninja: Nível 9"},
		{Flag: "--cap=21 --topics", Description: "Canais"},
		{Flag: "--cap=22 --topics", Description: "Exercícios Ninja: Nível 10"},
		{Flag: "--cap=23 --topics", Description: "Tratamento de Erro"},
		{Flag: "--cap=24 --topics", Description: "Exercícios Ninja: Nível 11"},
		{Flag: "--cap=25 --topics", Description: "Documentação"},
	}

	fmt.Println("Capítulos do Curso")
	format.PrintHelpMe(hlp)
}
