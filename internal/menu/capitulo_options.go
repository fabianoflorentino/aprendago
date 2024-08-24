package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func MenuCapituloOptions([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap-1", ExecFunc: func() { visao_geral_do_curso.HelpMeVisaoGeralDoCurso() }},
		{Options: "--cap-2", ExecFunc: func() { variaveis_valores_tipos.HelpMeVariaveisValoresTipos() }},
		{Options: "--cap-3", ExecFunc: func() { exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1() }},
		{Options: "--cap-4", ExecFunc: func() { fundamentos_da_programacao.HelpMeFundamentosDaProgramacao() }},
		{Options: "--cap-5", ExecFunc: func() { exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2() }},
		{Options: "--cap-6", ExecFunc: func() { fluxo_de_controle.HelpMeFluxoDeControle() }},
		{Options: "--cap-7", ExecFunc: func() { exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3() }},
		{Options: "--cap-8", ExecFunc: func() { agrupamento_de_dados.HelpMeAgrupamentoDeDados() }},
		{Options: "--cap-9", ExecFunc: func() { exercicios_ninja_nivel_4.HelpMeExerciciosNinjaNivel4() }},
		{Options: "--cap-10", ExecFunc: func() { structs.HelpMeStructs() }},
		{Options: "--cap-11", ExecFunc: func() { exercicios_ninja_nivel_5.HelpMeExerciciosNinjaNivel5() }},
		{Options: "--cap-12", ExecFunc: func() { funcoes.HelpMeFuncoes() }},
		{Options: "--cap-13", ExecFunc: func() { exercicios_ninja_nivel_6.HelMeExerciciosNinjaNivel6() }},
	}
}

func HelpMeCapituloOptions() {
	hlp := []format.HelpMe{
		{Flag: "--cap-1", Description: "Visão Geral do Curso", Width: 0},
		{Flag: "--cap-2", Description: "Variáveis, Valores e Tipos", Width: 0},
		{Flag: "--cap-3", Description: "Exercícios Ninja: Nível 1", Width: 0},
		{Flag: "--cap-4", Description: "Fundamentos da Programação", Width: 0},
		{Flag: "--cap-5", Description: "Exercícios Ninja: Nível 2", Width: 0},
		{Flag: "--cap-6", Description: "Fluxo de Controle", Width: 0},
		{Flag: "--cap-7", Description: "Exercícios Ninja: Nível 3", Width: 0},
		{Flag: "--cap-8", Description: "Agrupamento de Dados", Width: 0},
		{Flag: "--cap-9", Description: "Exercícios Ninja: Nível 4", Width: 0},
		{Flag: "--cap-10", Description: "Structs", Width: 0},
		{Flag: "--cap-11", Description: "Exercícios Ninja: Nível 5", Width: 0},
		{Flag: "--cap-12", Description: "Funções", Width: 0},
		{Flag: "--cap-13", Description: "Exercícios Ninja: Nível 6", Width: 0},
	}

	fmt.Println("Capítulos do Curso")
	format.PrintHelpMe(hlp)
}
