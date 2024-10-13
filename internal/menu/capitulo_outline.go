package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func MenuCapituloOutline([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--cap-1 --outline", ExecFunc: func() { visao_geral_do_curso.VisaoGeralDoCurso() }},
		{Options: "--cap-2 --variaveis-valores-tipos", ExecFunc: func() { variaveis_valores_tipos.VariaveisValoresETipos() }},
	}
}

func HelpMeCapituloOutline() {
	hlp := []format.HelpMe{
		{Flag: "--cap-1 --outline", Description: "Visão Geral do Curso", Width: 0},
		{Flag: "--cap-2 --variaveis-valores-tipos", Description: "Variáveis, Valores & Tipos", Width: 0},
	}

	fmt.Printf("\nOutline do Curso por Capítulo\n")
	format.PrintHelpMe(hlp)
}
