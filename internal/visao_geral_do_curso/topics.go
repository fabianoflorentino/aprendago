package visao_geral_do_curso

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/visao_geral_do_curso"

func VisaoGeralDoCurso() {
	fmt.Printf("\n01 - Visão Geral do Curso\n")

	executeSection("Bem-vindo!")
	executeSection("Por que Go?")
	executeSection("Sucesso")
	executeSection("Recursos")
	executeSection("Como esse curso funciona")
}

func MenuVisaoGeralDoCurso([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--bem-vindo", ExecFunc: func() { executeSection("Bem-vindo!") }},
		{Options: "--porque-go", ExecFunc: func() { executeSection("Por que Go?") }},
		{Options: "--sucesso", ExecFunc: func() { executeSection("Sucesso") }},
		{Options: "--recursos", ExecFunc: func() { executeSection("Recursos") }},
		{Options: "--como-esse-curso-funciona", ExecFunc: func() { executeSection("Como esse curso funciona") }},
	}
}

func HelpMeVisaoGeralDoCurso() {
	hlp := []format.HelpMe{
		{Flag: "--bem-vindo", Description: "Exibe a mensagem de boas-vindas ao curso Aprenda Go.", Width: 0},
		{Flag: "--porque-go", Description: "Descreve os benefícios e razões para aprender a linguagem Go.", Width: 0},
		{Flag: "--sucesso", Description: "Apresenta dicas e estratégias para ter sucesso no curso.", Width: 0},
		{Flag: "--recursos", Description: "Lista recursos e materiais de apoio para o curso.", Width: 0},
		{Flag: "--como-esse-curso-funciona", Description: "Explica a estrutura e metodologia do curso.", Width: 0},
	}

	fmt.Println("\nCapítulo 1: Visão Geral do Curso")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
