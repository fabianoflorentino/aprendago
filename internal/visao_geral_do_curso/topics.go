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

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
