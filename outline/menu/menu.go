package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
	variaveis_valores_tipos "github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	visao_geral_do_curso "github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

func Menu(args []string) {
	fmt.Printf("Aprenda GO\n\n")

	buildMenu(
		args,
		variaveis_valores_tipos.MenuVariaveisValoresTipos(args),
		visao_geral_do_curso.MenuVisaoGeralDoCurso(args),
	)
}

func buildMenu(args []string, options ...[]format.MenuOptions) {
	var opt []format.MenuOptions

	for _, o := range options {
		opt = append(opt, o...)
	}

	format.FormatMenuOptions(args, opt)
}
