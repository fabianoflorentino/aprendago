package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuVisaoGeralDoCurso([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--bem-vindo", ExecFunc: func() { BemVindo() }},
		{Options: "--porque-go", ExecFunc: func() { PorQueGo() }},
		{Options: "--sucesso", ExecFunc: func() { Sucesso() }},
		{Options: "--recursos", ExecFunc: func() { Recursos() }},
		{Options: "--como-esse-curso-funciona", ExecFunc: func() { ComoEsseCursoFunciona() }},
	}
}
